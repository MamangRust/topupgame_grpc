package api

import (
	"net/http"
	"strconv"
	"strings"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_api "topup_game/internal/mapper/response/api"
	"topup_game/internal/pb"
	"topup_game/pkg/errors/voucher_errors"
	"topup_game/pkg/logger"
	"topup_game/pkg/upload_image"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type voucherHandleApi struct {
	voucher      pb.VoucherServiceClient
	logger       logger.LoggerInterface
	mapping      response_api.VoucherResponseMapper
	upload_image upload_image.ImageUploads
}

func NewHandlerVoucher(router *echo.Echo, voucher pb.VoucherServiceClient, logger logger.LoggerInterface, mapping response_api.VoucherResponseMapper, upload_image upload_image.ImageUploads) *voucherHandleApi {
	voucherHandle := &voucherHandleApi{
		voucher:      voucher,
		logger:       logger,
		mapping:      mapping,
		upload_image: upload_image,
	}

	routerVoucher := router.Group("/api/voucher")

	routerVoucher.GET("", voucherHandle.FindAll)
	routerVoucher.POST("/create", voucherHandle.Create)
	routerVoucher.GET("/:id", voucherHandle.FindById)
	routerVoucher.GET("/active", voucherHandle.FindByActive)
	routerVoucher.GET("/trashed", voucherHandle.FindByTrashed)

	routerVoucher.GET("/monhtly-amount-success", voucherHandle.FindMonthAmountVoucherSuccess)
	routerVoucher.GET("/yearly-amount-success", voucherHandle.FindYearAmountVoucherSuccess)
	routerVoucher.GET("/monthly-amount-failed", voucherHandle.FindMonthAmountVoucherFailed)
	routerVoucher.GET("/yearly-amount-failed", voucherHandle.FindYearAmountVoucherFailed)

	routerVoucher.GET("/monthly-method-success", voucherHandle.FindMonthMethodVoucherSuccess)
	routerVoucher.GET("/yearly-method-success", voucherHandle.FindYearMethodVoucherSuccess)
	routerVoucher.GET("/monthly-method-failed", voucherHandle.FindMonthMethodVoucherFailed)
	routerVoucher.GET("/yearly-method-failed", voucherHandle.FindYearMethodVoucherFailed)

	routerVoucher.GET("/myvoucher/monhtly-amount-success/:id", voucherHandle.FindMonthAmountVoucherSuccessById)
	routerVoucher.GET("/myvoucher/yearly-amount-success/:id", voucherHandle.FindYearAmountVoucherSuccessById)
	routerVoucher.GET("/myvoucher/monthly-amount-failed/:id", voucherHandle.FindMonthAmountVoucherFailedById)
	routerVoucher.GET("/myvoucher/yearly-amount-failed/:id", voucherHandle.FindYearAmountVoucherFailedById)

	routerVoucher.GET("/myvoucher/monthly-method-success/:id", voucherHandle.FindMonthMethodVoucherSuccessById)
	routerVoucher.GET("/myvoucher/yearly-method-success/:id", voucherHandle.FindYearMethodVoucherSuccessById)
	routerVoucher.GET("/myvoucher/monthly-method-failed/:id", voucherHandle.FindMonthMethodVoucherSuccessById)
	routerVoucher.GET("/myvoucher/yearly-method-failed/:id", voucherHandle.FindYearMethodVoucherSuccessById)

	routerVoucher.GET("/merchant/monhtly-amount-success/:merchant_id", voucherHandle.FindMonthAmountVoucherSuccessByMerchant)
	routerVoucher.GET("/merchant/yearly-amount-success/:merchant_id", voucherHandle.FindYearAmountVoucherSuccessByMerchant)
	routerVoucher.GET("/merchant/monthly-amount-failed/:merchant_id", voucherHandle.FindMonthAmountVoucherFailedByMerchant)
	routerVoucher.GET("/merchant/yearly-amount-failed/:merchant_id", voucherHandle.FindYearAmountVoucherFailedByMerchant)

	routerVoucher.GET("/merchant/monthly-method-success/:merchant_id", voucherHandle.FindMonthMethodVoucherSuccessByMerchant)
	routerVoucher.GET("/merchant/yearly-method-success/:merchant_id", voucherHandle.FindYearMethodVoucherSuccessByMerchant)
	routerVoucher.GET("/merchant/monthly-method-failed/:merchant_id", voucherHandle.FindMonthMethodVoucherSuccessByMerchant)
	routerVoucher.GET("/merchant/yearly-method-failed/:merchant_id", voucherHandle.FindYearMethodVoucherSuccessByMerchant)

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
		h.logger.Debug("Failed to fetch Voucher records", zap.Error(err))
		return voucher_errors.ErrApiFailedFindAll(c)
	}

	so := h.mapping.ToApiResponsePaginationVoucher(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find monthly voucher success amounts
// @Tags Voucher
// @Description Retrieve monthly voucher success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesVoucherMonthSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/monhtly-amount-success [get]
func (h *voucherHandleApi) FindMonthAmountVoucherSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountVoucherRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.voucher.FindMonthAmountVoucherSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly voucher success amounts", zap.Error(err))
		return voucher_errors.ErrApiFindMonthAmountVoucherSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly voucher success amounts
// @Tags Voucher
// @Description Retrieve yearly voucher success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/yearly-amount-success [get]
func (h *voucherHandleApi) FindYearAmountVoucherSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountVoucherRequest{
		Year: int32(year),
	}

	res, err := h.voucher.FindYearAmountVoucherSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly voucher success amounts", zap.Error(err))
		return voucher_errors.ErrApiFindYearAmountVoucherSuccess(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly voucher failed amounts
// @Tags Voucher
// @Description Retrieve monthly voucher failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesVoucherMonthFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/monthly-amount-failed [get]
func (h *voucherHandleApi) FindMonthAmountVoucherFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountVoucherRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.voucher.FindMonthAmountVoucherFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly voucher failed amounts", zap.Error(err))
		return voucher_errors.ErrApiFindMonthAmountVoucherFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly voucher failed amounts
// @Tags Voucher
// @Description Retrieve yearly voucher failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/yearly-amount-failed [get]
func (h *voucherHandleApi) FindYearAmountVoucherFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountVoucherRequest{
		Year: int32(year),
	}

	res, err := h.voucher.FindYearAmountVoucherFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly voucher failed amounts", zap.Error(err))
		return voucher_errors.ErrApiFindYearAmountVoucherFailed(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly voucher success methods
// @Tags Voucher
// @Description Retrieve monthly voucher success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/monhtly-method-success [get]
func (h *voucherHandleApi) FindMonthMethodVoucherSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountVoucherRequest{
		Year: int32(year),
	}

	res, err := h.voucher.FindMonthMethodVoucherSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly voucher success methods", zap.Error(err))
		return voucher_errors.ErrApiFindMonthMethodVoucherSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly voucher success methods
// @Tags Voucher
// @Description Retrieve yearly voucher success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/yearly-method-success [get]
func (h *voucherHandleApi) FindYearMethodVoucherSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountVoucherRequest{
		Year: int32(year),
	}

	res, err := h.voucher.FindYearMethodVoucherSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly voucher success methods", zap.Error(err))
		return voucher_errors.ErrApiFindYearMethodVoucherSuccess(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly voucher failed methods
// @Tags Voucher
// @Description Retrieve monthly voucher failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/monthly-method-failed [get]
func (h *voucherHandleApi) FindMonthMethodVoucherFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountVoucherRequest{
		Year: int32(year),
	}

	res, err := h.voucher.FindMonthMethodVoucherFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly voucher failed methods", zap.Error(err))
		return voucher_errors.ErrApiFindMonthMethodVoucherFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly voucher failed methods
// @Tags Voucher
// @Description Retrieve yearly voucher failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/yearly-method-failed [get]
func (h *voucherHandleApi) FindYearMethodVoucherFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountVoucherRequest{
		Year: int32(year),
	}

	res, err := h.voucher.FindYearMethodVoucherFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly voucher failed methods", zap.Error(err))
		return voucher_errors.ErrApiFindYearMethodVoucherFailed(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Get monthly success amounts by Voucher ID
// @Tags Voucher
// @Description Retrieve monthly success amounts for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponsesVoucherMonthSuccess "Monthly success amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/monhtly-amount-success/{id} [get]
func (h *voucherHandleApi) FindMonthAmountVoucherSuccessById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthAmountVoucherByIdRequest{
		Id:    int32(VoucherID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.voucher.FindMonthAmountVoucherSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher success amounts", zap.Error(err))
		return voucher_errors.ErrApiFindMonthAmountVoucherSuccessById(c)
	}

	so := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly success amounts by Voucher ID
// @Tags Voucher
// @Description Retrieve yearly success amounts for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearSuccess "Yearly success amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/yearly-amount-success/{id} [get]
func (h *voucherHandleApi) FindYearAmountVoucherSuccessById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearAmountVoucherByIdRequest{
		Id:   int32(VoucherID),
		Year: int32(year),
	}

	res, err := h.voucher.FindYearAmountVoucherSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher success amounts", zap.Error(err))
		return voucher_errors.ErrApiFindYearAmountVoucherSuccessById(c)
	}

	so := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get monthly failed amounts by Voucher ID
// @Tags Voucher
// @Description Retrieve monthly failed amounts for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponsesVoucherMonthFailed "Monthly failed amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/monthly-amount-failed/{id} [get]
func (h *voucherHandleApi) FindMonthAmountVoucherFailedById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthAmountVoucherByIdRequest{
		Id:    int32(VoucherID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.voucher.FindMonthAmountVoucherFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher failed amounts", zap.Error(err))
		return voucher_errors.ErrApiFindMonthAmountVoucherFailedById(c)
	}

	so := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly failed amounts by Voucher ID
// @Tags Voucher
// @Description Retrieve yearly failed amounts for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearFailed "Yearly failed amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/yearly-amount-failed/{id} [get]
func (h *voucherHandleApi) FindYearAmountVoucherFailedById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearAmountVoucherByIdRequest{
		Id:   int32(VoucherID),
		Year: int32(year),
	}

	res, err := h.voucher.FindYearAmountVoucherFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher failed amounts", zap.Error(err))
		return voucher_errors.ErrApiFindYearAmountVoucherFailedById(c)
	}

	so := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get monthly success methods by Voucher ID
// @Tags Voucher
// @Description Retrieve monthly success methods for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherMonthMethod "Monthly success methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/monthly-method-success/{id} [get]
func (h *voucherHandleApi) FindMonthMethodVoucherSuccessById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthMethodVoucherByIdRequest{
		Id:   int32(VoucherID),
		Year: int32(year),
	}

	res, err := h.voucher.FindMonthMethodVoucherSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher success methods", zap.Error(err))
		return voucher_errors.ErrApiFindMonthMethodVoucherSuccessById(c)
	}

	so := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly success methods by Voucher ID
// @Tags Voucher
// @Description Retrieve yearly success methods for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearMethod "Yearly success methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/yearly-method-success/{id} [get]
func (h *voucherHandleApi) FindYearMethodVoucherSuccessById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearMethodVoucherByIdRequest{
		Id:   int32(VoucherID),
		Year: int32(year),
	}

	res, err := h.voucher.FindYearMethodVoucherSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher success methods", zap.Error(err))
		return voucher_errors.ErrApiFindYearMethodVoucherSuccessById(c)
	}

	so := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get monthly failed methods by Voucher ID
// @Tags Voucher
// @Description Retrieve monthly failed methods for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherMonthMethod "Monthly failed methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/monthly-method-failed/{id} [get]
func (h *voucherHandleApi) FindMonthMethodVoucherFailedById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthMethodVoucherByIdRequest{
		Id:   int32(VoucherID),
		Year: int32(year),
	}

	res, err := h.voucher.FindMonthMethodVoucherFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher failed methods", zap.Error(err))
		return voucher_errors.ErrApiFindMonthMethodVoucherFailedById(c)
	}

	so := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly failed methods by Voucher ID
// @Tags Voucher
// @Description Retrieve yearly failed methods for a specific Voucher
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearMethod "Yearly failed methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/myvoucher/yearly-method-failed/{id} [get]
func (h *voucherHandleApi) FindYearMethodVoucherFailedById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearMethodVoucherByIdRequest{
		Id:   int32(VoucherID),
		Year: int32(year),
	}

	res, err := h.voucher.FindYearMethodVoucherFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher failed methods", zap.Error(err))
		return voucher_errors.ErrApiFindYearMethodVoucherFailedById(c)
	}

	so := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get monthly success amounts by Merchant ID
// @Tags Voucher
// @Description Retrieve monthly success amounts for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponsesVoucherMonthSuccess "Monthly success amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/monthly-amount-success/{merchant_id} [get]
func (h *voucherHandleApi) FindMonthAmountVoucherSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthAmountVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.voucher.FindMonthAmountVoucherSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher success amounts by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindMonthAmountVoucherSuccessByMerchant(c)
	}

	so := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly success amounts by Merchant ID
// @Tags Voucher
// @Description Retrieve yearly success amounts for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearSuccess "Yearly success amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/yearly-amount-success/{merchant_id} [get]
func (h *voucherHandleApi) FindYearAmountVoucherSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearAmountVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.voucher.FindYearAmountVoucherSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher success amounts by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindYearAmountVoucherSuccessByMerchant(c)
	}

	so := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get monthly failed amounts by Merchant ID
