package service

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	constants "github.com/sandy0786/skill-assessment-service/constants"
	categoryDao "github.com/sandy0786/skill-assessment-service/dao/category"
	categoryDocument "github.com/sandy0786/skill-assessment-service/documents/category"
	categoryDto "github.com/sandy0786/skill-assessment-service/dto/category"
	categoryRequest "github.com/sandy0786/skill-assessment-service/request/category"
	categoryResponse "github.com/sandy0786/skill-assessment-service/response/category"
	successResponse "github.com/sandy0786/skill-assessment-service/response/success"
	categoryValidation "github.com/sandy0786/skill-assessment-service/validations/category"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService interface {
	// GetServiceStatus(context.Context) (string, error)
	AddCategory(context.Context, categoryRequest.CategoryRequest) (successResponse.SuccessResponse, error)
	// AddMultipleQuestions(context.Context, []categoryRequest.CategoryRequest) (successResponse.SuccessResponse, error)
	GetAllCategories(context.Context, *http.Request) (categoryResponse.CategoryResults, error)
	// GetEmployeeById(context.Context, int64) (employeeResponse.EmployeeResponse, error)
}

// service for druid
type categoryService struct {
	questionServiceConfig configuration.ConfigurationInterface
	dao                   categoryDao.CategoryDAO
	validator             categoryValidation.CategoryValidator
}

func NewCategoryService(c configuration.ConfigurationInterface, dao categoryDao.CategoryDAO, validator categoryValidation.CategoryValidator) *categoryService {
	return &categoryService{
		questionServiceConfig: c,
		dao:                   dao,
		validator:             validator,
	}
}

func (t *categoryService) AddCategory(_ context.Context, categoryRequest categoryRequest.CategoryRequest) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add Category")

	// create collection with provided collection name
	// _, err := t.dao.CreateCollection(categoryRequest.CollectionName)
	// if err != nil {
	// 	return successResponse.SuccessResponse{}, err
	// }

	// update category details in db
	var category categoryDocument.Category
	copier.Copy(&category, &categoryRequest)
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()
	category.ID = primitive.NewObjectID()
	categoryCreated, err := t.dao.Save(&category)
	if !categoryCreated {
		return successResponse.SuccessResponse{}, err
	}

	return successResponse.SuccessResponse{
		Status:    http.StatusCreated,
		Message:   "Category Created successfully",
		TimeStamp: time.Now().UTC().String(),
	}, err
}

func (t *categoryService) GetAllCategories(_ context.Context, r *http.Request) (categoryResponse.CategoryResults, error) {
	log.Println("Inside GetAllCategories")

	isValid, validationError := t.validator.ValidateGetAllUsersRequest(r)
	if !isValid {
		return categoryResponse.CategoryResults{}, validationError
	}

	queryParamValues := r.URL.Query()

	startQueryParam := queryParamValues[constants.PAGE][0]
	lengthQueryParam := queryParamValues[constants.PAGE_SIZE][0]

	start, err := strconv.ParseInt(startQueryParam, 10, 64)
	if err != nil {
		log.Println("error wile converting page", err)
	}

	pageSize, err := strconv.ParseInt(lengthQueryParam, 10, 64)
	if err != nil {
		log.Println("error wile converting pageSize", err)
	}

	var orderBy = 1
	var orderByQueryParam string
	if len(queryParamValues[constants.ORDER_BY]) > 0 {
		orderByQueryParam = strings.ToLower(queryParamValues[constants.ORDER_BY][0])
	}

	if orderByQueryParam == constants.ASCENDING {
		orderBy = 1
	} else if orderByQueryParam == constants.DESCENDING {
		orderBy = -1
	}

	start--

	// create pagination object
	pagination := categoryDto.Pagination{
		Search:  queryParamValues[constants.SEARCH][0],
		Start:   &start,
		Length:  &pageSize,
		OrderBy: orderBy,
	}

	var categoryResponses []categoryResponse.CategoryResponse
	categories, err := t.dao.FindAll(pagination)

	for _, category := range categories {
		categoryResponse := categoryResponse.CategoryResponse{
			ID:        category.ID.Hex(),
			Category:  category.Category,
			CreatedAt: category.CreatedAt.Unix(),
			UpdatedAt: category.UpdatedAt.Unix(),
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	// get total count
	totalCount, err := t.dao.GetCount(pagination.Search)

	var categoryResults categoryResponse.CategoryResults
	categoryResults.Data = categoryResponses
	categoryResults.TotalRecords = totalCount
	return categoryResults, err
}
