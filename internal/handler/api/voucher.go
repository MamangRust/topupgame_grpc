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

type voucherHandleApi struct {
	voucher pb.VoucherServiceClient
	logger  logger.LoggerInterface
	mapping response_api.VoucherResponseMapper
}

func NewHandlerVoucher(router *echo.Echo, voucher pb.VoucherServiceClient, logger logger.LoggerInterface, mapping response_api.VoucherResponseMapper) *voucherHandleApi {
	voucherHandle := &voucherHandleApi{
		voucher: voucher,
		logger:  logger,
		mapping: mapping,
	}

	routerVoucher := router.Group("/api/voucher")

	routerVoucher.GET("", voucherHandle.FindAll)
	routerVoucher.GET("/:id", voucherHandle.FindById)
	routerVoucher.GET("/active", voucherHandle.FindByActive)
	routerVoucher.GET("/trashed", voucherHandle.FindByTrashed)
	routerVoucher.POST("", voucherHandle.Create)
	routerVoucher.POST("/:id", voucherHandle.Update)
	routerVoucher.DELETE("/:id", voucherHandle.Trashed)
	routerVoucher.PUT("/restore/:id", voucherHandle.Restore)
	routerVoucher.DELETE("/permanent/:id", voucherHandle.DeletePermanent)
	routerVoucher.PUT("/restore-all", voucherHandle.RestoreAll)
	routerVoucher.DELETE("/permanent-all", voucherHandle.DeleteAllPermanent)

	return voucherHandle
}

// @Security Bearer
// @Summary Find all Vouchers
// @Tags Voucher
// @Description Retrieve a list of all Vouchers
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationVoucher "List of Vouchers"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Voucher data"
// @Router /api/voucher [get]
func (h *voucherHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllVoucherRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.voucher.FindAll(ctx, req)
	if err != nil {
		if errors.Is(err, echo.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status:  "error",
				Message: "Unauthorized",
			})
		}

		h.logger.Debug("Failed to fetch Voucher records", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Voucher records",
		})
	}

	so := h.mapping.ToApiResponsePaginationVoucher(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find Voucher by ID
// @Tags Voucher
// @Description Retrieve a Voucher by ID
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Success 200 {object} pb.ApiResponseVoucher "Voucher data"
// @Failure 400 {object} response.ErrorResponse "Invalid Voucher ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Voucher data"
// @Router /api/voucher/{id} [get]
func (h *voucherHandleApi) FindById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Voucher ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Voucher", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Voucher",
		})
	}

	so := h.mapping.ToApiResponseVoucher(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active Vouchers
// @Tags Voucher
// @Description Retrieve a list of active Vouchers
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationVoucherDeleteAt "List of active Vouchers"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Voucher data"
// @Router /api/voucher/active [get]
func (h *voucherHandleApi) FindByActive(c echo.Context) error {
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

	req := &pb.FindAllVoucherRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.voucher.FindByActive(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch active Vouchers", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Vouchers",
		})
	}

	so := h.mapping.ToApiResponsePaginationVoucherDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve trashed Vouchers
// @Tags Voucher
// @Description Retrieve a list of trashed Voucher records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationVoucherDeleteAt "List of trashed Voucher data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Voucher data"
// @Router /api/voucher/trashed [get]
func (h *voucherHandleApi) FindByTrashed(c echo.Context) error {
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

	req := &pb.FindAllVoucherRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.voucher.FindByTrashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch trashed Vouchers", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Vouchers",
		})
	}

	so := h.mapping.ToApiResponsePaginationVoucherDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create an existing Voucher
// @Tags Voucher
// @Description Create an existing Voucher record with the provided details
// @Accept json
// @Produce json
// @Param CreateVoucherRequest body requests.CreateVoucherRequest true "Create Voucher request"
// @Success 200 {object} response.ApiResponseVoucher "Successfully created Voucher"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create Voucher"
// @Router /api/voucher/create [post]
func (h *voucherHandleApi) Create(c echo.Context) error {
	var req pb.CreateVoucherRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	ctx := c.Request().Context()

	res, err := h.voucher.Create(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to create Voucher", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create Voucher",
		})
	}

	so := h.mapping.ToApiResponseVoucher(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update an existing Voucher
// @Tags Voucher
// @Description Update an existing Voucher record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param UpdateVoucherRequest body requests.UpdateVoucherRequest true "Update Voucher request"
// @Success 200 {object} response.ApiResponseVoucher "Successfully updated Voucher"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update Voucher"
// @Router /api/voucher/update/{id} [post]
func (h *voucherHandleApi) Update(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Voucher ID",
		})
	}

	var req pb.UpdateVoucherRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	req.Id = int32(VoucherID)

	ctx := c.Request().Context()

	res, err := h.voucher.Update(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to update Voucher", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update Voucher",
		})
	}

	so := h.mapping.ToApiResponseVoucher(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve a trashed Voucher
// @Tags Voucher
// @Description Retrieve a trashed Voucher record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Success 200 {object} response.ApiResponseVoucherDeleteAt "Successfully retrieved trashed Voucher"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed Voucher"
// @Router /api/voucher/trashed/{id} [get]
func (h *voucherHandleApi) Trashed(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Voucher ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Voucher", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Voucher",
		})
	}

	so := h.mapping.ToApiResponseVoucherDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore a trashed Voucher
// @Tags Voucher
// @Description Restore a trashed Voucher record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Success 200 {object} response.ApiResponseVoucherDeleteAt "Successfully restored Voucher"
// @Failure 400 {object} response.ErrorResponse "Invalid Voucher ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Voucher"
// @Router /api/voucher/restore/{id} [post]
func (h *voucherHandleApi) Restore(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Voucher ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Voucher", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Voucher",
		})
	}

	so := h.mapping.ToApiResponseVoucherDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete a Voucher
// @Tags Voucher
// @Description Permanently delete a Voucher record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Success 200 {object} pb.ApiResponseVoucherDelete "Successfully deleted Voucher record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Voucher"
// @Router /api/voucher/delete/{id} [post]
func (h *voucherHandleApi) DeletePermanent(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Voucher ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Voucher permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Voucher permanently",
		})
	}

	so := h.mapping.ToApiResponseVoucherDelete(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore all trashed Vouchers
// @Tags Voucher
// @Description Restore all trashed Voucher records.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseVoucherAll "Successfully restored all Vouchers"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Vouchers"
// @Router /api/voucher/restore/all [post]
func (h *voucherHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.voucher.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Vouchers", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Vouchers",
		})
	}

	so := h.mapping.ToApiResponseVoucherAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete all trashed Vouchers
// @Tags Voucher
// @Description Permanently delete all trashed Voucher records.
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseVoucherDelete "Successfully deleted all Voucher records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Vouchers"
// @Router /api/voucher/delete/all [post]
func (h *voucherHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.voucher.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Vouchers permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Vouchers permanently",
		})
	}

	so := h.mapping.ToApiResponseVoucherAll(res)

	return c.JSON(http.StatusOK, so)
}
