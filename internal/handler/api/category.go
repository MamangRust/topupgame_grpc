package api

import (
	"net/http"
	"strconv"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_api "topup_game/internal/mapper/response/api"
	"topup_game/internal/pb"
	"topup_game/pkg/errors/category_errors"
	"topup_game/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryHandleApi struct {
	category pb.CategoryServiceClient
	logger   logger.LoggerInterface
	mapping  response_api.CategoryResponseMapper
}

func NewHandlerCategory(router *echo.Echo, category pb.CategoryServiceClient, logger logger.LoggerInterface, mapping response_api.CategoryResponseMapper) *categoryHandleApi {
	categoryHandler := &categoryHandleApi{
		category: category,
		logger:   logger,
		mapping:  mapping,
	}

	routerCategory := router.Group("/api/category")

	routerCategory.GET("", categoryHandler.FindAll)
	routerCategory.GET("/:id", categoryHandler.FindById)
	routerCategory.GET("/active", categoryHandler.FindByActive)
	routerCategory.GET("/trashed", categoryHandler.FindByTrashed)
	routerCategory.POST("/create", categoryHandler.Create)

	routerCategory.GET("/monthly-amount-success", categoryHandler.FindMonthAmountCategorySuccess)
	routerCategory.GET("/yearly-amount-success", categoryHandler.FindYearAmountCategorySuccess)
	routerCategory.GET("/monthly-amount-failed", categoryHandler.FindMonthAmountCategoryFailed)
	routerCategory.GET("/yearly-amount-failed", categoryHandler.FindYearAmountCategoryFailed)

	routerCategory.GET("/monthly-method-success", categoryHandler.FindMonthMethodCategorySuccess)
	routerCategory.GET("/yearly-method-success", categoryHandler.FindYearMethodCategorySuccess)
	routerCategory.GET("/monthly-method-failed", categoryHandler.FindMonthMethodCategorySuccess)
	routerCategory.GET("/yearly-method-failed", categoryHandler.FindYearMethodCategorySuccess)

	routerCategory.GET("/mycategory/monthly-amount-success/:id", categoryHandler.FindMonthAmountCategorySuccessById)
	routerCategory.GET("/mycategory/yearly-amount-success/:id", categoryHandler.FindYearAmountCategorySuccessById)
	routerCategory.GET("/mycategory/monthly-amount-failed/:id", categoryHandler.FindMonthAmountCategoryFailedById)
	routerCategory.GET("/mycategory/yearly-amount-failed/:id", categoryHandler.FindYearAmountCategoryFailedById)

	routerCategory.GET("/mycategory/monthly-method-success/:id", categoryHandler.FindMonthMethodCategorySuccessById)
	routerCategory.GET("/mycategory/yearly-method-success/:id", categoryHandler.FindYearMethodCategorySuccessById)
	routerCategory.GET("/mycategory/monthly-method-failed/:id", categoryHandler.FindMonthMethodCategorySuccessById)
	routerCategory.GET("/mycategory/yearly-method-failed/:id", categoryHandler.FindYearMethodCategorySuccessById)

	routerCategory.GET("/merchant/monthly-amount-success/:merchant_merchant_id", categoryHandler.FindMonthAmountCategorySuccessByMerchant)
	routerCategory.GET("/merchant/yearly-amount-success/:merchant_id", categoryHandler.FindYearAmountCategorySuccessByMerchant)
	routerCategory.GET("/merchant/monthly-amount-failed/:merchant_id", categoryHandler.FindMonthAmountCategoryFailedByMerchant)
	routerCategory.GET("/merchant/yearly-amount-failed/:merchant_id", categoryHandler.FindYearAmountCategoryFailedByMerchant)

	routerCategory.GET("/merchant/monthly-method-success/:merchant_id", categoryHandler.FindMonthMethodCategorySuccessByMerchant)
	routerCategory.GET("/merchant/yearly-method-success/:merchant_id", categoryHandler.FindYearMethodCategorySuccessByMerchant)
	routerCategory.GET("/merchant/monthly-method-failed/:merchant_id", categoryHandler.FindMonthMethodCategorySuccessByMerchant)
	routerCategory.GET("/merchant/yearly-method-failed/:merchant_id", categoryHandler.FindYearMethodCategorySuccessByMerchant)

	routerCategory.POST("/:id", categoryHandler.Update)
	routerCategory.DELETE("/:id", categoryHandler.Trashed)
	routerCategory.PUT("/restore/:id", categoryHandler.Restore)
	routerCategory.DELETE("/permanent/:id", categoryHandler.DeletePermanent)
	routerCategory.PUT("/restore-all", categoryHandler.RestoreAll)
	routerCategory.DELETE("/permanent-all", categoryHandler.DeleteAllPermanent)

	return categoryHandler
}

