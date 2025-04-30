package api

import (
	"net/http"
	"strconv"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_api "topup_game/internal/mapper/response/api"
	"topup_game/internal/pb"
	"topup_game/pkg/errors/bank_errors"
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
	routerBank.POST("/create", bankHandler.Create)
	routerBank.POST("/update/:id", bankHandler.Update)
	routerBank.DELETE("/delete/:id", bankHandler.Trashed)
	routerBank.PUT("/restore/:id", bankHandler.Restore)
	routerBank.DELETE("/permanent/:id", bankHandler.DeletePermanent)
	routerBank.POST("/restore/all", bankHandler.RestoreAll)
	routerBank.DELETE("/permanent/all", bankHandler.DeleteAllPermanent)

	routerBank.GET("/monthly-amount-success", bankHandler.FindMonthAmountBankSuccess)
	routerBank.GET("/yearly-amount-success", bankHandler.FindYearAmountBankSuccess)
	routerBank.GET("/monthly-amount-failed", bankHandler.FindMonthAmountBankFailed)
	routerBank.GET("/yearly-amount-failed", bankHandler.FindYearAmountBankFailed)

	routerBank.GET("/monthly-method-success", bankHandler.FindMonthMethodBankSuccess)
	routerBank.GET("/yearly-method-success", bankHandler.FindYearMethodBankSuccess)
	routerBank.GET("/monthly-method-failed", bankHandler.FindMonthMethodBankSuccess)
	routerBank.GET("/yearly-method-failed", bankHandler.FindYearMethodBankSuccess)

	routerBank.GET("/mybank/monthly-amount-success/:id", bankHandler.FindMonthAmountBankSuccessById)
	routerBank.GET("/mybank/yearly-amount-success/:id", bankHandler.FindYearAmountBankSuccessById)
	routerBank.GET("/mybank/monthly-amount-failed/:id", bankHandler.FindMonthAmountBankFailedById)
	routerBank.GET("/mybank/yearly-amount-failed/:id", bankHandler.FindYearAmountBankFailedById)

	routerBank.GET("/mybank/monthly-method-success/:id", bankHandler.FindMonthMethodBankSuccessById)
	routerBank.GET("/mybank/yearly-method-success/:id", bankHandler.FindYearMethodBankSuccessById)
	routerBank.GET("/mybank/monthly-method-failed/:id", bankHandler.FindMonthMethodBankSuccessById)
	routerBank.GET("/mybank/yearly-method-failed/:id", bankHandler.FindYearMethodBankSuccessById)

	routerBank.GET("/merchant/monthly-amount-success/:merchant_id", bankHandler.FindMonthAmountBankSuccessByMerchant)
	routerBank.GET("/merchant/yearly-amount-success/:merchant_id", bankHandler.FindYearAmountBankSuccessByMerchant)
	routerBank.GET("/merchant/monthly-amount-failed/:merchant_id", bankHandler.FindMonthAmountBankFailedByMerchant)
	routerBank.GET("/merchant/yearly-amount-failed/:merchant_id", bankHandler.FindYearAmountBankFailedByMerchant)

	routerBank.GET("/merchant/monthly-method-success/:merchant_id", bankHandler.FindMonthMethodBankSuccessByMerchant)
	routerBank.GET("/merchant/yearly-method-success/:merchant_id", bankHandler.FindYearMethodBankSuccessByMerchant)
	routerBank.GET("/merchant/monthly-method-failed/:merchant_id", bankHandler.FindMonthMethodBankSuccessByMerchant)
	routerBank.GET("/merchant/yearly-method-failed/:merchant_id", bankHandler.FindYearMethodBankSuccessByMerchant)

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
		h.logger.Debug("Failed to fetch Bank records", zap.Error(err))
		return bank_errors.ErrApiBankNotFound(c)
	}

	so := h.mapping.ToApiResponsePaginationBank(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find monthly bank success amounts
// @Tags Bank
// @Description Retrieve monthly bank success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseBankMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/monthly-amount-success [get]
func (h *bankHandleApi) FindMonthAmountBankSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountBankRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.bank.FindMonthAmountBankSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank success amounts", zap.Error(err))
		return bank_errors.ErrApiFindMonthAmountBankSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank success amounts
// @Tags Bank
// @Description Retrieve yearly bank success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/yearly-amount-success [get]
func (h *bankHandleApi) FindYearAmountBankSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankRequest{
		Year: int32(year),
	}

	res, err := h.bank.FindYearAmountBankSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank success amounts", zap.Error(err))
		return bank_errors.ErrApiFindYearAmountBankSuccess(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank failed amounts
// @Tags Bank
// @Description Retrieve monthly bank failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseBankMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/monthly-amount-failed [get]
func (h *bankHandleApi) FindMonthAmountBankFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountBankRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.bank.FindMonthAmountBankFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank failed amounts", zap.Error(err))
		return bank_errors.ErrApiFindMonthAmountBankFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank failed amounts
// @Tags Bank
// @Description Retrieve yearly bank failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/yearly-amount-failed [get]
func (h *bankHandleApi) FindYearAmountBankFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankRequest{
		Year: int32(year),
	}

	res, err := h.bank.FindYearAmountBankFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank failed amounts", zap.Error(err))
		return bank_errors.ErrApiFindYearAmountBankFailed(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank success methods
// @Tags Bank
// @Description Retrieve monthly bank success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/monthly-method-success/{id} [get]
func (h *bankHandleApi) FindMonthMethodBankSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankRequest{
		Year: int32(year),
	}

	res, err := h.bank.FindMonthMethodBankSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank success methods", zap.Error(err))
		return bank_errors.ErrApiFindMonthMethodBankSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank success methods
// @Tags Bank
// @Description Retrieve yearly bank success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/yearly-method-success/{id} [get]
func (h *bankHandleApi) FindYearMethodBankSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankRequest{
		Year: int32(year),
	}

	res, err := h.bank.FindYearMethodBankSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank success methods", zap.Error(err))
		return bank_errors.ErrApiFindYearMethodBankSuccess(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank failed methods
// @Tags Bank
// @Description Retrieve monthly bank failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/monthly-method-failed/{id} [get]
func (h *bankHandleApi) FindMonthMethodBankFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankRequest{
		Year: int32(year),
	}

	res, err := h.bank.FindMonthMethodBankFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank failed methods", zap.Error(err))
		return bank_errors.ErrApiFindMonthMethodBankFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank failed methods
// @Tags Bank
// @Description Retrieve yearly bank failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/yearly-method-failed/{id} [get]
func (h *bankHandleApi) FindYearMethodBankFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankRequest{
		Year: int32(year),
	}

	res, err := h.bank.FindYearMethodBankFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank failed methods", zap.Error(err))
		return bank_errors.ErrApiFindYearMethodBankFailed(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank success amounts by ID
// @Tags Bank
// @Description Retrieve monthly bank success amounts by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseBankMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/monthly-amount-success/{id} [get]
func (h *bankHandleApi) FindMonthAmountBankSuccessById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
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
	req := &pb.MonthAmountBankByIdRequest{
		Id:    int32(bankID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.bank.FindMonthAmountBankSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank success amounts by ID", zap.Error(err))
		return bank_errors.ErrApiFindMonthAmountBankSuccessById(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank success amounts by ID
// @Tags Bank
// @Description Retrieve yearly bank success amounts by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/yearly-amount-success/{id} [get]
func (h *bankHandleApi) FindYearAmountBankSuccessById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankByIdRequest{
		Id:   int32(bankID),
		Year: int32(year),
	}

	res, err := h.bank.FindYearAmountBankSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank success amounts by ID", zap.Error(err))
		return bank_errors.ErrApiFindYearAmountBankSuccessById(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank failed amounts by ID
// @Tags Bank
// @Description Retrieve monthly bank failed amounts by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseBankMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/monthly-amount-failed/{id} [get]
func (h *bankHandleApi) FindMonthAmountBankFailedById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
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
	req := &pb.MonthAmountBankByIdRequest{
		Id:    int32(bankID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.bank.FindMonthAmountBankFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank failed amounts by ID", zap.Error(err))
		return bank_errors.ErrApiFindMonthAmountBankFailedById(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank failed amounts by ID
// @Tags Bank
// @Description Retrieve yearly bank failed amounts by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/yearly-amount-failed/{id} [get]
func (h *bankHandleApi) FindYearAmountBankFailedById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankByIdRequest{
		Id:   int32(bankID),
		Year: int32(year),
	}

	res, err := h.bank.FindYearAmountBankFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank failed amounts by ID", zap.Error(err))
		return bank_errors.ErrApiFindYearAmountBankFailedById(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank success methods by ID
// @Tags Bank
// @Description Retrieve monthly bank success methods by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/monthly-method-success/{id} [get]
func (h *bankHandleApi) FindMonthMethodBankSuccessById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodBankByIdRequest{
		Id:   int32(bankID),
		Year: int32(year),
	}

	res, err := h.bank.FindMonthMethodBankSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank success methods by ID", zap.Error(err))
		return bank_errors.ErrApiFindMonthMethodBankSuccessById(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank success methods by ID
// @Tags Bank
// @Description Retrieve yearly bank success methods by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/yearly-method-success/{id} [get]
func (h *bankHandleApi) FindYearMethodBankSuccessById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodBankByIdRequest{
		Id:   int32(bankID),
		Year: int32(year),
	}

	res, err := h.bank.FindYearMethodBankSuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank success methods by ID", zap.Error(err))
		return bank_errors.ErrApiFindYearMethodBankSuccessById(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank failed methods by ID
// @Tags Bank
// @Description Retrieve monthly bank failed methods by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/monthly-method-failed/{id} [get]
func (h *bankHandleApi) FindMonthMethodBankFailedById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodBankByIdRequest{
		Id:   int32(bankID),
		Year: int32(year),
	}

	res, err := h.bank.FindMonthMethodBankFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank failed methods by ID", zap.Error(err))
		return bank_errors.ErrApiFindMonthMethodBankFailedById(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank failed methods by ID
// @Tags Bank
// @Description Retrieve yearly bank failed methods by bank ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/mycategory/yearly-method-failed/{id} [get]
func (h *bankHandleApi) FindYearMethodBankFailedById(c echo.Context) error {
	bankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || bankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodBankByIdRequest{
		Id:   int32(bankID),
		Year: int32(year),
	}

	res, err := h.bank.FindYearMethodBankFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank failed methods by ID", zap.Error(err))
		return bank_errors.ErrApiFindYearMethodBankFailedById(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank success amounts by merchant
// @Tags Bank
// @Description Retrieve monthly bank success amounts by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseBankMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/monthly-amount-success/{merchant_id} [get]
func (h *bankHandleApi) FindMonthAmountBankSuccessByMerchant(c echo.Context) error {
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
	req := &pb.MonthAmountBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.bank.FindMonthAmountBankSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank success amounts by merchant", zap.Error(err))
		return bank_errors.ErrApiFindMonthAmountBankSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank success amounts by merchant
// @Tags Bank
// @Description Retrieve yearly bank success amounts by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/yearly-amount-success/{merchant_id} [get]
func (h *bankHandleApi) FindYearAmountBankSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.bank.FindYearAmountBankSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank success amounts by merchant", zap.Error(err))
		return bank_errors.ErrApiFindYearAmountBankSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank failed amounts by merchant
// @Tags Bank
// @Description Retrieve monthly bank failed amounts by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseBankMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/monthly-amount-failed/{merchant_id} [get]
func (h *bankHandleApi) FindMonthAmountBankFailedByMerchant(c echo.Context) error {
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
	req := &pb.MonthAmountBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.bank.FindMonthAmountBankFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank failed amounts by merchant", zap.Error(err))
		return bank_errors.ErrApiFindMonthAmountBankFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank failed amounts by merchant
// @Tags Bank
// @Description Retrieve yearly bank failed amounts by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/yearly-amount-failed/{merchant_id} [get]
func (h *bankHandleApi) FindYearAmountBankFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.bank.FindYearAmountBankFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank failed amounts by merchant", zap.Error(err))
		return bank_errors.ErrApiFindYearAmountBankFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank success methods by merchant
// @Tags Bank
// @Description Retrieve monthly bank success methods by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/monthly-method-success/{merchant_id} [get]
func (h *bankHandleApi) FindMonthMethodBankSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.bank.FindMonthMethodBankSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank success methods by merchant", zap.Error(err))
		return bank_errors.ErrApiFindMonthMethodBankSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank success methods by merchant
// @Tags Bank
// @Description Retrieve yearly bank success methods by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/yearly-method-success/{merchant_id} [get]
func (h *bankHandleApi) FindYearMethodBankSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.bank.FindYearMethodBankSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank success methods by merchant", zap.Error(err))
		return bank_errors.ErrApiFindYearMethodBankSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly bank failed methods by merchant
// @Tags Bank
// @Description Retrieve monthly bank failed methods by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/monthly-method-failed/{merchant_id} [get]
func (h *bankHandleApi) FindMonthMethodBankFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.bank.FindMonthMethodBankFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly bank failed methods by merchant", zap.Error(err))
		return bank_errors.ErrApiFindMonthMethodBankFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly bank failed methods by merchant
// @Tags Bank
// @Description Retrieve yearly bank failed methods by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseBankYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/bank/merchant/yearly-method-failed/{merchant_id} [get]
func (h *bankHandleApi) FindYearMethodBankFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodBankByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.bank.FindYearMethodBankFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly bank failed methods by merchant", zap.Error(err))
		return bank_errors.ErrApiFindYearMethodBankFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find bank by ID
// @Tags Bank
// @Description Retrieve a bank by ID
// @Accept json
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} response.ApiResponseBank "Bank data"
// @Failure 400 {object} response.ErrorResponse "Invalid bank ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve bank data"
// @Router /api/bank/{id} [get]
func (h *bankHandleApi) FindById(c echo.Context) error {
	BankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || BankID <= 0 {
		return bank_errors.ErrApiBankInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Bank", zap.Error(err))
		return bank_errors.ErrApiBankNotFound(c)
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
		return bank_errors.ErrApiFailedFindActive(c)
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
		return bank_errors.ErrApiFailedFindTrashed(c)
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
	var req requests.CreateBankRequest
	if err := c.Bind(&req); err != nil {
		return bank_errors.ErrApiBindCreateBank(c)
	}

	if err := req.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return bank_errors.ErrApiValidateCreateBank(c)
	}

	ctx := c.Request().Context()

	reqPb := &pb.CreateBankRequest{
		Name: req.Name,
	}

	res, err := h.bank.Create(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to create Bank", zap.Error(err))
		return bank_errors.ErrApiFailedCreateBank(c)
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
		return bank_errors.ErrInvalidBankId(c)
	}

	var req requests.UpdateBankRequest

	if err := c.Bind(&req); err != nil {
		return bank_errors.ErrApiBindUpdateBank(c)
	}

	if err := req.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return bank_errors.ErrApiValidateCreateBank(c)
	}

	reqPb := &pb.UpdateBankRequest{
		Id:   int32(BankID),
		Name: req.Name,
	}

	ctx := c.Request().Context()

	res, err := h.bank.Update(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to update Bank", zap.Error(err))
		return bank_errors.ErrApiFailedUpdateBank(c)
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
		return bank_errors.ErrInvalidBankId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Bank", zap.Error(err))
		return bank_errors.ErrApiFailedTrashedBank(c)
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
		return bank_errors.ErrInvalidBankId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Bank", zap.Error(err))
		return bank_errors.ErrApiFailedRestoreBank(c)
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
// @Success 200 {object} response.ApiResponseBankDelete "Successfully deleted bank record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete bank"
// @Router /api/bank/delete/{id} [post]
func (h *bankHandleApi) DeletePermanent(c echo.Context) error {
	BankID, err := strconv.Atoi(c.Param("id"))
	if err != nil || BankID <= 0 {
		return bank_errors.ErrInvalidBankId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdBankRequest{
		BankId: int32(BankID),
	}

	res, err := h.bank.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Bank permanently", zap.Error(err))
		return bank_errors.ErrApiFailedDeletePermanent(c)
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
// @Success 200 {object} response.ApiResponseBankAll "Successfully restored all banks"
// @Failure 500 {object} response.ErrorResponse "Failed to restore banks"
// @Router /api/bank/restore/all [post]
func (h *bankHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.bank.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Banks", zap.Error(err))
		return bank_errors.ErrApiFailedRestoreAll(c)
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
// @Success 200 {object} response.ApiResponseBankDelete "Successfully deleted all bank records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete banks"
// @Router /api/bank/delete/all [post]
func (h *bankHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.bank.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Banks permanently", zap.Error(err))
		return bank_errors.ErrApiFailedDeleteAll(c)
	}

	so := h.mapping.ToApiResponseBankAll(res)

	return c.JSON(http.StatusOK, so)
}
