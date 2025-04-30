package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/errors/category_errors"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
	logger             logger.LoggerInterface
	mapping            response_service.CategoryResponseMapper
}

func NewCategoryService(categoryRepository repository.CategoryRepository, logger logger.LoggerInterface, mapping response_service.CategoryResponseMapper) *categoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *categoryService) FindAll(request *requests.FindAllCategory) ([]*response.CategoryResponse, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching category",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.categoryRepository.FindAllCategories(request)
	if err != nil {
		s.logger.Error("Failed to fetch category",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, category_errors.ErrFailedFindAll
	}

	s.logger.Debug("Successfully fetched category",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToCategoriesResponse(res)

	return so, totalRecords, nil
}

func (s *categoryService) FindMonthAmountCategorySuccess(req *requests.MonthAmountCategoryRequest) ([]*response.MonthAmountCategorySuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category success amounts", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthAmountCategorySuccess(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category success amounts", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthAmountCategorySuccess
	}

	responses := s.mapping.ToCategoriesResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly category success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearAmountCategorySuccess(year int) ([]*response.YearAmountCategorySuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category success amounts", zap.Int("year", year))

	records, err := s.categoryRepository.FindYearAmountCategorySuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category success amounts", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearAmountCategorySuccess
	}

	responses := s.mapping.ToCategoriesResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly category success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthAmountCategoryFailed(req *requests.MonthAmountCategoryRequest) ([]*response.MonthAmountCategoryFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category failed amounts", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthAmountCategoryFailed(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category failed amounts", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthAmountCategoryFailed
	}

	responses := s.mapping.ToCategoriesResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly category failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearAmountCategoryFailed(year int) ([]*response.YearAmountCategoryFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category failed amounts", zap.Int("year", year))

	records, err := s.categoryRepository.FindYearAmountCategoryFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category failed amounts", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearAmountCategoryFailed
	}

	responses := s.mapping.ToCategoriesResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly category failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthMethodCategorySuccess(year int) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category success methods", zap.Int("year", year))

	records, err := s.categoryRepository.FindMonthMethodCategorySuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category success methods", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthMethodCategorySuccess
	}

	responses := s.mapping.ToCategoriesResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly category success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearMethodCategorySuccess(year int) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category success methods", zap.Int("year", year))

	records, err := s.categoryRepository.FindYearMethodCategorySuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category success methods", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearMethodCategorySuccess
	}

	responses := s.mapping.ToCategoriesResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly category success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthMethodCategoryFailed(year int) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category failed methods", zap.Int("year", year))

	records, err := s.categoryRepository.FindMonthMethodCategoryFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category failed methods", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthMethodCategoryFailed
	}

	responses := s.mapping.ToCategoriesResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly category failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearMethodCategoryFailed(year int) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category failed methods", zap.Int("year", year))

	records, err := s.categoryRepository.FindYearMethodCategoryFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category failed methods", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearMethodCategoryFailed
	}

	responses := s.mapping.ToCategoriesResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly category failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthAmountCategorySuccessById(req *requests.MonthAmountCategoryByIdRequest) ([]*response.MonthAmountCategorySuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category success amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthAmountCategorySuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category success amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthAmountCategorySuccessById
	}

	responses := s.mapping.ToCategoriesResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly category success amounts by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearAmountCategorySuccessById(req *requests.YearAmountCategoryByIdRequest) ([]*response.YearAmountCategorySuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category success amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearAmountCategorySuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category success amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearAmountCategorySuccessById
	}

	responses := s.mapping.ToCategoriesResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly category success amounts by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthAmountCategoryFailedById(req *requests.MonthAmountCategoryByIdRequest) ([]*response.MonthAmountCategoryFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category failed amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthAmountCategoryFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category failed amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthAmountCategoryFailedById
	}

	responses := s.mapping.ToCategoriesResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly category failed amounts by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearAmountCategoryFailedById(req *requests.YearAmountCategoryByIdRequest) ([]*response.YearAmountCategoryFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category failed amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearAmountCategoryFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category failed amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearAmountCategoryFailedById
	}

	responses := s.mapping.ToCategoriesResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly category failed amounts by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthMethodCategorySuccessById(req *requests.MonthMethodCategoryByIdRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category success methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthMethodCategorySuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category success methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthMethodCategorySuccessById
	}

	responses := s.mapping.ToCategoriesResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly category success methods by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearMethodCategorySuccessById(req *requests.YearMethodCategoryByIdRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category success methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearMethodCategorySuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category success methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearMethodCategorySuccessById
	}

	responses := s.mapping.ToCategoriesResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly category success methods by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthMethodCategoryFailedById(req *requests.MonthMethodCategoryByIdRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category failed methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthMethodCategoryFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category failed methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthMethodCategoryFailedById
	}

	responses := s.mapping.ToCategoriesResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly category failed methods by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearMethodCategoryFailedById(req *requests.YearMethodCategoryByIdRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category failed methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearMethodCategoryFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category failed methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearMethodCategoryFailedByMerchant
	}

	responses := s.mapping.ToCategoriesResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly category failed methods by ID",
		zap.Int("category_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthAmountCategorySuccessByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*response.MonthAmountCategorySuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category success amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthAmountCategorySuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category success amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthAmountCategorySuccessByMerchant
	}

	responses := s.mapping.ToCategoriesResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly category success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearAmountCategorySuccessByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*response.YearAmountCategorySuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category success amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearAmountCategorySuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category success amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearAmountCategorySuccessByMerchant
	}

	responses := s.mapping.ToCategoriesResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly category success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthAmountCategoryFailedByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*response.MonthAmountCategoryFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category failed amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthAmountCategoryFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category failed amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthAmountCategoryFailedByMerchant
	}

	responses := s.mapping.ToCategoriesResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly category failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearAmountCategoryFailedByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*response.YearAmountCategoryFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category failed amounts by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearAmountCategoryFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category failed amounts by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearAmountCategoryFailedByMerchant
	}

	responses := s.mapping.ToCategoriesResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly category failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthMethodCategorySuccessByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category success methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthMethodCategorySuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category success methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthMethodCategorySuccessByMerchant
	}

	responses := s.mapping.ToCategoriesResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly category success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearMethodCategorySuccessByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category success methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearMethodCategorySuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category success methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearMethodCategorySuccessByMerchant
	}

	responses := s.mapping.ToCategoriesResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly category success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindMonthMethodCategoryFailedByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly category failed methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindMonthMethodCategoryFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly category failed methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindMonthMethodCategoryFailedByMerchant
	}

	responses := s.mapping.ToCategoriesResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly category failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindYearMethodCategoryFailedByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly category failed methods by ID", zap.Any("request", req))

	records, err := s.categoryRepository.FindYearMethodCategoryFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly category failed methods by ID", zap.Error(err))
		return nil, category_errors.ErrFailedFindYearMethodCategoryFailedByMerchant
	}

	responses := s.mapping.ToCategoriesResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly category failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *categoryService) FindById(id int) (*response.CategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching category by ID", zap.Int("id", id))

	res, err := s.categoryRepository.FindById(id)

	if err != nil {
		s.logger.Error("Failed to fetch category record by ID", zap.Error(err))

		return nil, category_errors.ErrCategoryNotFoundRes
	}

	s.logger.Debug("Successfully fetched category", zap.Int("id", id))

	so := s.mapping.ToCategoryResponse(res)

	return so, nil
}