// @Tags Voucher
// @Description Retrieve monthly failed amounts for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponsesVoucherMonthFailed "Monthly failed amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/monthly-amount-failed/{merchant_id} [get]
func (h *voucherHandleApi) FindMonthAmountVoucherFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthAmountVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.voucher.FindMonthAmountVoucherFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher failed amounts by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindMonthAmountVoucherFailedByMerchant(c)
	}

	so := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly failed amounts by Merchant ID
// @Tags Voucher
// @Description Retrieve yearly failed amounts for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearFailed "Yearly failed amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/yearly-amount-failed/{merchant_id} [get]
func (h *voucherHandleApi) FindYearAmountVoucherFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearAmountVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.voucher.FindYearAmountVoucherFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher failed amounts by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindYearAmountVoucherFailedByMerchant(c)
	}

	so := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get monthly success methods by Merchant ID
// @Tags Voucher
// @Description Retrieve monthly success methods for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherMonthMethod "Monthly success methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/monthly-method-success/{merchant_id} [get]
func (h *voucherHandleApi) FindMonthMethodVoucherSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthMethodVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.voucher.FindMonthMethodVoucherSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher success methods by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindMonthMethodVoucherSuccessByMerchant(c)
	}

	so := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly success methods by Merchant ID
