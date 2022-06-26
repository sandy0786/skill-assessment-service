//  Skill Assessment Service:
//   version: 1.0.0
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
	authDao "github.com/sandy0786/skill-assessment-service/dao/auth"
	categoryDao "github.com/sandy0786/skill-assessment-service/dao/category"
	initDatabase "github.com/sandy0786/skill-assessment-service/dao/init"
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

	// Create and Get mongodb metadata
	metadata := initDatabase.InitializeMongoMetadataObject()

	// init database
	initDatabase.InitMongoDBCollections(dbObj, metadata)

	empDao := userDao.NewUserDAO(dbObj, "users")
	qsnDao := questionDao.NewQuestionDAO(dbObj, "questions")
	ctgDao := categoryDao.NewCategoryDAO(dbObj, "categories")
	authhDao := authDao.NewAuthDAO(dbObj, "users")

	// create service
	userSrv := service.NewUserService(configobj, empDao)
	qsnSrv := service.NewQuestionService(configobj, qsnDao)
	ctgSrv := service.NewCategoryService(configobj, ctgDao)
	authSrv := service.NewAuthService(configobj, authhDao)

	errChan := make(chan error)

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		StatusEndpoint:              endpoint.MakeStatusEndpoint(userSrv),
		AddUserEndpoint:             endpoint.MakeAddUserEndpoint(userSrv),
		DeleteUserEndpoint:          endpoint.MakeDeleteUserEndpoint(userSrv),
		GetAllUsersEndpoint:         endpoint.MakeGetAllUsersEndpoint(userSrv),
		RevokeUserEndpoint:          endpoint.MakeRevokeUserEndpoint(userSrv),
		ResetPasswordEndpoint:       endpoint.MakeResetPasswordEndpoint(userSrv),
		LoginEndpoint:               endpoint.MakeLoginEndpoint(authSrv),
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