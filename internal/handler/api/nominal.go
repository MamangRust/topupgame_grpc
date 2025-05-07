package api

import (
	"net/http"
	"strconv"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_api "topup_game/internal/mapper/response/api"
	"topup_game/internal/pb"
	"topup_game/pkg/errors/nominal_errors"
	"topup_game/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type nominalHandleApi struct {
	nominal pb.NominalServiceClient
	logger  logger.LoggerInterface
	mapping response_api.NominalResponseMapper
}

func NewHandlerNominal(router *echo.Echo, nominal pb.NominalServiceClient, logger logger.LoggerInterface, mapping response_api.NominalResponseMapper) *nominalHandleApi {
	nominalHandler := &nominalHandleApi{
		nominal: nominal,
		logger:  logger,
		mapping: mapping,
	}

	routerNominal := router.Group("/api/nominal")

	routerNominal.GET("", nominalHandler.FindAll)
	routerNominal.GET("/:id", nominalHandler.FindById)
	routerNominal.GET("/active", nominalHandler.FindByActive)
	routerNominal.GET("/trashed", nominalHandler.FindByTrashed)
	routerNominal.POST("/create", nominalHandler.Create)

	routerNominal.GET("/monthly-amount-success", nominalHandler.FindMonthAmountNominalSuccess)
	routerNominal.GET("/yearly-amount-success", nominalHandler.FindYearAmountNominalSuccess)
	routerNominal.GET("/monthly-amount-failed", nominalHandler.FindMonthAmountNominalFailed)
	routerNominal.GET("/yearly-amount-failed", nominalHandler.FindYearAmountNominalFailed)

	routerNominal.GET("/monthly-method-success", nominalHandler.FindMonthMethodNominalSuccess)
	routerNominal.GET("/yearly-method-success", nominalHandler.FindYearMethodNominalSuccess)
	routerNominal.GET("/monthly-method-failed", nominalHandler.FindMonthMethodNominalSuccess)
	routerNominal.GET("/yearly-method-failed", nominalHandler.FindYearMethodNominalSuccess)

	routerNominal.GET("/mynominal/monthly-amount-success/:id", nominalHandler.FindMonthAmountNominalSuccessById)
	routerNominal.GET("/mynominal/yearly-amount-success/:id", nominalHandler.FindYearAmountNominalSuccessById)
	routerNominal.GET("/mynominal/monthly-amount-failed/:id", nominalHandler.FindMonthAmountNominalFailedById)
	routerNominal.GET("/mynominal/yearly-amount-failed/:id", nominalHandler.FindYearAmountNominalFailedById)

	routerNominal.GET("/mynominal/monthly-method-success/:id", nominalHandler.FindMonthMethodNominalSuccessById)
	routerNominal.GET("/mynominal/yearly-method-success/:id", nominalHandler.FindYearMethodNominalSuccessById)
	routerNominal.GET("/mynominal/monthly-method-failed/:id", nominalHandler.FindMonthMethodNominalSuccessById)
	routerNominal.GET("/mynominal/yearly-method-failed/:id", nominalHandler.FindYearMethodNominalSuccessById)

	routerNominal.GET("/merchant/monthly-amount-success/:merchant_id", nominalHandler.FindMonthAmountNominalSuccessByMerchant)
	routerNominal.GET("/merchant/yearly-amount-success/:merchant_id", nominalHandler.FindYearAmountNominalSuccessByMerchant)
	routerNominal.GET("/merchant/monthly-amount-failed/:merchant_id", nominalHandler.FindMonthAmountNominalFailedByMerchant)
	routerNominal.GET("/merchant/yearly-amount-failed/:merchant_id", nominalHandler.FindYearAmountNominalFailedByMerchant)

	routerNominal.GET("/merchant/monthly-method-success/:merchant_id", nominalHandler.FindMonthMethodNominalSuccessByMerchant)
	routerNominal.GET("/merchant/yearly-method-success/:merchant_id", nominalHandler.FindYearMethodNominalSuccessByMerchant)
	routerNominal.GET("/merchant/monthly-method-failed/:merchant_id", nominalHandler.FindMonthMethodNominalSuccessByMerchant)
	routerNominal.GET("/merchant/yearly-method-failed/:merchant_id", nominalHandler.FindYearMethodNominalSuccessByMerchant)

	routerNominal.POST("/update/:id", nominalHandler.Update)
	routerNominal.POST("/trashed/:id", nominalHandler.Trashed)
	routerNominal.POST("/restore/:id", nominalHandler.Restore)
	routerNominal.DELETE("/permanent/:id", nominalHandler.DeletePermanent)
	routerNominal.POST("/restore/all", nominalHandler.RestoreAll)
	routerNominal.DELETE("/permanent/all", nominalHandler.DeleteAllPermanent)

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
func (h *nominalHandleApi) FindAll(c echo.Context) error {
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
		h.logger.Debug("Failed to fetch Nominal records", zap.Error(err))
		return nominal_errors.ErrApiFailedFindAll(c)
	}

	so := h.mapping.ToApiResponsePaginationNominal(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find monthly nominal success amounts
// @Tags Nominal
// @Description Retrieve monthly nominal success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseNominalMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/monthly/success/amount [get]
func (h *nominalHandleApi) FindMonthAmountNominalSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountNominalRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.nominal.FindMonthAmountNominalSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal success amounts", zap.Error(err))
		return nominal_errors.ErrApiFindMonthAmountNominalSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal success amounts
// @Tags Nominal
// @Description Retrieve yearly nominal success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/yearly/success/amount [get]
func (h *nominalHandleApi) FindYearAmountNominalSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalRequest{
		Year: int32(year),
	}

	res, err := h.nominal.FindYearAmountNominalSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal success amounts", zap.Error(err))
		return nominal_errors.ErrApiFindYearAmountNominalSuccess(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal failed amounts
// @Tags Nominal
// @Description Retrieve monthly nominal failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseNominalMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/monthly/failed/amount [get]
func (h *nominalHandleApi) FindMonthAmountNominalFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountNominalRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.nominal.FindMonthAmountNominalFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal failed amounts", zap.Error(err))
		return nominal_errors.ErrApiFindMonthAmountNominalFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal failed amounts
// @Tags Nominal
// @Description Retrieve yearly nominal failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/yearly/failed/amount [get]
func (h *nominalHandleApi) FindYearAmountNominalFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalRequest{
		Year: int32(year),
	}

	res, err := h.nominal.FindYearAmountNominalFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal failed amounts", zap.Error(err))
		return nominal_errors.ErrApiFindYearAmountNominalFailed(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal success methods
// @Tags Nominal
// @Description Retrieve monthly nominal success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/monthly/success/method [get]
func (h *nominalHandleApi) FindMonthMethodNominalSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalRequest{
		Year: int32(year),
	}

	res, err := h.nominal.FindMonthMethodNominalSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal success methods", zap.Error(err))
		return nominal_errors.ErrApiFindMonthMethodNominalSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal success methods
// @Tags Nominal
// @Description Retrieve yearly nominal success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/yearly/success/method [get]
func (h *nominalHandleApi) FindYearMethodNominalSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalRequest{
		Year: int32(year),
	}

	res, err := h.nominal.FindYearMethodNominalSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal success methods", zap.Error(err))
		return nominal_errors.ErrApiFindYearMethodNominalSuccess(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal failed methods
// @Tags Nominal
// @Description Retrieve monthly nominal failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/monthly/failed/method [get]
func (h *nominalHandleApi) FindMonthMethodNominalFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalRequest{
		Year: int32(year),
	}

	res, err := h.nominal.FindMonthMethodNominalFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal failed methods", zap.Error(err))
		return nominal_errors.ErrApiFindMonthMethodNominalFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal failed methods
// @Tags Nominal
// @Description Retrieve yearly nominal failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/yearly/failed/method [get]
func (h *nominalHandleApi) FindYearMethodNominalFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalRequest{
		Year: int32(year),
	}

	res, err := h.nominal.FindYearMethodNominalFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal failed methods", zap.Error(err))
		return nominal_errors.ErrApiFindYearMethodNominalFailed(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal success amounts by ID
// @Tags Nominal
// @Description Retrieve monthly nominal success amounts by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseNominalMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/monthly-amount-success/{id} [get]
func (h *nominalHandleApi) FindMonthAmountNominalSuccessById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountNominalByIdRequest{
		Id:    int32(nominalID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.nominal.FindMonthAmountNominalSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal success amounts by ID", zap.Error(err))
		return nominal_errors.ErrApiFindMonthAmountNominalSuccessById(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal success amounts by ID
// @Tags Nominal
// @Description Retrieve yearly nominal success amounts by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/yearly-amount-nominal/{id} [get]
func (h *nominalHandleApi) FindYearAmountNominalSuccessById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalByIdRequest{
		Id:   int32(nominalID),
		Year: int32(year),
	}

	res, err := h.nominal.FindYearAmountNominalSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal success amounts by ID", zap.Error(err))
		return nominal_errors.ErrApiFindYearAmountNominalSuccessById(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal failed amounts by ID
// @Tags Nominal
// @Description Retrieve monthly nominal failed amounts by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseNominalMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/monthly-amount-failed/{id} [get]
func (h *nominalHandleApi) FindMonthAmountNominalFailedById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountNominalByIdRequest{
		Id:    int32(nominalID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.nominal.FindMonthAmountNominalFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal failed amounts by ID", zap.Error(err))
		return nominal_errors.ErrApiFindMonthAmountNominalFailedById(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal failed amounts by ID
// @Tags Nominal
// @Description Retrieve yearly nominal failed amounts by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/yearly-amount-failed/{id} [get]
func (h *nominalHandleApi) FindYearAmountNominalFailedById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalByIdRequest{
		Id:   int32(nominalID),
		Year: int32(year),
	}

	res, err := h.nominal.FindYearAmountNominalFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal failed amounts by ID", zap.Error(err))
		return nominal_errors.ErrApiFindYearAmountNominalFailedById(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal success methods by ID
// @Tags Nominal
// @Description Retrieve monthly nominal success methods by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/monthly-method-success/{id} [get]
func (h *nominalHandleApi) FindMonthMethodNominalSuccessById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodNominalByIdRequest{
		Id:   int32(nominalID),
		Year: int32(year),
	}

	res, err := h.nominal.FindMonthMethodNominalSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal success methods by ID", zap.Error(err))
		return nominal_errors.ErrApiFindMonthMethodNominalSuccessById(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal success methods by ID
// @Tags Nominal
// @Description Retrieve yearly nominal success methods by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/yearly-method-success/{id} [get]
func (h *nominalHandleApi) FindYearMethodNominalSuccessById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodNominalByIdRequest{
		Id:   int32(nominalID),
		Year: int32(year),
	}

	res, err := h.nominal.FindYearMethodNominalSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal success methods by ID", zap.Error(err))
		return nominal_errors.ErrApiFindYearMethodNominalSuccessById(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal failed methods by ID
// @Tags Nominal
// @Description Retrieve monthly nominal failed methods by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/monthly-method-failed/{id} [get]
func (h *nominalHandleApi) FindMonthMethodNominalFailedById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodNominalByIdRequest{
		Id:   int32(nominalID),
		Year: int32(year),
	}

	res, err := h.nominal.FindMonthMethodNominalFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal failed methods by ID", zap.Error(err))
		return nominal_errors.ErrApiFindMonthMethodNominalFailedById(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal failed methods by ID
// @Tags Nominal
// @Description Retrieve yearly nominal failed methods by nominal ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid nominal ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/mynominal/yearly-method-failed/{id} [get]
func (h *nominalHandleApi) FindYearMethodNominalFailedById(c echo.Context) error {
	nominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || nominalID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid nominal ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodNominalByIdRequest{
		Id:   int32(nominalID),
		Year: int32(year),
	}

	res, err := h.nominal.FindYearMethodNominalFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal failed methods by ID", zap.Error(err))
		return nominal_errors.ErrApiFindYearMethodNominalFailedById(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal success amounts by merchant
// @Tags Nominal
// @Description Retrieve monthly nominal success amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseNominalMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/monthly-amount-success/{merchant_id} [get]
func (h *nominalHandleApi) FindMonthAmountNominalSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.nominal.FindMonthAmountNominalSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal success amounts by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindMonthAmountNominalSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal success amounts by merchant
// @Tags Nominal
// @Description Retrieve yearly nominal success amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/yearly-amount-success/{merchant_id} [get]
func (h *nominalHandleApi) FindYearAmountNominalSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.nominal.FindYearAmountNominalSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal success amounts by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindYearAmountNominalSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal failed amounts by merchant
// @Tags Nominal
// @Description Retrieve monthly nominal failed amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseNominalMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/monthly-amount-failed/{merchant_id} [get]
func (h *nominalHandleApi) FindMonthAmountNominalFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.nominal.FindMonthAmountNominalFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal failed amounts by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindMonthAmountNominalFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal failed amounts by merchant
// @Tags Nominal
// @Description Retrieve yearly nominal failed amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/yearly-amount-failed/{merchant_is} [get]
func (h *nominalHandleApi) FindYearAmountNominalFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.nominal.FindYearAmountNominalFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal failed amounts by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindYearAmountNominalFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal success methods by merchant
// @Tags Nominal
// @Description Retrieve monthly nominal success methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/monthly-method-success/{merchant_id} [get]
func (h *nominalHandleApi) FindMonthMethodNominalSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.nominal.FindMonthMethodNominalSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal success methods by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindMonthMethodNominalSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal success methods by merchant
// @Tags Nominal
// @Description Retrieve yearly nominal success methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/yearly-method-success/{merchant_id} [get]
func (h *nominalHandleApi) FindYearMethodNominalSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.nominal.FindYearMethodNominalSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal success methods by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindYearMethodNominalSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly nominal failed methods by merchant
// @Tags Nominal
// @Description Retrieve monthly nominal failed methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/monthly-method-failed/{merchant_id} [get]
func (h *nominalHandleApi) FindMonthMethodNominalFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.nominal.FindMonthMethodNominalFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly nominal failed methods by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindMonthMethodNominalFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly nominal failed methods by merchant
// @Tags Nominal
// @Description Retrieve yearly nominal failed methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseNominalYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/nominal/merchant/yearly-method-failed/{merchant_id} [get]
func (h *nominalHandleApi) FindYearMethodNominalFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodNominalByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.nominal.FindYearMethodNominalFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly nominal failed methods by merchant", zap.Error(err))
		return nominal_errors.ErrApiFindYearMethodNominalFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find Nominal by ID
// @Tags Nominal
// @Description Retrieve a Nominal by ID
// @Accept json
// @Produce json
// @Param id path int true "Nominal ID"
// @Success 200 {object} response.ApiResponseNominal "Nominal data"
// @Failure 400 {object} response.ErrorResponse "Invalid Nominal ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Nominal data"
// @Router /api/nominal/{id} [get]
func (h *nominalHandleApi) FindById(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return nominal_errors.ErrInvalidNominalId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Nominal", zap.Error(err))
		return nominal_errors.ErrApiNominalNotFound(c)
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
func (h *nominalHandleApi) FindByActive(c echo.Context) error {
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
		return nominal_errors.ErrApiFailedFindActive(c)
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
func (h *nominalHandleApi) FindByTrashed(c echo.Context) error {
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
		return nominal_errors.ErrApiFailedFindTrashed(c)
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
func (h *nominalHandleApi) Create(c echo.Context) error {
	var req requests.CreateNominalRequest

	if err := c.Bind(&req); err != nil {
		return nominal_errors.ErrApiBindCreateNominal(c)
	}

	if err := req.Validate(); err != nil {
		return nominal_errors.ErrApiValidateCreateNominal(c)
	}

	ctx := c.Request().Context()

	reqPb := &pb.CreateNominalRequest{
		VoucherId: int32(req.VoucherID),
		Name:      req.Name,
		Quantity:  int32(req.Quantity),
		Price:     req.Price,
	}

	res, err := h.nominal.Create(ctx, reqPb)

	if err != nil {
		h.logger.Debug("Failed to create Nominal", zap.Error(err))
		return nominal_errors.ErrApiFailedCreateNominal(c)
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
func (h *nominalHandleApi) Update(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return nominal_errors.ErrApiNominalInvalidId(c)
	}

	var req requests.UpdateNominalRequest
	if err := c.Bind(&req); err != nil {
		return nominal_errors.ErrApiBindUpdateNominal(c)
	}

	if err := req.Validate(); err != nil {
		return nominal_errors.ErrApiValidateUpdateNominal(c)
	}

	reqPb := &pb.UpdateNominalRequest{
		Id:        int32(NominalID),
		VoucherId: int32(req.VoucherID),
		Name:      req.Name,
		Quantity:  int32(req.Quantity),
		Price:     req.Price,
	}

	ctx := c.Request().Context()

	res, err := h.nominal.Update(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to update Nominal", zap.Error(err))
		return nominal_errors.ErrApiFailedUpdateNominal(c)
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
// @Router /api/nominal/trashed/{id} [post]
func (h *nominalHandleApi) Trashed(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return nominal_errors.ErrApiNominalInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Nominal", zap.Error(err))
		return nominal_errors.ErrApiFailedTrashedNominal(c)
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
func (h *nominalHandleApi) Restore(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return nominal_errors.ErrApiNominalInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Nominal", zap.Error(err))
		return nominal_errors.ErrApiFailedRestoreNominal(c)
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
// @Success 200 {object} response.ApiResponseNominalDelete "Successfully deleted Nominal record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Nominal"
// @Router /api/nominal/delete/{id} [post]
func (h *nominalHandleApi) DeletePermanent(c echo.Context) error {
	NominalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || NominalID <= 0 {
		return nominal_errors.ErrApiNominalInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdNominalRequest{
		NominalId: int32(NominalID),
	}

	res, err := h.nominal.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Nominal permanently", zap.Error(err))
		return nominal_errors.ErrApiFailedDeletePermanent(c)
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
// @Success 200 {object} response.ApiResponseNominalAll "Successfully restored all Nominals"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Nominals"
// @Router /api/nominal/restore/all [post]
func (h *nominalHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.nominal.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Nominals", zap.Error(err))
		return nominal_errors.ErrApiFailedRestoreAll(c)
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
// @Success 200 {object} response.ApiResponseNominalDelete "Successfully deleted all Nominal records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Nominals"
// @Router /api/nominal/delete/all [post]
func (h *nominalHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.nominal.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Nominals permanently", zap.Error(err))
		return nominal_errors.ErrApiFailedDeleteAll(c)
	}

	so := h.mapping.ToApiResponseNominalAll(res)

	return c.JSON(http.StatusOK, so)
}