// @Tags Voucher
// @Description Retrieve yearly success methods for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearMethod "Yearly success methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/yearly-method-success/{merchant_id} [get]
func (h *voucherHandleApi) FindYearMethodVoucherSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearMethodVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.voucher.FindYearMethodVoucherSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher success methods by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindYearMethodVoucherSuccessByMerchant(c)
	}

	so := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get monthly failed methods by Merchant ID
// @Tags Voucher
// @Description Retrieve monthly failed methods for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherMonthMethod "Monthly failed methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/monthly-method-failed/{merchant_id} [get]
func (h *voucherHandleApi) FindMonthMethodVoucherFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.MonthMethodVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.voucher.FindMonthMethodVoucherFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly Voucher failed methods by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindMonthMethodVoucherFailedByMerchant(c)
	}

	so := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Get yearly failed methods by Merchant ID
// @Tags Voucher
// @Description Retrieve yearly failed methods for a specific Merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesVoucherYearMethod "Yearly failed methods"
// @Failure 400 {object} response.ErrorResponse "Invalid parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/voucher/merchant/yearly-method-failed/{merchant_id} [get]
func (h *voucherHandleApi) FindYearMethodVoucherFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year parameter", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	req := &pb.YearMethodVoucherByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.voucher.FindYearMethodVoucherFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly Voucher failed methods by merchant", zap.Error(err))
		return voucher_errors.ErrApiFindYearMethodVoucherFailedByMerchant(c)
	}

	so := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find Voucher by ID
