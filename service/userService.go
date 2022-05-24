package service

import (
	"context"
	"log"
	"net/http"
	"time"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	userDao "github.com/sandy0786/skill-assessment-service/dao/user"
	userDocument "github.com/sandy0786/skill-assessment-service/documents/user"
	userDto "github.com/sandy0786/skill-assessment-service/dto/user"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	successResponse "github.com/sandy0786/skill-assessment-service/response/success"
	userResponse "github.com/sandy0786/skill-assessment-service/response/user"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jinzhu/copier"
)

type UserService interface {
	GetServiceStatus(context.Context) (string, error)
	AddUser(context.Context, userRequest.UserRequest) (successResponse.SuccessResponse, error)
	GetAllUsers(context.Context) ([]userResponse.UserResponse, error)
	DeleteUserByUserName(context.Context, string) (successResponse.SuccessResponse, error)
	RevokeUserByUserName(context.Context, string) (successResponse.SuccessResponse, error)
	ResetUserPassword(context.Context, userDto.UserDTO) (successResponse.SuccessResponse, error)
}

// service for druid
type userService struct {
	testServiceConfig configuration.ConfigurationInterface
	dao               userDao.UserDAO
}

func NewUserService(c configuration.ConfigurationInterface, dao userDao.UserDAO) *userService {
	return &userService{
		testServiceConfig: c,
		dao:               dao,
	}
}

func (t *userService) GetServiceStatus(ctx context.Context) (string, error) {
	log.Println("Inside getServiceStatus")
	return `UP`, nil
}

func (t *userService) AddUser(ctx context.Context, userRequest userRequest.UserRequest) (successResponse.SuccessResponse, error) {
	log.Println("Inside Add user")
	// var userResponse userResponse.UserSuccessResponse
	var user userDocument.User
	copier.Copy(&user, &userRequest)
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	user.ID = primitive.NewObjectID()
	user.Active = true
	userCreated, err := t.dao.Save(&user)
	// copier.Copy(&userResponse, &userRequest)
	if userCreated {
		return successResponse.SuccessResponse{
			Status:    http.StatusOK,
			Message:   "User created successfully",
			TimeStamp: time.Now().UTC().String(),
		}, err
	}
	return successResponse.SuccessResponse{}, err
}

func (t *userService) GetAllUsers(context.Context) ([]userResponse.UserResponse, error) {
	log.Println("Inside GetAllusers")
	var userResponses []userResponse.UserResponse
	users, err := t.dao.FindAll()
	copier.Copy(&userResponses, &users)
	return userResponses, err
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
