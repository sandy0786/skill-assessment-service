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
	roleDao "github.com/sandy0786/skill-assessment-service/dao/role"
	userDao "github.com/sandy0786/skill-assessment-service/dao/user"
	userDocument "github.com/sandy0786/skill-assessment-service/documents/user"
	userDto "github.com/sandy0786/skill-assessment-service/dto/user"
	globalErr "github.com/sandy0786/skill-assessment-service/errors"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	successResponse "github.com/sandy0786/skill-assessment-service/response/success"
	userResponse "github.com/sandy0786/skill-assessment-service/response/user"
	userValidation "github.com/sandy0786/skill-assessment-service/validations/user"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetServiceStatus(context.Context) (string, error)
	AddUser(context.Context, userRequest.UserRequest) (successResponse.SuccessResponse, error)
	GetAllUsers(context.Context, *http.Request) (userResponse.UserResults, error)
	DeleteUserByUserName(context.Context, string) (successResponse.SuccessResponse, error)
	RevokeUserByUserName(context.Context, string) (successResponse.SuccessResponse, error)
	ResetUserPassword(context.Context, userDto.UserDTO) (successResponse.SuccessResponse, error)
}

// service for druid
type userService struct {
	testServiceConfig configuration.ConfigurationInterface
	dao               userDao.UserDAO
	roleDao           roleDao.RoleDAO
	validator         userValidation.UserValidator
}

func NewUserService(c configuration.ConfigurationInterface, dao userDao.UserDAO, roleDao roleDao.RoleDAO, validator userValidation.UserValidator) *userService {
	return &userService{
		testServiceConfig: c,
		dao:               dao,
		roleDao:           roleDao,
		validator:         validator,
	}
}

func (t *userService) GetServiceStatus(ctx context.Context) (string, error) {
	log.Println("Inside getServiceStatus")
	return `UP`, nil
}

func (t *userService) AddUser(ctx context.Context, userRequest userRequest.UserRequest) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add user")
	// var userResponse userResponse.UserSuccessResponse
	isRoleValid := t.roleDao.ValidateRole(userRequest.Role)

	if !isRoleValid {
		return successResponse.SuccessResponse{}, globalErr.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusBadRequest,
			Message:   "Invalid role",
		}
	}

	var user userDocument.User
	copier.Copy(&user, &userRequest)
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	user.ID = primitive.NewObjectID()
	user.Active = true
	user.Role, _ = primitive.ObjectIDFromHex(userRequest.Role)
	userCreated, err := t.dao.Save(&user)

	if userCreated {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "User created successfully",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}

func (t *userService) GetAllUsers(_ context.Context, r *http.Request) (userResponse.UserResults, error) {
	log.Println("Inside GetAllusers")
	var userResults userResponse.UserResults
	var userResponses []userResponse.UserResponse

	isValid, validationError := t.validator.ValidateGetAllUsersRequest(r)
	if !isValid {
		return userResponse.UserResults{}, validationError
	}

	queryParamValues := r.URL.Query()
	pageQueryParam := queryParamValues[constants.PAGE][0]
	lengthQueryParam := queryParamValues[constants.PAGE_SIZE][0]

	page, err := strconv.ParseInt(pageQueryParam, 10, 64)
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

	start := pageSize * (page - 1)

	var sortBy string
	if len(queryParamValues[constants.SORT_BY]) > 0 {
		sortBy = queryParamValues[constants.SORT_BY][0]
	}

	// create pagination object
	pagination := userDto.UserPaginationDTO{
		Search:  queryParamValues[constants.SEARCH][0],
		Start:   &start,
		Length:  &pageSize,
		SortBy:  sortBy,
		OrderBy: orderBy,
	}

	// get all users
	users, err := t.dao.FindAll(pagination.Start, pagination.Length, pagination.Search, pagination.SortBy, pagination.OrderBy)
	if err != nil {
		return userResults, err
	}

	// loop through each user and process
	for _, user := range users {
		userResponse := userResponse.UserResponse{
			Username:  user.Username,
			Email:     user.Email,
			Active:    user.Active,
			Role:      user.Role.Hex(),
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		}
		userResponses = append(userResponses, userResponse)
	}

	// get total count
	totalCount, err := t.dao.GetCount(pagination.Search)

	userResults.Data = userResponses
	userResults.TotalRecords = totalCount
	return userResults, err
}

func (t *userService) DeleteUserByUserName(_ context.Context, username string) (successResponse.SuccessResponse, error) {
	log.Println("Inside DeleteUserByUserName")
	userDeleted, err := t.dao.DeleteByUserName(username)
	if userDeleted {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "User deleted successfully",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}

func (t *userService) RevokeUserByUserName(_ context.Context, username string) (successResponse.SuccessResponse, error) {
	log.Println("Inside RevokeUserByUserName")
	userDeleted, err := t.dao.RevokeByUserName(username)
	if userDeleted {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "User revoked successfully",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}

func (t *userService) ResetUserPassword(_ context.Context, userdto userDto.UserDTO) (successResponse.SuccessResponse, error) {
	log.Println("Inside ResetUserPassword ")
	resetSuccess, err := t.dao.ResetUserPassword(userdto.Username, userdto.OldPassword, userdto.NewPassword)
	if resetSuccess {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "Password reset success",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}
