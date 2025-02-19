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
	routerCategory.POST("", categoryHandler.Create)
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
		if errors.Is(err, echo.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status:  "error",
				Message: "Unauthorized",
			})
		}

		h.logger.Debug("Failed to fetch Category records", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Category records",
		})
	}

	so := h.mapping.ToApiResponsePaginationCategory(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find Category by ID
// @Tags Category
// @Description Retrieve a Category by ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} pb.ApiResponseCategory "Category data"
// @Failure 400 {object} response.ErrorResponse "Invalid Category ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve Category data"
// @Router /api/category/{id} [get]
func (h *categoryHandleApi) FindById(c echo.Context) error {
	CategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || CategoryID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Category ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.FindById(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to fetch Category", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Category",
		})
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Categorys",
		})
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Categorys",
		})
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
	var req pb.CreateCategoryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	ctx := c.Request().Context()

	res, err := h.category.Create(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to create Category", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create Category",
		})
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Category ID",
		})
	}

	var req pb.UpdateCategoryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	req.CategoryId = int32(CategoryID)

	ctx := c.Request().Context()

	res, err := h.category.Update(ctx, &req)
	if err != nil {
		h.logger.Debug("Failed to update Category", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update Category",
		})
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Category ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.Trashed(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to trash Category", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Category",
		})
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Category ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.Restore(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to restore Category", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Category",
		})
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
// @Success 200 {object} pb.ApiResponseCategoryDelete "Successfully deleted Category record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Category"
// @Router /api/category/delete/{id} [post]
func (h *categoryHandleApi) DeletePermanent(c echo.Context) error {
	CategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil || CategoryID <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid Category ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCategoryRequest{
		Id: int32(CategoryID),
	}

	res, err := h.category.DeletePermanent(ctx, req)
	if err != nil {
		h.logger.Debug("Failed to delete Category permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Category permanently",
		})
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
// @Success 200 {object} pb.ApiResponseCategoryAll "Successfully restored all Categorys"
// @Failure 500 {object} response.ErrorResponse "Failed to restore Categorys"
// @Router /api/category/restore/all [post]
func (h *categoryHandleApi) RestoreAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.category.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to restore all Categorys", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Categorys",
		})
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
// @Success 200 {object} pb.ApiResponseCategoryDelete "Successfully deleted all Category records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete Categorys"
// @Router /api/category/delete/all [post]
func (h *categoryHandleApi) DeleteAllPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.category.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Debug("Failed to delete all Categorys permanently", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Categorys permanently",
		})
	}

	so := h.mapping.ToApiResponseCategoryAll(res)

	return c.JSON(http.StatusOK, so)
}
