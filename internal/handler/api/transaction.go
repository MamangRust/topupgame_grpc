package api

import (
	"net/http"
	"strconv"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_api "topup_game/internal/mapper/response/api"
	"topup_game/internal/pb"
	"topup_game/pkg/errors/nominal_errors"
	"topup_game/pkg/errors/transaction_errors"
	"topup_game/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transactionHandleApi struct {
	transaction pb.TransactionServiceClient
	logger      logger.LoggerInterface
	mapping     response_api.TransactionResponseMapper
}

func NewHandlerTransaction(router *echo.Echo, transaction pb.TransactionServiceClient, logger logger.LoggerInterface, mapping response_api.TransactionResponseMapper) *transactionHandleApi {
	transactionHandler := &transactionHandleApi{
		transaction: transaction,
		logger:      logger,
		mapping:     mapping,
	}

	routerTransaction := router.Group("/api/transaction")

	routerTransaction.GET("", transactionHandler.FindAll)
	routerTransaction.GET("/:id", transactionHandler.FindById)
	routerTransaction.GET("/active", transactionHandler.FindByActive)
	routerTransaction.GET("/trashed", transactionHandler.FindByTrashed)
	routerTransaction.POST("/create", transactionHandler.Create)

	routerTransaction.GET("/monthly-amount-success", transactionHandler.FindMonthAmountTransactionSuccess)
	routerTransaction.GET("/yearly-amount-success", transactionHandler.FindYearAmountTransactionSuccess)
	routerTransaction.GET("/monthly-amount-failed", transactionHandler.FindMonthAmountTransactionFailed)
	routerTransaction.GET("/yearly-amount-failed", transactionHandler.FindYearAmountTransactionFailed)

	routerTransaction.GET("/monthly-method-success", transactionHandler.FindMonthMethodTransactionSuccess)
	routerTransaction.GET("/yearly-method-success", transactionHandler.FindYearMethodTransactionSuccess)
	routerTransaction.GET("/monthly-method-failed", transactionHandler.FindMonthMethodTransactionSuccess)
	routerTransaction.GET("/yearly-method-failed", transactionHandler.FindYearMethodTransactionSuccess)

	routerTransaction.GET("/merchant/monthly-amount-success/:id", transactionHandler.FindMonthAmountTransactionSuccessByMerchant)
	routerTransaction.GET("/merchant/yearly-amount-success/:id", transactionHandler.FindYearAmountTransactionSuccessByMerchant)
	routerTransaction.GET("/merchant/monthly-amount-failed/:id", transactionHandler.FindMonthAmountTransactionFailedByMerchant)
	routerTransaction.GET("/merchant/yearly-amount-failed/:id", transactionHandler.FindYearAmountTransactionFailedByMerchant)

	routerTransaction.GET("/merchant/monthly-method-success/:id", transactionHandler.FindMonthMethodTransactionSuccessByMerchant)
	routerTransaction.GET("/merchant/yearly-method-success/:id", transactionHandler.FindYearMethodTransactionSuccessByMerchant)
	routerTransaction.GET("/merchant/monthly-method-failed/:id", transactionHandler.FindMonthMethodTransactionSuccessByMerchant)
	routerTransaction.GET("/merchant/yearly-method-failed/:id", transactionHandler.FindYearMethodTransactionSuccessByMerchant)

	routerTransaction.POST("/:id", transactionHandler.Update)
	routerTransaction.DELETE("/:id", transactionHandler.Trashed)
	routerTransaction.PUT("/restore/:id", transactionHandler.Restore)
	routerTransaction.DELETE("/permanent/:id", transactionHandler.DeletePermanent)
	routerTransaction.PUT("/restore-all", transactionHandler.RestoreAll)
	routerTransaction.DELETE("/permanent-all", transactionHandler.DeleteAllPermanent)

	return transactionHandler
}

