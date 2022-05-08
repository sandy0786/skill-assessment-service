package service

import (
	"context"
	"log"

	configuration "github.com/sandy0786/skill-assessment-service/configuration"
	userDao "github.com/sandy0786/skill-assessment-service/dao/user"
	userDocument "github.com/sandy0786/skill-assessment-service/documents/user"
	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	userResponse "github.com/sandy0786/skill-assessment-service/response/user"

	"github.com/jinzhu/copier"
)

type UserService interface {
	GetServiceStatus(context.Context) (string, error)
	AddUser(context.Context, userRequest.UserRequest) (userResponse.UserResponse, error)
	// GetAllEmployees(context.Context) ([]employeeResponse.EmployeeResponse, error)
	// GetEmployeeById(context.Context, int64) (employeeResponse.EmployeeResponse, error)
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
	return `ok`, nil
}

func (t *userService) AddUser(ctx context.Context, userRequest userRequest.UserRequest) (userResponse.UserResponse, error) {
	log.Println("Inside Add user")
	var userResponse userResponse.UserResponse
	var user userDocument.User
	copier.Copy(&user, &userRequest)
	_, err := t.dao.Save(&user)
	copier.Copy(&userResponse, &userRequest)
	return userResponse, err
}

// func (t *testService) GetAllEmployees(context.Context) ([]employeeResponse.EmployeeResponse, error) {
// 	log.Println("Inside GetAllEmployees")
// 	var employeeResponses []employeeResponse.EmployeeResponse
// 	employees, err := t.dao.FindAll()
// 	copier.Copy(&employeeResponses, &employees)
// 	return employeeResponses, err
// }

// func (t *testService) GetEmployeeById(_ context.Context, id int64) (employeeResponse.EmployeeResponse, error) {
// 	log.Println("Inside GetEmployeeById : ", id)
// 	var employeeResponse employeeResponse.EmployeeResponse
// 	employee, err := t.dao.FindById(id)
// 	copier.Copy(&employeeResponse, &employee)
// 	return employeeResponse, err
// }
