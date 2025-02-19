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
	routerTransaction.POST("", transactionHandler.Create)
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
		if errors.Is(err, echo.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status:  "error",
				Message: "Unauthorized",
			})
		}

		h.logger.Debug("Failed to fetch Transaction records", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Transaction records",
		})
	}

	so := h.mapping.ToApiResponsePaginationTransaction(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find Transaction by ID
// @Tags Transaction
// @Description Retrieve a Transaction by ID
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} pb.ApiResponseTransaction "Transaction data"
// @Failure 400 {object} response.ErrorResponse "Invalid Transaction ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Transaction data"
// @Router /api/transaction/{id} [get]
func (h *transactionHandleApi) FindById(c echo.Context) error {
	TransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || TransactionID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Transaction ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Transaction", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Transaction",
		})
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Transactions",
		})
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Transactions",
		})
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
	var req pb.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	ctx := c.Request().Context()

	res, err := h.transaction.Create(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to create Transaction", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create Transaction",
		})
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Transaction ID",
		})
	}

	var req pb.UpdateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	req.Id = int32(TransactionID)

	ctx := c.Request().Context()

	res, err := h.transaction.Update(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to update Transaction", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update Transaction",
		})
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Transaction ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Transaction", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Transaction",
		})
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Transaction ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Transaction", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Transaction",
		})
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
// @Success 200 {object} pb.ApiResponseTransactionDelete "Successfully deleted Transaction record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Transaction"
// @Router /api/transaction/delete/{id} [post]
func (h *transactionHandleApi) DeletePermanent(c echo.Context) error {
	TransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil || TransactionID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Transaction ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdTransactionRequest{
		Id: int32(TransactionID),
	}

	res, err := h.transaction.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Transaction permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Transaction permanently",
		})
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
// @Success 200 {object} pb.ApiResponseTransactionAll "Successfully restored all Transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Transactions"
// @Router /api/transaction/restore/all [post]
func (h *transactionHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Transactions", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Transactions",
		})
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
// @Success 200 {object} pb.ApiResponseTransactionDelete "Successfully deleted all Transaction records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Transactions"
// @Router /api/transaction/delete/all [post]
func (h *transactionHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.transaction.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Transactions permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Transactions permanently",
		})
	}

	so := h.mapping.ToApiResponseTransactionAll(res)

	return c.JSON(http.StatusOK, so)
}
