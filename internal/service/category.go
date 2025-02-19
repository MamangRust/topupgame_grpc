package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
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

func (s *categoryService) FindAll(page int, pageSize int, search string) ([]*response.CategoryResponse, int, *response.ErrorResponse) {
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

	res, totalRecords, err := s.categoryRepository.FindAllCategories(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to fetch category",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch category records",
		}
	}

	s.logger.Debug("Successfully fetched category",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToCategorysResponse(res)

	return so, totalRecords, nil
}

func (s *categoryService) FindById(id int) (*response.CategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching category by ID", zap.Int("id", id))

	res, err := s.categoryRepository.FindById(id)

	if err != nil {
		s.logger.Error("Failed to fetch category record by ID", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch category record by ID",
		}
	}

	s.logger.Debug("Successfully fetched category", zap.Int("id", id))

	so := s.mapping.ToCategoryResponse(res)

	return so, nil
}

func (s *categoryService) FindByActive(page int, pageSize int, search string) ([]*response.CategoryResponseDeleteAt, int, *response.ErrorResponse) {
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

	res, totalRecords, err := s.categoryRepository.FindByActiveCategories(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to fetch active category",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch category records",
		}
	}

	s.logger.Debug("Successfully fetched active category",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToCategorysResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *categoryService) FindByTrashed(page int, pageSize int, search string) ([]*response.CategoryResponseDeleteAt, int, *response.ErrorResponse) {
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

	res, totalRecords, err := s.categoryRepository.FindByTrashedCategory(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to fetch trashed category",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch category records",
		}
	}

	s.logger.Debug("Successfully fetched trashed category",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToCategorysResponseDeleteAt(res)

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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create category record",
		}
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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update category record",
		}
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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed category record",
		}
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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore category record",
		}
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

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete category record",
		}
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
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all categories: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all categories")
	return true, nil
}

func (s *categoryService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all categories")

	_, err := s.categoryRepository.DeleteAllCategoriesPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all categories", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all categories: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted all categories permanently")
	return true, nil
}
