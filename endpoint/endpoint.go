package endpoint

import (
	"context"
	"log"

	userRequest "github.com/sandy0786/skill-assessment-service/request/user"
	service "github.com/sandy0786/skill-assessment-service/service"
	transport "github.com/sandy0786/skill-assessment-service/transport"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints are exposed here
//Contains endpoints that map request from client to our service
type Endpoints struct {
	StatusEndpoint      endpoint.Endpoint
	AddEmployeeEndpoint endpoint.Endpoint
	// GetAllEmployeesEndpoint endpoint.Endpoint
	// GetEmployeeByIdEndpoint endpoint.Endpoint
}

//MakeStatusEndpoint returns response
func MakeStatusEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeStatusEndpoint")
		responseFromService, err := srv.GetServiceStatus(ctx)
		if err != nil {
			return &transport.StatusResponse{Status: responseFromService}, err
		}
		return &transport.StatusResponse{Status: responseFromService}, nil
	}
}

//MakeAddEmployeeEndpoint returns response
func MakeAddEmployeeEndpoint(srv service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("MakeAddEmployeeEndpoint")
		employeeRequest := request.(userRequest.UserRequest)
		responseFromService, err := srv.AddUser(ctx, employeeRequest)
		return responseFromService, err
	}
}

//MakeGetAllEmployeesEndpoint returns response
// func MakeGetAllEmployeesEndpoint(srv service.Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		log.Println("MakeGetAllEmployeesEndpoint")
// 		responseFromService, err := srv.GetAllEmployees(ctx)
// 		return responseFromService, err
// 	}
// }

// //MakeGetEmployeeByIdEndpoint returns response
// func MakeGetEmployeeByIdEndpoint(srv service.Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		log.Println("MakeGetEmployeeByIdEndpoint")
// 		id, err := strconv.ParseInt(request.(string), 10, 64)
// 		responseFromService, err := srv.GetEmployeeById(ctx, id)
// 		return responseFromService, err
// 	}
// }
