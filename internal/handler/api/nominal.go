package api

import (
	"errors"
	"net/http"
	"strconv"
	"topup_game/internal/domain/response"
	response_api "topup_game/internal/mapper/response/api"
	"topup_game/internal/pb"
	"topup_game/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type nominalHandleGrpc struct {
	nominal pb.NominalServiceClient
	logger  logger.LoggerInterface
	mapping response_api.NominalResponseMapper
}

func NewHandlerNominal(router *echo.Echo, nominal pb.NominalServiceClient, logger logger.LoggerInterface, mapping response_api.NominalResponseMapper) *nominalHandleGrpc {
	nominalHandler := &nominalHandleGrpc{
		nominal: nominal,
		logger:  logger,
		mapping: mapping,
	}

	routerNominal := router.Group("/api/nominal")

	routerNominal.GET("", nominalHandler.FindAll)
	routerNominal.GET("/:id", nominalHandler.FindById)
	routerNominal.GET("/active", nominalHandler.FindByActive)
	routerNominal.GET("/trashed", nominalHandler.FindByTrashed)
	routerNominal.POST("", nominalHandler.Create)
	routerNominal.POST("/:id", nominalHandler.Update)
	routerNominal.DELETE("/:id", nominalHandler.Trashed)
	routerNominal.PUT("/restore/:id", nominalHandler.Restore)
	routerNominal.DELETE("/permanent/:id", nominalHandler.DeletePermanent)
	routerNominal.PUT("/restore-all", nominalHandler.RestoreAll)
	routerNominal.DELETE("/permanent-all", nominalHandler.DeleteAllPermanent)

	return nominalHandler
}

// @Security Bearer
// @Summary Find all Nominals
// @Tags Nominal
// @Description Retrieve a list of all Nominals
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationNominal "List of Nominals"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Nominal data"
// @Router /api/nominal [get]
func (h *nominalHandleGrpc) FindAll(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllNominalRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.nominal.FindAll(ctx, req)
	if err != nil {
		if errors.Is(err, echo.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status:  "error",
				Message: "Unauthorized",
			})
		}

		h.logger.Debug("Failed to fetch Nominal records", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Nominal records",
		})
	}

	so := h.mapping.ToApiResponsePaginationNominal(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find Nominal by ID
// @Tags Nominal
// @Description Retrieve a Nominal by ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Success 200 {object} pb.ApiResponseNominal "Nominal data"
// @Failure 400 {object} response.ErrorResponse "Invalid Nominal ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Nominal data"
// @Router /api/nominal/{id} [get]
func (h *nominalHandleGrpc) FindById(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Nominal ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Nominal", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Nominal",
		})
	}

	so := h.mapping.ToApiResponseNominal(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active Nominals
// @Tags Nominal
// @Description Retrieve a list of active Nominals
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationNominalDeleteAt "List of active Nominals"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Nominal data"
// @Router /api/nominal/active [get]
func (h *nominalHandleGrpc) FindByActive(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllNominalRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.nominal.FindByActive(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch active Nominals", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Nominals",
		})
	}

	so := h.mapping.ToApiResponsePaginationNominalDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve trashed Nominals
// @Tags Nominal
// @Description Retrieve a list of trashed Nominal records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationNominalDeleteAt "List of trashed Nominal data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Nominal data"
// @Router /api/nominal/trashed [get]
func (h *nominalHandleGrpc) FindByTrashed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllNominalRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.nominal.FindByTrashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch trashed Nominals", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Nominals",
		})
	}

	so := h.mapping.ToApiResponsePaginationNominalDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create an existing Nominal
// @Tags Nominal
// @Description Create an existing Nominal record with the provided details
// @Accept json
// @Produce json
// @Param CreateNominalRequest body requests.CreateNominalRequest true "Create Nominal request"
// @Success 200 {object} response.ApiResponseNominal "Successfully created Nominal"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create Nominal"
// @Router /api/nominal/create [post]
func (h *nominalHandleGrpc) Create(c echo.Context) error {
	var req pb.CreateNominalRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	ctx := c.Request().Context()

	res, err := h.nominal.Create(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to create Nominal", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create Nominal",
		})
	}

	so := h.mapping.ToApiResponseNominal(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update an existing Nominal
// @Tags Nominal
// @Description Update an existing Nominal record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param UpdateNominalRequest body requests.UpdateNominalRequest true "Update Nominal request"
// @Success 200 {object} response.ApiResponseNominal "Successfully updated Nominal"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update Nominal"
// @Router /api/nominal/update/{id} [post]
func (h *nominalHandleGrpc) Update(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Nominal ID",
		})
	}

	var req pb.UpdateNominalRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	req.Id = int32(NominalID)

	ctx := c.Request().Context()

	res, err := h.nominal.Update(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to update Nominal", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update Nominal",
		})
	}

	so := h.mapping.ToApiResponseNominal(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve a trashed Nominal
// @Tags Nominal
// @Description Retrieve a trashed Nominal record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Success 200 {object} response.ApiResponseNominalDeleteAt "Successfully retrieved trashed Nominal"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed Nominal"
// @Router /api/nominal/trashed/{id} [get]
func (h *nominalHandleGrpc) Trashed(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Nominal ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Nominal", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Nominal",
		})
	}

	so := h.mapping.ToApiResponseNominalDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore a trashed Nominal
// @Tags Nominal
// @Description Restore a trashed Nominal record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Success 200 {object} response.ApiResponseNominalDeleteAt "Successfully restored Nominal"
// @Failure 400 {object} response.ErrorResponse "Invalid Nominal ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Nominal"
// @Router /api/nominal/restore/{id} [post]
func (h *nominalHandleGrpc) Restore(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Nominal ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Nominal", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Nominal",
		})
	}

	so := h.mapping.ToApiResponseNominalDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete a Nominal
// @Tags Nominal
// @Description Permanently delete a Nominal record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Success 200 {object} pb.ApiResponseNominalDelete "Successfully deleted Nominal record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Nominal"
// @Router /api/nominal/delete/{id} [post]
func (h *nominalHandleGrpc) DeletePermanent(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Nominal ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Nominal permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Nominal permanently",
		})
	}

	so := h.mapping.ToApiResponseNominalDelete(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore all trashed Nominals
// @Tags Nominal
// @Description Restore all trashed Nominal records.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseNominalAll "Successfully restored all Nominals"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Nominals"
// @Router /api/nominal/restore/all [post]
func (h *nominalHandleGrpc) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.nominal.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Nominals", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Nominals",
		})
	}

	so := h.mapping.ToApiResponseNominalAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete all trashed Nominals
// @Tags Nominal
// @Description Permanently delete all trashed Nominal records.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseNominalDelete "Successfully deleted all Nominal records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Nominals"
// @Router /api/nominal/delete/all [post]
func (h *nominalHandleGrpc) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.nominal.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Nominals permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Nominals permanently",
		})
	}

	so := h.mapping.ToApiResponseNominalAll(res)

	return c.JSON(http.StatusOK, so)
}
