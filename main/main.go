package main

import (
	"context"
	"log"

	config "github.com/sandy0786/skill-assessment-service/configuration"
	questionDao "github.com/sandy0786/skill-assessment-service/dao/question"
	userDao "github.com/sandy0786/skill-assessment-service/dao/user"
	database "github.com/sandy0786/skill-assessment-service/database"
	endpoint "github.com/sandy0786/skill-assessment-service/endpoint"
	server "github.com/sandy0786/skill-assessment-service/server"
	service "github.com/sandy0786/skill-assessment-service/service"
	transport "github.com/sandy0786/skill-assessment-service/transport"

	"github.com/go-playground/validator"
	// "github.com/ArthurHlt/go-eureka-client/eureka"
	// "github.com/go-playground/validator"
	// "github.com/ArthurHlt/go-eureka-client/eureka"
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

	// create service
	userSrv := service.NewUserService(configobj, empDao)
	qsnSrv := service.NewQuestionService(configobj, qsnDao)

	errChan := make(chan error)

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		StatusEndpoint:      endpoint.MakeStatusEndpoint(userSrv),
		AddUserEndpoint:     endpoint.MakeAddUserEndpoint(userSrv),
		GetAllUsersEndpoint: endpoint.MakeGetAllUsersEndpoint(userSrv),
		// GetEmployeeByIdEndpoint: endpoint.MakeGetEmployeeByIdEndpoint(userSrv),
		AddQuestionEndpoint:     endpoint.MakeAddQuestionEndpoint(qsnSrv),
		GetAllQuestionsEndpoint: endpoint.MakeGetAllQuestionsEndpoint(qsnSrv),
	}

	// HTTP transport
	srv := server.CreateNewServer(configobj.GetConfigDetails(), server.NewHTTPServer(ctx, endpoints))
	go func() {
		errChan <- srv.ListenAndServe()
	}()

	log.Println("Main: main: Microservice started")

	// client := eureka.NewClient([]string{
	// 	"http://127.0.0.1:8761/eureka", //From a spring boot based eureka server
	// 	// add others servers here
	// })
	// instance := eureka.NewInstanceInfo("localhost", "test-service", "127.0.0.1", 8084, 30, false) //Create a new instance to register
	// instance.Metadata = &eureka.MetaData{
	// 	Map: make(map[string]string),
	// }
	// instance.Metadata.Map["foo"] = "bar"        //add metadata for example
	// client.RegisterInstance("test-service", instance) // Register new instance in your eureka(s)
	// applications, _ := client.GetApplications() // Retrieves all applications from eureka server(s)
	// log.Println("applications : ", applications)
	// client.GetApplication(instance.App)                 // retrieve the application "test"
	// client.GetInstance(instance.App, instance.HostName) // retrieve the instance from "test.com" inside "test"" app

	// go func() {
	// 	// send heartbeat every 30sec
	// client.SendHeartbeat(instance.App, instance.HostName) // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
	// }()

	// s, err := scheduler.NewScheduler(1)
	// if err != nil {
	// 	log.Println(err)
	// }

	// // s.Every().Do(client.SendHeartbeat, instance.App, instance.HostName)
	// s.Every().Do(test, client, instance)

	log.Println("msg", <-errChan)
}

// func test(client *eureka.Client, instance *eureka.InstanceInfo) {
// 	log.Println("print")
// 	applications, _ := client.GetApplications() // Retrieves all applications from eureka server(s)
// 	log.Println("applications : ", applications)
// 	client.GetApplication(instance.App) // retrieve the application "test"
// 	client.GetInstance(instance.App, instance.HostName)
// 	client.SendHeartbeat(instance.App, instance.HostName)
// }
