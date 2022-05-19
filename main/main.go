//  Skill assessment service:
//   version: 0.0.1
//   title: skill-assessment-service
//  Schemes: http, https
//  Host: localhost:8084
//  BasePath: /
//  Produces:
//    - application/json
//
// securityDefinitions:
//  apiKey:
//    type: apiKey
//    in: header
//    name: authorization
// swagger:meta
package main

import (
	"context"
	"log"

	config "github.com/sandy0786/skill-assessment-service/configuration"
	categoryDao "github.com/sandy0786/skill-assessment-service/dao/category"
	questionDao "github.com/sandy0786/skill-assessment-service/dao/question"
	userDao "github.com/sandy0786/skill-assessment-service/dao/user"
	database "github.com/sandy0786/skill-assessment-service/database"
	endpoint "github.com/sandy0786/skill-assessment-service/endpoint"
	server "github.com/sandy0786/skill-assessment-service/server"
	service "github.com/sandy0786/skill-assessment-service/service"
	transport "github.com/sandy0786/skill-assessment-service/transport"

	"github.com/go-playground/validator"
)

func init() {
	log.Println("inside init")
	transport.Validate = validator.New()
}

func main() {

	ctx := context.Background()

	// config
	configobj := config.NewConfigObject()
	configobj.LoadConfiguration()

	// log.Println("config > ", configobj.GetConfigDetails())
	dbDetails := configobj.GetConfigDetails().DatabaseDetails
	// connect to db
	dbObj := database.NewMongoObj(dbDetails.Host, dbDetails.Port, dbDetails.User, dbDetails.Password, dbDetails.Name, dbDetails.ConnectionString)
	err := dbObj.Connect()
	if err != nil {
		log.Fatal("Db connection error : ", err)
	}

	empDao := userDao.NewUserDAO(dbObj, "user")
	qsnDao := questionDao.NewQuestionDAO(dbObj, "questions")
	ctgDao := categoryDao.NewCategoryDAO(dbObj, "categories")

	// create service
	userSrv := service.NewUserService(configobj, empDao)
	qsnSrv := service.NewQuestionService(configobj, qsnDao)
	ctgSrv := service.NewCategoryService(configobj, ctgDao)

	errChan := make(chan error)

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		StatusEndpoint:      endpoint.MakeStatusEndpoint(userSrv),
		AddUserEndpoint:     endpoint.MakeAddUserEndpoint(userSrv),
		GetAllUsersEndpoint: endpoint.MakeGetAllUsersEndpoint(userSrv),
		// GetEmployeeByIdEndpoint: endpoint.MakeGetEmployeeByIdEndpoint(userSrv),
		AddQuestionEndpoint:         endpoint.MakeAddQuestionEndpoint(qsnSrv),
		AddMultipleQuestionEndpoint: endpoint.MakeAddMultipleQuestionsEndpoint(qsnSrv),
		GetAllQuestionsEndpoint:     endpoint.MakeGetAllQuestionsEndpoint(qsnSrv),
		AddCategoryEndpoint:         endpoint.MakeAddCategoryEndpoint(ctgSrv),
		GetAllCategoriesEndpoint:    endpoint.MakeGetAllCategoriesEndpoint(ctgSrv),
	}

	// HTTP transport
	srv := server.CreateNewServer(configobj.GetConfigDetails(), server.NewHTTPServer(ctx, endpoints))
	go func() {
		errChan <- srv.ListenAndServe()
	}()

	log.Println("Main: main: Microservice started")
	log.Println("msg", <-errChan)
}