// @Security Bearer
// @Summary Find all Transactions
// @Tags Transaction
// @Description Retrieve a list of all Transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransaction "List of Transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Transaction data"
// @Router /api/transaction [get]
func (h *transactionHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.transaction.FindAll(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Transaction records", zap.Error(err))
		return transaction_errors.ErrApiFailedFindAll(c)
	}

	so := h.mapping.ToApiResponsePaginationTransaction(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find monthly transaction success amounts
// @Tags Transaction
// @Description Retrieve monthly transaction success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/monthly-amount-success [get]
func (h *transactionHandleApi) FindMonthAmountTransactionSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountTransactionRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.transaction.FindMonthAmountTransactionSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction success amounts", zap.Error(err))
		return transaction_errors.ErrApiFindMonthAmountTransactionSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction success amounts
// @Tags Transaction
// @Description Retrieve yearly transaction success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/yearly-amount-success [get]
func (h *transactionHandleApi) FindYearAmountTransactionSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionRequest{
		Year: int32(year),
	}

	res, err := h.transaction.FindYearAmountTransactionSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction success amounts", zap.Error(err))
		return transaction_errors.ErrApiFindYearAmountTransactionSuccess(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly transaction failed amounts
// @Tags Transaction
// @Description Retrieve monthly transaction failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/monthly-amount-failed [get]
func (h *transactionHandleApi) FindMonthAmountTransactionFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountTransactionRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.transaction.FindMonthAmountTransactionFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction failed amounts", zap.Error(err))
		return transaction_errors.ErrApiFindMonthAmountTransactionFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction failed amounts
// @Tags Transaction
// @Description Retrieve yearly transaction failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/yearly-amount-failed [get]
func (h *transactionHandleApi) FindYearAmountTransactionFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionRequest{
		Year: int32(year),
	}

	res, err := h.transaction.FindYearAmountTransactionFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction failed amounts", zap.Error(err))
		return transaction_errors.ErrApiFindYearAmountTransactionFailed(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly transaction success methods
// @Tags Transaction
// @Description Retrieve monthly transaction success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/monthly-method-successs [get]
func (h *transactionHandleApi) FindMonthMethodTransactionSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionRequest{
		Year: int32(year),
	}

	res, err := h.transaction.FindMonthMethodTransactionSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction success methods", zap.Error(err))
		return transaction_errors.ErrApiFindMonthMethodTransactionSuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction success methods
// @Tags Transaction
// @Description Retrieve yearly transaction success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/yearly-method-success [get]
func (h *transactionHandleApi) FindYearMethodTransactionSuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionRequest{
		Year: int32(year),
	}

	res, err := h.transaction.FindYearMethodTransactionSuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction success methods", zap.Error(err))
		return transaction_errors.ErrApiFindYearMethodTransactionSuccess(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly transaction failed methods
// @Tags Transaction
// @Description Retrieve monthly transaction failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/monthly-method-failed [get]
func (h *transactionHandleApi) FindMonthMethodTransactionFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionRequest{
		Year: int32(year),
	}

	res, err := h.transaction.FindMonthMethodTransactionFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction failed methods", zap.Error(err))
		return transaction_errors.ErrApiFindMonthMethodTransactionFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction failed methods
// @Tags Transaction
// @Description Retrieve yearly transaction failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/yearly-method-failed [get]
func (h *transactionHandleApi) FindYearMethodTransactionFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionRequest{
		Year: int32(year),
	}

	res, err := h.transaction.FindYearMethodTransactionFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction failed methods", zap.Error(err))
		return transaction_errors.ErrApiFindYearMethodTransactionFailed(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly transaction success amounts by merchant
// @Tags Transaction
// @Description Retrieve monthly transaction success amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/monthly-amount-success/{merchant_id} [get]
func (h *transactionHandleApi) FindMonthAmountTransactionSuccessByMerchant(c echo.Context) error {
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
	req := &pb.MonthAmountTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.transaction.FindMonthAmountTransactionSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction success amounts by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindMonthAmountTransactionSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction success amounts by merchant
// @Tags Transaction
// @Description Retrieve yearly transaction success amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/yearly-success-amount/{merchant_id} [get]
func (h *transactionHandleApi) FindYearAmountTransactionSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.transaction.FindYearAmountTransactionSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction success amounts by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindYearAmountTransactionSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly transaction failed amounts by merchant
// @Tags Transaction
// @Description Retrieve monthly transaction failed amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/monthly-amount-failed/{merchant_id} [get]
func (h *transactionHandleApi) FindMonthAmountTransactionFailedByMerchant(c echo.Context) error {
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
	req := &pb.MonthAmountTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.transaction.FindMonthAmountTransactionFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction failed amounts by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindMonthAmountTransactionFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction failed amounts by merchant
// @Tags Transaction
// @Description Retrieve yearly transaction failed amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/yearly-amount-failed/{merchant_id} [get]
func (h *transactionHandleApi) FindYearAmountTransactionFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.transaction.FindYearAmountTransactionFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction failed amounts by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindYearAmountTransactionFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly transaction success methods by merchant
// @Tags Transaction
// @Description Retrieve monthly transaction success methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/monthly-method-success/{merchant_id} [get]
func (h *transactionHandleApi) FindMonthMethodTransactionSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.transaction.FindMonthMethodTransactionSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction success methods by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindMonthMethodTransactionSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction success methods by merchant
// @Tags Transaction
// @Description Retrieve yearly transaction success methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/yearly-method-success/{merchant_id} [get]
func (h *transactionHandleApi) FindYearMethodTransactionSuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.transaction.FindYearMethodTransactionSuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction success methods by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindYearMethodTransactionSuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly transaction failed methods by merchant
// @Tags Transaction
// @Description Retrieve monthly transaction failed methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/monthly-method-failed/{merchant_id} [get]
func (h *transactionHandleApi) FindMonthMethodTransactionFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.transaction.FindMonthMethodTransactionFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly transaction failed methods by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindMonthMethodTransactionFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly transaction failed methods by merchant
// @Tags Transaction
// @Description Retrieve yearly transaction failed methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/transaction/merchant/yearly-method-failed/{merchant_id} [get]
func (h *transactionHandleApi) FindYearMethodTransactionFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodTransactionByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.transaction.FindYearMethodTransactionFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly transaction failed methods by merchant", zap.Error(err))
		return transaction_errors.ErrApiFindYearMethodTransactionFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find Transaction by ID
// @Tags Transaction
// @Description Retrieve a Transaction by ID
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransaction "Transaction data"
// @Failure 400 {object} response.ErrorResponse "Invalid Transaction ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Transaction data"
// @Router /api/transaction/{id} [get]
func (h *transactionHandleApi) FindById(c echo.Context) error {
	TransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || TransactionID <= 0 {
		return transaction_errors.ErrApiTransactionInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Transaction", zap.Error(err))
		return transaction_errors.ErrApiTransactionNotFound(c)
	}

	so := h.mapping.ToApiResponseTransaction(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active Transactions
// @Tags Transaction
// @Description Retrieve a list of active Transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransactionDeleteAt "List of active Transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Transaction data"
// @Router /api/transaction/active [get]
func (h *transactionHandleApi) FindByActive(c echo.Context) error {
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

	req := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.transaction.FindByActive(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch active Transactions", zap.Error(err))
		return nominal_errors.ErrApiFailedFindActive(c)
	}

	so := h.mapping.ToApiResponsePaginationTransactionDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve trashed Transactions
// @Tags Transaction
// @Description Retrieve a list of trashed Transaction records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransactionDeleteAt "List of trashed Transaction data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Transaction data"
// @Router /api/transaction/trashed [get]
func (h *transactionHandleApi) FindByTrashed(c echo.Context) error {
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

	req := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.transaction.FindByTrashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch trashed Transactions", zap.Error(err))
		return nominal_errors.ErrApiFailedFindTrashed(c)
	}

	so := h.mapping.ToApiResponsePaginationTransactionDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create an existing Transaction
// @Tags Transaction
// @Description Create an existing Transaction record with the provided details
// @Accept json
// @Produce json
// @Param CreateTransactionRequest body requests.CreateTransactionRequest true "Create Transaction request"
// @Success 200 {object} response.ApiResponseTransaction "Successfully created Transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create Transaction"
// @Router /api/transaction/create [post]
func (h *transactionHandleApi) Create(c echo.Context) error {
	var req requests.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return nominal_errors.ErrApiBindCreateNominal(c)
	}

	if err := req.Validate(); err != nil {
		return nominal_errors.ErrApiValidateCreateNominal(c)
	}

	reqPb := &pb.CreateTransactionRequest{
		UserId:        int32(req.UserID),
		MerchantId:    int32(req.MerchantID),
		NominalId:     int32(req.NominalID),
		BankId:        int32(req.BankID),
		PaymentMethod: req.PaymentMethod,
	}

	ctx := c.Request().Context()

	res, err := h.transaction.Create(ctx, reqPb)

	if err != nil {
		h.logger.Debug("Failed to create Transaction", zap.Error(err))
		return nominal_errors.ErrApiFailedCreateNominal(c)
	}

	so := h.mapping.ToApiResponseTransaction(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update an existing Transaction
// @Tags Transaction
// @Description Update an existing Transaction record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param UpdateTransactionRequest body requests.UpdateTransactionRequest true "Update Transaction request"
// @Success 200 {object} response.ApiResponseTransaction "Successfully updated Transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update Transaction"
// @Router /api/transaction/update/{id} [post]
func (h *transactionHandleApi) Update(c echo.Context) error {
	TransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || TransactionID <= 0 {
		return transaction_errors.ErrApiTransactionInvalidId(c)
	}

	var req requests.UpdateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return transaction_errors.ErrApiBindUpdateTransaction(c)
	}

	if err := req.Validate(); err != nil {
		return transaction_errors.ErrApiValidateUpdateTransaction(c)
	}

	reqPb := &pb.UpdateTransactionRequest{
		Id:            int32(TransactionID),
		UserId:        int32(req.UserID),
		MerchantId:    int32(req.MerchantID),
		NominalId:     int32(req.NominalID),
		BankId:        int32(req.BankID),
		PaymentMethod: req.PaymentMethod,
	}

	ctx := c.Request().Context()

	res, err := h.transaction.Update(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to update Transaction", zap.Error(err))
		return transaction_errors.ErrApiFailedUpdateTransaction(c)
	}

	so := h.mapping.ToApiResponseTransaction(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve a trashed Transaction
// @Tags Transaction
// @Description Retrieve a trashed Transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransactionDeleteAt "Successfully retrieved trashed Transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed Transaction"
// @Router /api/transaction/trashed/{id} [get]
func (h *transactionHandleApi) Trashed(c echo.Context) error {
	TransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || TransactionID <= 0 {
		return transaction_errors.ErrApiTransactionInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Transaction", zap.Error(err))
		return transaction_errors.ErrApiFailedTrashedTransaction(c)
	}

	so := h.mapping.ToApiResponseTransactionDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore a trashed Transaction
// @Tags Transaction
// @Description Restore a trashed Transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransactionDeleteAt "Successfully restored Transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid Transaction ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Transaction"
// @Router /api/transaction/restore/{id} [post]
func (h *transactionHandleApi) Restore(c echo.Context) error {
	TransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || TransactionID <= 0 {
		return nominal_errors.ErrApiNominalInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Transaction", zap.Error(err))
		return nominal_errors.ErrApiFailedRestoreNominal(c)
	}

	so := h.mapping.ToApiResponseTransactionDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete a Transaction
// @Tags Transaction
// @Description Permanently delete a Transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransactionDelete "Successfully deleted Transaction record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Transaction"
// @Router /api/transaction/delete/{id} [post]
func (h *transactionHandleApi) DeletePermanent(c echo.Context) error {
	TransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || TransactionID <= 0 {
		return nominal_errors.ErrApiNominalInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Transaction permanently", zap.Error(err))
		return nominal_errors.ErrApiFailedDeletePermanent(c)
	}

	so := h.mapping.ToApiResponseTransactionDelete(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore all trashed Transactions
// @Tags Transaction
// @Description Restore all trashed Transaction records.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseTransactionAll "Successfully restored all Transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Transactions"
// @Router /api/transaction/restore/all [post]
func (h *transactionHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Transactions", zap.Error(err))
		return nominal_errors.ErrApiFailedRestoreAll(c)
	}

	so := h.mapping.ToApiResponseTransactionAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete all trashed Transactions
// @Tags Transaction
// @Description Permanently delete all trashed Transaction records.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseTransactionDelete "Successfully deleted all Transaction records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Transactions"
// @Router /api/transaction/delete/all [post]
func (h *transactionHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Transactions permanently", zap.Error(err))
		return nominal_errors.ErrApiFailedDeleteAll(c)
	}

	so := h.mapping.ToApiResponseTransactionAll(res)

	return c.JSON(http.StatusOK, so)
}