// @Security Bearer
// @Summary Find all Categorys
// @Tags Category
// @Description Retrieve a list of all Categorys
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationCategory "List of Categorys"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Category data"
// @Router /api/category [get]
func (h *categoryHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllCategoryRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.category.FindAll(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Category records", zap.Error(err))
		return category_errors.ErrApiFailedFindAll(c)
	}

	so := h.mapping.ToApiResponsePaginationCategory(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find monthly category success amounts
// @Tags Category
// @Description Retrieve monthly category success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseCategoryMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/monthly-amount-success [get]
func (h *categoryHandleApi) FindMonthAmountCategorySuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountCategoryRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.category.FindMonthAmountCategorySuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category success amounts", zap.Error(err))
		return category_errors.ErrApiFindMonthAmountCategorySuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category success amounts
// @Tags Category
// @Description Retrieve yearly category success amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/yearly-amount-success [get]
func (h *categoryHandleApi) FindYearAmountCategorySuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryRequest{
		Year: int32(year),
	}

	res, err := h.category.FindYearAmountCategorySuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category success amounts", zap.Error(err))
		return category_errors.ErrApiFindYearAmountCategorySuccess(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category failed amounts
// @Tags Category
// @Description Retrieve monthly category failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseCategoryMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/monthly-amount-failed [get]
func (h *categoryHandleApi) FindMonthAmountCategoryFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil || month < 1 || month > 12 {
		return response.NewApiErrorResponse(c, "error", "invalid month", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthAmountCategoryRequest{
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.category.FindMonthAmountCategoryFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category failed amounts", zap.Error(err))
		return category_errors.ErrApiFindMonthAmountCategoryFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category failed amounts
// @Tags Category
// @Description Retrieve yearly category failed amounts
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/yearly-amount-failed [get]
func (h *categoryHandleApi) FindYearAmountCategoryFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryRequest{
		Year: int32(year),
	}

	res, err := h.category.FindYearAmountCategoryFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category failed amounts", zap.Error(err))
		return category_errors.ErrApiFindYearAmountCategoryFailed(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category success methods
// @Tags Category
// @Description Retrieve monthly category success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/monthly-method-success/{id} [get]
func (h *categoryHandleApi) FindMonthMethodCategorySuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryRequest{
		Year: int32(year),
	}

	res, err := h.category.FindMonthMethodCategorySuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category success methods", zap.Error(err))
		return category_errors.ErrApiFindMonthMethodCategorySuccess(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category success methods
// @Tags Category
// @Description Retrieve yearly category success methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/yearly-method-success [get]
func (h *categoryHandleApi) FindYearMethodCategorySuccess(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryRequest{
		Year: int32(year),
	}

	res, err := h.category.FindYearMethodCategorySuccess(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category success methods", zap.Error(err))
		return category_errors.ErrApiFindYearMethodCategorySuccess(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category failed methods
// @Tags Category
// @Description Retrieve monthly category failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/monthly-method-failed [get]
func (h *categoryHandleApi) FindMonthMethodCategoryFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryRequest{
		Year: int32(year),
	}

	res, err := h.category.FindMonthMethodCategoryFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category failed methods", zap.Error(err))
		return category_errors.ErrApiFindMonthMethodCategoryFailed(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category failed methods
// @Tags Category
// @Description Retrieve yearly category failed methods
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/yearly-method-failed [get]
func (h *categoryHandleApi) FindYearMethodCategoryFailed(c echo.Context) error {
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryRequest{
		Year: int32(year),
	}

	res, err := h.category.FindYearMethodCategoryFailed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category failed methods", zap.Error(err))
		return category_errors.ErrApiFindYearMethodCategoryFailed(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category success amounts by ID
// @Tags Category
// @Description Retrieve monthly category success amounts by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseCategoryMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid category ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/monthly-amount-success/{id} [get]
func (h *categoryHandleApi) FindMonthAmountCategorySuccessById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
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
	req := &pb.MonthAmountCategoryByIdRequest{
		Id:    int32(categoryID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.category.FindMonthAmountCategorySuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category success amounts by ID", zap.Error(err))
		return category_errors.ErrApiFindMonthAmountCategorySuccessById(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category success amounts by ID
// @Tags Category
// @Description Retrieve yearly category success amounts by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid category ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/yearly-amount-success/{id} [get]
func (h *categoryHandleApi) FindYearAmountCategorySuccessById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryByIdRequest{
		Id:   int32(categoryID),
		Year: int32(year),
	}

	res, err := h.category.FindYearAmountCategorySuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category success amounts by ID", zap.Error(err))
		return category_errors.ErrApiFindYearAmountCategorySuccessById(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category failed amounts by ID
// @Tags Category
// @Description Retrieve monthly category failed amounts by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseCategoryMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid category ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/monthly-amount-failed/{id} [get]
func (h *categoryHandleApi) FindMonthAmountCategoryFailedById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
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
	req := &pb.MonthAmountCategoryByIdRequest{
		Id:    int32(categoryID),
		Year:  int32(year),
		Month: int32(month),
	}

	res, err := h.category.FindMonthAmountCategoryFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category failed amounts by ID", zap.Error(err))
		return category_errors.ErrApiFindMonthAmountCategoryFailedById(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category failed amounts by ID
// @Tags Category
// @Description Retrieve yearly category failed amounts by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid category ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/yearly-amount-failed/{id} [get]
func (h *categoryHandleApi) FindYearAmountCategoryFailedById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryByIdRequest{
		Id:   int32(categoryID),
		Year: int32(year),
	}

	res, err := h.category.FindYearAmountCategoryFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category failed amounts by ID", zap.Error(err))
		return category_errors.ErrApiFindYearAmountCategoryFailedById(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category success methods by ID
// @Tags Category
// @Description Retrieve monthly category success methods by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid category ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/monthly-method-success/{id} [get]
func (h *categoryHandleApi) FindMonthMethodCategorySuccessById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodCategoryByIdRequest{
		Id:   int32(categoryID),
		Year: int32(year),
	}

	res, err := h.category.FindMonthMethodCategorySuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category success methods by ID", zap.Error(err))
		return category_errors.ErrApiFindMonthMethodCategorySuccessById(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category success methods by ID
// @Tags Category
// @Description Retrieve yearly category success methods by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid category ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/yearly-method-success/{id} [get]
func (h *categoryHandleApi) FindYearMethodCategorySuccessById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodCategoryByIdRequest{
		Id:   int32(categoryID),
		Year: int32(year),
	}

	res, err := h.category.FindYearMethodCategorySuccessById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category success methods by ID", zap.Error(err))
		return category_errors.ErrApiFindYearMethodCategorySuccessById(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category failed methods by ID
// @Tags Category
// @Description Retrieve monthly category failed methods by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid category ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/monthly-method-category/{id} [get]
func (h *categoryHandleApi) FindMonthMethodCategoryFailedById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodCategoryByIdRequest{
		Id:   int32(categoryID),
		Year: int32(year),
	}

	res, err := h.category.FindMonthMethodCategoryFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category failed methods by ID", zap.Error(err))
		return category_errors.ErrApiFindMonthMethodCategoryFailedById(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category failed methods by ID
// @Tags Category
// @Description Retrieve yearly category failed methods by category ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid category ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/mycategory/yearly-method-failed/{id} [get]
func (h *categoryHandleApi) FindYearMethodCategoryFailedById(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || categoryID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid category ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodCategoryByIdRequest{
		Id:   int32(categoryID),
		Year: int32(year),
	}

	res, err := h.category.FindYearMethodCategoryFailedById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category failed methods by ID", zap.Error(err))
		return category_errors.ErrApiFindYearMethodCategoryFailedById(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category success amounts by merchant
// @Tags Category
// @Description Retrieve monthly category success amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseCategoryMonthAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/monthly-amount-success/{merchant_id} [get]
func (h *categoryHandleApi) FindMonthAmountCategorySuccessByMerchant(c echo.Context) error {
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
	req := &pb.MonthAmountCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.category.FindMonthAmountCategorySuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category success amounts by merchant", zap.Error(err))
		return category_errors.ErrApiFindMonthAmountCategorySuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category success amounts by merchant
// @Tags Category
// @Description Retrieve yearly category success amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearAmountSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/yearly-amount-success/{merchant_id} [get]
func (h *categoryHandleApi) FindYearAmountCategorySuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.category.FindYearAmountCategorySuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category success amounts by merchant", zap.Error(err))
		return category_errors.ErrApiFindYearAmountCategorySuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountSuccess(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category failed amounts by merchant
// @Tags Category
// @Description Retrieve monthly category failed amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Param month query int true "Month (1-12)"
// @Success 200 {object} response.ApiResponseCategoryMonthAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/monthly-amount-failed/{merchant_id} [get]
func (h *categoryHandleApi) FindMonthAmountCategoryFailedByMerchant(c echo.Context) error {
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
	req := &pb.MonthAmountCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
		Month:      int32(month),
	}

	res, err := h.category.FindMonthAmountCategoryFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category failed amounts by merchant", zap.Error(err))
		return category_errors.ErrApiFindMonthAmountCategoryFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category failed amounts by merchant
// @Tags Category
// @Description Retrieve yearly category failed amounts by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearAmountFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/yearly-amount-failed/{merchant_id} [get]
func (h *categoryHandleApi) FindYearAmountCategoryFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearAmountCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.category.FindYearAmountCategoryFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category failed amounts by merchant", zap.Error(err))
		return category_errors.ErrApiFindYearAmountCategoryFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearAmountFailed(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category success methods by merchant
// @Tags Category
// @Description Retrieve monthly category success methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/monthly-method-success/{merchant_id} [get]
func (h *categoryHandleApi) FindMonthMethodCategorySuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.category.FindMonthMethodCategorySuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category success methods by merchant", zap.Error(err))
		return category_errors.ErrApiFindMonthMethodCategorySuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category success methods by merchant
// @Tags Category
// @Description Retrieve yearly category success methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/yearly-method-success/{merchant_id} [get]
func (h *categoryHandleApi) FindYearMethodCategorySuccessByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.category.FindYearMethodCategorySuccessByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category success methods by merchant", zap.Error(err))
		return category_errors.ErrApiFindYearMethodCategorySuccessByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find monthly category failed methods by merchant
// @Tags Category
// @Description Retrieve monthly category failed methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/monthly-method-failed/{merchant_id} [get]
func (h *categoryHandleApi) FindMonthMethodCategoryFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.MonthMethodCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.category.FindMonthMethodCategoryFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch monthly category failed methods by merchant", zap.Error(err))
		return category_errors.ErrApiFindMonthMethodCategoryFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponsesMonthMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find yearly category failed methods by merchant
// @Tags Category
// @Description Retrieve yearly category failed methods by merchant ID
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseCategoryYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve data"
// @Router /api/category/merchant/yearly-method-failed/{merchant_id} [get]
func (h *categoryHandleApi) FindYearMethodCategoryFailedByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid merchant ID", http.StatusBadRequest)
	}

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil || year <= 0 {
		return response.NewApiErrorResponse(c, "error", "invalid year", http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	req := &pb.YearMethodCategoryByMerchantRequest{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.category.FindYearMethodCategoryFailedByMerchant(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch yearly category failed methods by merchant", zap.Error(err))
		return category_errors.ErrApiFindYearMethodCategoryFailedByMerchant(c)
	}

	response := h.mapping.ToApiResponseYearMethod(res)
	return c.JSON(http.StatusOK, response)
}

// @Security Bearer
// @Summary Find Category by ID
// @Tags Category
// @Description Retrieve a Category by ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.ApiResponseCategory "Category data"
// @Failure 400 {object} response.ErrorResponse "Invalid Category ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Category data"
// @Router /api/category/{id} [get]
func (h *categoryHandleApi) FindById(c echo.Context) error {
	CategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || CategoryID <= 0 {
		return category_errors.ErrApiCategoryInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Category", zap.Error(err))
		return category_errors.ErrApiCategoryNotFound(c)
	}

	so := h.mapping.ToApiResponseCategory(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active Categorys
// @Tags Category
// @Description Retrieve a list of active Categorys
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationCategoryDeleteAt "List of active Categorys"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Category data"
// @Router /api/category/active [get]
func (h *categoryHandleApi) FindByActive(c echo.Context) error {
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

	req := &pb.FindAllCategoryRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.category.FindByActive(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch active Categorys", zap.Error(err))
		return category_errors.ErrApiFailedFindActive(c)
	}

	so := h.mapping.ToApiResponsePaginationCategoryDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve trashed Categorys
// @Tags Category
// @Description Retrieve a list of trashed Category records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationCategoryDeleteAt "List of trashed Category data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Category data"
// @Router /api/category/trashed [get]
func (h *categoryHandleApi) FindByTrashed(c echo.Context) error {
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

	req := &pb.FindAllCategoryRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.category.FindByTrashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch trashed Categorys", zap.Error(err))
		return category_errors.ErrApiFailedFindTrashed(c)
	}

	so := h.mapping.ToApiResponsePaginationCategoryDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create an existing Category
// @Tags Category
// @Description Create an existing Category record with the provided details
// @Accept json
// @Produce json
// @Param CreateCategoryRequest body requests.CreateCategoryRequest true "Create Category request"
// @Success 200 {object} response.ApiResponseCategory "Successfully created Category"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create Category"
// @Router /api/category/create [post]
func (h *categoryHandleApi) Create(c echo.Context) error {
	var req requests.CreateCategoryRequest
	if err := c.Bind(&req); err != nil {
		return category_errors.ErrApiBindCreateCategory(c)
	}

	if err := req.Validate(); err != nil {
		return category_errors.ErrApiValidateCreateCategory(c)
	}

	reqPb := &pb.CreateCategoryRequest{
		Name: req.Name,
	}

	ctx := c.Request().Context()

	res, err := h.category.Create(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to create Category", zap.Error(err))
		return category_errors.ErrApiFailedCreateCategory(c)
	}

	so := h.mapping.ToApiResponseCategory(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update an existing Category
// @Tags Category
// @Description Update an existing Category record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param UpdateCategoryRequest body requests.UpdateCategoryRequest true "Update Category request"
// @Success 200 {object} response.ApiResponseCategory "Successfully updated Category"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update Category"
// @Router /api/category/update/{id} [post]
func (h *categoryHandleApi) Update(c echo.Context) error {
	CategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || CategoryID <= 0 {
		return category_errors.ErrApiCategoryInvalidId(c)
	}

	var req requests.UpdateCategoryRequest
	if err := c.Bind(&req); err != nil {
		return category_errors.ErrApiBindUpdateCategory(c)
	}

	if err := req.Validate(); err != nil {
		return category_errors.ErrApiValidateUpdateCategory(c)
	}

	reqPb := &pb.UpdateCategoryRequest{
		CategoryId: int32(CategoryID),
		Name:       req.Name,
	}

	ctx := c.Request().Context()

	res, err := h.category.Update(ctx, reqPb)
	if err != nil {
		h.logger.Debug("Failed to update Category", zap.Error(err))
		return category_errors.ErrApiFailedUpdateCategory(c)
	}

	so := h.mapping.ToApiResponseCategory(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve a trashed Category
// @Tags Category
// @Description Retrieve a trashed Category record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.ApiResponseCategoryDeleteAt "Successfully retrieved trashed Category"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed Category"
// @Router /api/category/trashed/{id} [get]
func (h *categoryHandleApi) Trashed(c echo.Context) error {
	CategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || CategoryID <= 0 {
		return category_errors.ErrInvalidCategoryId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Category", zap.Error(err))
		return category_errors.ErrApiFailedTrashedCategory(c)
	}

	so := h.mapping.ToApiResponseCategoryDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore a trashed Category
// @Tags Category
// @Description Restore a trashed Category record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.ApiResponseCategoryDeleteAt "Successfully restored Category"
// @Failure 400 {object} response.ErrorResponse "Invalid Category ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Category"
// @Router /api/category/restore/{id} [post]
func (h *categoryHandleApi) Restore(c echo.Context) error {
	CategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || CategoryID <= 0 {
		return category_errors.ErrApiCategoryInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Category", zap.Error(err))
		return category_errors.ErrApiFailedRestoreCategory(c)
	}

	so := h.mapping.ToApiResponseCategoryDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete a Category
// @Tags Category
// @Description Permanently delete a Category record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.ApiResponseCategoryDelete "Successfully deleted Category record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Category"
// @Router /api/category/delete/{id} [post]
func (h *categoryHandleApi) DeletePermanent(c echo.Context) error {
	CategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || CategoryID <= 0 {
		return category_errors.ErrApiCategoryInvalidId(c)
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Category permanently", zap.Error(err))
		return category_errors.ErrApiFailedDeletePermanent(c)
	}

	so := h.mapping.ToApiResponseCategoryDelete(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore all trashed Categorys
// @Tags Category
// @Description Restore all trashed Category records.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseCategoryAll "Successfully restored all Categorys"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Categorys"
// @Router /api/category/restore/all [post]
func (h *categoryHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.category.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Categorys", zap.Error(err))
		return category_errors.ErrApiFailedRestoreAll(c)
	}

	so := h.mapping.ToApiResponseCategoryAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Permanently delete all trashed Categorys
// @Tags Category
// @Description Permanently delete all trashed Category records.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseCategoryDelete "Successfully deleted all Category records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Categorys"
// @Router /api/category/delete/all [post]
func (h *categoryHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.category.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Categorys permanently", zap.Error(err))
		return category_errors.ErrApiFailedDeleteAll(c)
	}

	so := h.mapping.ToApiResponseCategoryAll(res)

	return c.JSON(http.StatusOK, so)
}
