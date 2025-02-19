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

type bankHandleApi struct {
	bank    pb.BankServiceClient
	logger  logger.LoggerInterface
	mapping response_api.BankResponseMapper
}

func NewHandlerBank(router *echo.Echo, bank pb.BankServiceClient, logger logger.LoggerInterface, mapping response_api.BankResponseMapper) *bankHandleApi {
	bankHandler := &bankHandleApi{
		bank:    bank,
		logger:  logger,
		mapping: mapping,
	}

	routerBank := router.Group("/api/bank")

	routerBank.GET("", bankHandler.FindAll)
	routerBank.GET("/:id", bankHandler.FindById)
	routerBank.GET("/active", bankHandler.FindByActive)
	routerBank.GET("/trashed", bankHandler.FindByTrashed)
	routerBank.POST("", bankHandler.Create)
	routerBank.POST("/:id", bankHandler.Update)
	routerBank.DELETE("/:id", bankHandler.Trashed)
	routerBank.PUT("/restore/:id", bankHandler.Restore)
	routerBank.DELETE("/permanent/:id", bankHandler.DeletePermanent)
	routerBank.PUT("/restore-all", bankHandler.RestoreAll)
	routerBank.DELETE("/permanent-all", bankHandler.DeleteAllPermanent)

	return bankHandler
}

// @Security Bearer
// @Summary Find all banks
// @Tags Bank
// @Description Retrieve a list of all banks
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationBank "List of banks"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve bank data"
// @Router /api/bank [get]
func (h *bankHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllBankRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.bank.FindAll(ctx, req)
	if err != nil {
		if errors.Is(err, echo.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status:  "error",
				Message: "Unauthorized",
			})
		}

		h.logger.Debug("Failed to fetch Bank records", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Bank records",
		})
	}

	so := h.mapping.ToApiResponsePaginationBank(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find bank by ID
// @Tags Bank
// @Description Retrieve a bank by ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} pb.ApiResponseBank "Bank data"
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve bank data"
// @Router /api/bank/{id} [get]
func (h *bankHandleApi) FindById(c echo.Context) error {
	BankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || BankID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Bank ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Bank", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Bank",
		})
	}

	so := h.mapping.ToApiResponseBank(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active banks
// @Tags Bank
// @Description Retrieve a list of active banks
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationBankDeleteAt "List of active banks"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve bank data"
// @Router /api/bank/active [get]
func (h *bankHandleApi) FindByActive(c echo.Context) error {
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

	req := &pb.FindAllBankRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.bank.FindByActive(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch active Banks", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Banks",
		})
	}

	so := h.mapping.ToApiResponsePaginationBankDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve trashed banks
// @Tags Bank
// @Description Retrieve a list of trashed bank records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationBankDeleteAt "List of trashed bank data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve bank data"
// @Router /api/bank/trashed [get]
func (h *bankHandleApi) FindByTrashed(c echo.Context) error {
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

	req := &pb.FindAllBankRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.bank.FindByTrashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch trashed Banks", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Banks",
		})
	}

	so := h.mapping.ToApiResponsePaginationBankDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create an existing bank
// @Tags Bank
// @Description Create an existing bank record with the provided details
// @Accept json
// @Produce json
// @Param CreateBankRequest body requests.CreateBankRequest true "Create bank request"
// @Success 200 {object} response.ApiResponseBank "Successfully created bank"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create bank"
// @Router /api/bank/create [post]
func (h *bankHandleApi) Create(c echo.Context) error {
	var req pb.CreateBankRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	ctx := c.Request().Context()

	res, err := h.bank.Create(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to create Bank", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create Bank",
		})
	}

	so := h.mapping.ToApiResponseBank(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update an existing bank
// @Tags Bank
// @Description Update an existing bank record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param UpdateBankRequest body requests.UpdateBankRequest true "Update bank request"
// @Success 200 {object} response.ApiResponseBank "Successfully updated bank"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update bank"
// @Router /api/bank/update/{id} [post]
func (h *bankHandleApi) Update(c echo.Context) error {
	BankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || BankID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Bank ID",
		})
	}

	var req pb.UpdateBankRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	req.Id = int32(BankID)

	ctx := c.Request().Context()

	res, err := h.bank.Update(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to update Bank", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update Bank",
		})
	}

	so := h.mapping.ToApiResponseBank(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve a trashed bank
// @Tags Bank
// @Description Retrieve a trashed bank record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} response.ApiResponseBankDeleteAt "Successfully retrieved trashed bank"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed bank"
// @Router /api/bank/trashed/{id} [get]
func (h *bankHandleApi) Trashed(c echo.Context) error {
	BankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || BankID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Bank ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Bank", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Bank",
		})
	}

	so := h.mapping.ToApiResponseBankDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore a trashed bank
// @Tags Bank
// @Description Restore a trashed bank record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} response.ApiResponseBankDeleteAt "Successfully restored bank"
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore bank"
// @Router /api/bank/restore/{id} [post]
func (h *bankHandleApi) Restore(c echo.Context) error {
	BankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || BankID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Bank ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Bank", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Bank",
		})
	}

	so := h.mapping.ToApiResponseBankDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete a bank
// @Tags Bank
// @Description Permanently delete a bank record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} pb.ApiResponseBankDelete "Successfully deleted bank record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete bank"
// @Router /api/bank/delete/{id} [post]
func (h *bankHandleApi) DeletePermanent(c echo.Context) error {
	BankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || BankID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Bank ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Bank permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Bank permanently",
		})
	}

	so := h.mapping.ToApiResponseBankDelete(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore all trashed banks
// @Tags Bank
// @Description Restore all trashed bank records.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseBankAll "Successfully restored all banks"
// @Failure 500 {object} response.ErrorResponse "Failed to restore banks"
// @Router /api/bank/restore/all [post]
func (h *bankHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.bank.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Banks", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Banks",
		})
	}

	so := h.mapping.ToApiResponseBankAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete all trashed banks
// @Tags Bank
// @Description Permanently delete all trashed bank records.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseBankDelete "Successfully deleted all bank records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete banks"
// @Router /api/bank/delete/all [post]
func (h *bankHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.bank.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Banks permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Banks permanently",
		})
	}

	so := h.mapping.ToApiResponseBankAll(res)

	return c.JSON(http.StatusOK, so)
}