func (s *categoryService) FindByActive(request *requests.FindAllCategory) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching active category",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.categoryRepository.FindByActiveCategories(request)
	if err != nil {
		s.logger.Error("Failed to fetch active category",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, category_errors.ErrFailedFindActive
	}

	s.logger.Debug("Successfully fetched active category",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToCategoriesResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *categoryService) FindByTrashed(request *requests.FindAllCategory) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching trashed role",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.categoryRepository.FindByTrashedCategories(request)

	if err != nil {
		s.logger.Error("Failed to fetch trashed category",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, category_errors.ErrFailedFindTrashed
	}

	s.logger.Debug("Successfully fetched trashed category",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToCategoriesResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *categoryService) Create(request *requests.CreateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Create Category process",
		zap.String("name", request.Name),
	)

	res, err := s.categoryRepository.CreateCategory(request)

	if err != nil {
		s.logger.Error("Failed to create category",
			zap.String("name", request.Name),
			zap.Error(err),
		)

		return nil, category_errors.ErrFailedCreateCategory
	}

	so := s.mapping.ToCategoryResponse(res)

	s.logger.Debug("Create Category process completed",
		zap.String("name", request.Name),
		zap.Int("id", res.ID),
	)

	return so, nil
}

func (s *categoryService) Update(request *requests.UpdateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Update Category process",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
	)

	res, err := s.categoryRepository.UpdateCategory(request)

	if err != nil {
		s.logger.Error("Failed to update category",
			zap.Int("id", request.ID),
			zap.String("name", request.Name),
			zap.Error(err),
		)

		return nil, category_errors.ErrFailedCreateCategory
	}

	so := s.mapping.ToCategoryResponse(res)

	s.logger.Debug("Update Category process completed",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
	)

	return so, nil
}

func (s *categoryService) Trashed(id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting Trashed Category process",
		zap.Int("category_id", id),
	)

	res, err := s.categoryRepository.TrashedCategory(id)

	if err != nil {
		s.logger.Error("Failed to move Category to trash",
			zap.Int("category_id", id),
			zap.Error(err),
		)

		return nil, category_errors.ErrFailedTrashedCategory
	}

	so := s.mapping.ToCategoryResponseDeleteAt(res)

	s.logger.Debug("TrashedCategory process completed",
		zap.Int("category_id", id),
	)

	return so, nil
}

func (s *categoryService) Restore(id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting RestoreCategory process",
		zap.Int("category_id", id),
	)

	res, err := s.categoryRepository.RestoreCategory(id)

	if err != nil {
		s.logger.Error("Failed to restore category", zap.Error(err))

		return nil, category_errors.ErrFailedRestoreCategory
	}

	so := s.mapping.ToCategoryResponseDeleteAt(res)

	s.logger.Debug("RestoreCategory process completed",
		zap.Int("category_id", id),
	)

	return so, nil
}

func (s *categoryService) DeletePermanent(id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Starting DeleteCategoryPermanent process",
		zap.Int("category_id", id),
	)

	_, err := s.categoryRepository.DeleteCategoryPermanent(id)

	if err != nil {
		s.logger.Error("Failed to delete category permanently",
			zap.Int("category_id", id),
			zap.Error(err),
		)

		return false, category_errors.ErrFailedDeletePermanent
	}

	s.logger.Debug("DeleteCategoryPermanent process completed",
		zap.Int("category_id", id),
	)

	return true, nil
}

func (s *categoryService) RestoreAll() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all categories")

	_, err := s.categoryRepository.RestoreAllCategories()

	if err != nil {
		s.logger.Error("Failed to restore all categories", zap.Error(err))
		return false, category_errors.ErrFailedRestoreAll
	}

	s.logger.Debug("Successfully restored all categories")
	return true, nil
}

func (s *categoryService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all categories")

	_, err := s.categoryRepository.DeleteAllCategoriesPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all categories", zap.Error(err))
		return false, category_errors.ErrFailedDeleteAll
	}

	s.logger.Debug("Successfully deleted all categories permanently")
	return true, nil
}