// @Tags Voucher
// @Description Retrieve a Voucher by ID
// @Accept json
// @Produce json
// @Param id path int true "Voucher ID"
// @Success 200 {object} response.ApiResponseVoucher "Voucher data"
// @Failure 400 {object} response.ErrorResponse "Invalid Voucher ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Voucher data"
// @Router /api/voucher/{id} [get]
func (h *voucherHandleApi) FindById(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Voucher", zap.Error(err))
		return voucher_errors.ErrApiVoucherNotFound(c)
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
		return voucher_errors.ErrApiFailedFindActive(c)
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
		return voucher_errors.ErrApiFailedFindTrashed(c)
	}

	so := h.mapping.ToApiResponsePaginationVoucherDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create an existing Voucher
// @Tags Voucher
// @Description Create an existing Voucher record with the provided details
// @Accept multipart/form-data
// @Produce json
// @Param merchant_id formData string true "Merchant ID"
// @Param category_id formData string true "Category ID"
// @Param name formData string true "Product name"
// @Param image_path formData file true "Product image file"
// @Success 200 {object} response.ApiResponseVoucher "Successfully created Voucher"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create Voucher"
// @Router /api/voucher/create [post]
func (h *voucherHandleApi) Create(c echo.Context) error {
	formData, err := h.parseVoucherForm(c, true)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "invalid body",
			Code:    http.StatusBadRequest,
		})
	}

	reqPb := &pb.CreateVoucherRequest{
		MerchantId: int32(formData.MerchantID),
		CategoryId: int32(formData.CategoryID),
		Name:       formData.Name,
		ImageName:  formData.ImagePath,
	}

	ctx := c.Request().Context()

	res, err := h.voucher.Create(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to create Voucher", zap.Error(err))
		return voucher_errors.ErrApiFailedCreateVoucher(c)
	}

	so := h.mapping.ToApiResponseVoucher(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update an existing Voucher
// @Tags Voucher
// @Description Update an existing Voucher record with the provided details
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Voucher ID"
// @Param merchant_id formData string true "Merchant ID"
// @Param category_id formData string true "Category ID"
// @Param name formData string true "Product name"
// @Param image_path formData file true "Product image file"
// @Success 200 {object} response.ApiResponseVoucher "Successfully updated Voucher"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update Voucher"
// @Router /api/voucher/update/{id} [post]
func (h *voucherHandleApi) Update(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	formData, err := h.parseVoucherForm(c, false)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "invalid body",
			Code:    http.StatusBadRequest,
		})
	}

	reqPb := &pb.UpdateVoucherRequest{
		Id:         int32(VoucherID),
		MerchantId: int32(formData.MerchantID),
		CategoryId: int32(formData.CategoryID),
		Name:       formData.Name,
		ImageName:  formData.ImagePath,
	}

	ctx := c.Request().Context()

	res, err := h.voucher.Update(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to update Voucher", zap.Error(err))
		return voucher_errors.ErrApiFailedUpdateVoucher(c)
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
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Voucher", zap.Error(err))
		return voucher_errors.ErrApiFailedTrashedVoucher(c)
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
		return voucher_errors.ErrInvalidVoucherId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Voucher", zap.Error(err))
		return voucher_errors.ErrApiFailedRestoreVoucher(c)
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
// @Success 200 {object} response.ApiResponseVoucherDelete "Successfully deleted Voucher record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Voucher"
// @Router /api/voucher/delete/{id} [post]
func (h *voucherHandleApi) DeletePermanent(c echo.Context) error {
	VoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil || VoucherID <= 0 {
		return voucher_errors.ErrApiVoucherInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdVoucherRequest{
		VoucherId: int32(VoucherID),
	}

	res, err := h.voucher.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Voucher permanently", zap.Error(err))
		return voucher_errors.ErrApiFailedDeletePermanent(c)
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
// @Success 200 {object} response.ApiResponseVoucherAll "Successfully restored all Vouchers"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Vouchers"
// @Router /api/voucher/restore/all [post]
func (h *voucherHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.voucher.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Vouchers", zap.Error(err))
		return voucher_errors.ErrApiFailedRestoreAll(c)
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
// @Success 200 {object} response.ApiResponseVoucherDelete "Successfully deleted all Voucher records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Vouchers"
// @Router /api/voucher/delete/all [post]
func (h *voucherHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.voucher.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Vouchers permanently", zap.Error(err))
		return voucher_errors.ErrApiFailedDeleteAll(c)
	}

	so := h.mapping.ToApiResponseVoucherAll(res)

	return c.JSON(http.StatusOK, so)
}

func (h *voucherHandleApi) parseVoucherForm(c echo.Context, requireImage bool) (requests.VoucherFormData, error) {
	var formData requests.VoucherFormData
	var err error

	formData.MerchantID, err = strconv.Atoi(c.FormValue("merchant_id"))
	if err != nil || formData.MerchantID <= 0 {
		return formData, c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "invalid_merchant",
			Message: "Please provide a valid merchant ID",
			Code:    http.StatusBadRequest,
		})
	}

	formData.CategoryID, err = strconv.Atoi(c.FormValue("category_id"))
	if err != nil || formData.CategoryID <= 0 {
		return formData, c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "invalid_category",
			Message: "Please provide a valid category ID",
			Code:    http.StatusBadRequest,
		})
	}

	formData.Name = strings.TrimSpace(c.FormValue("name"))
	if formData.Name == "" {
		return formData, c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "validation_error",
			Message: "Voucher name is required",
			Code:    http.StatusBadRequest,
		})
	}

	file, err := c.FormFile("image_voucher")
	if err != nil {
		if requireImage {
			h.logger.Debug("Image upload error", zap.Error(err))
			return formData, c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status:  "image_required",
				Message: "A voucher image is required",
				Code:    http.StatusBadRequest,
			})
		}

		return formData, nil
	}

	imagePath, err := h.upload_image.ProcessImageUpload(c, file)
	if err != nil {
		return formData, err
	}

	formData.ImagePath = imagePath
	return formData, nil
}
