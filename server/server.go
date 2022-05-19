package server

import (
	"context"
	"log"
	"net/http"
	"strings"

	config "github.com/sandy0786/skill-assessment-service/configuration"
	constants "github.com/sandy0786/skill-assessment-service/constants"
	endpoint "github.com/sandy0786/skill-assessment-service/endpoint"
	transport "github.com/sandy0786/skill-assessment-service/transport"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//NewHTTPServer is an exported function
// added options for functions like validations before decoding the requests
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints, options ...httptransport.ServerOption) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleWare)
	log.Println("Inside Server implementation")

	Opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(transport.ErrorEncoder),
		httptransport.ServerErrorEncoder(transport.QuestionErrorEncoder),
		httptransport.ServerErrorEncoder(transport.CategoryErrorEncoder),
	}

	r.Methods("GET").Path(constants.STATUS_PATH).Handler(httptransport.NewServer(
		endpoints.StatusEndpoint,
		transport.DecodeStatusRequest,
		transport.EncodeStatusResponse,
		Opts[0],
	))

	r.Methods("POST").Path(constants.USER).Handler(httptransport.NewServer(
		endpoints.AddUserEndpoint,
		transport.DecodeAddUserRequest,
		transport.EncodeAddUserResponse,
		Opts[0],
	))

	r.Methods("GET").Path(constants.ALL_USERS).Handler(httptransport.NewServer(
		endpoints.GetAllUsersEndpoint,
		transport.DecodeGetAllUsersRequest,
		transport.EncodeGetAllUsersResponse,
		Opts[0],
	))

	r.Methods("POST").Path(constants.QUESTION).Handler(httptransport.NewServer(
		endpoints.AddQuestionEndpoint,
		transport.DecodeAddQuestionRequest,
		transport.EncodeAddQuestionResponse,
		Opts[1],
	))

	r.Methods("POST").Path(constants.ALL_QUESTIONS).Handler(httptransport.NewServer(
		endpoints.AddMultipleQuestionEndpoint,
		transport.DecodeAddMutlipleQuestionsRequest,
		transport.EncodeAddMultipleQuestionsResponse,
		Opts[1],
	))

	r.Methods("GET").Path(constants.ALL_QUESTIONS).Handler(httptransport.NewServer(
		endpoints.GetAllQuestionsEndpoint,
		transport.DecodeGetAllQuestionsRequest,
		transport.EncodeGetAllQuestionsResponse,
		Opts[1],
	))

	r.Methods("POST").Path(constants.CATEGORY).Handler(httptransport.NewServer(
		endpoints.AddCategoryEndpoint,
		transport.DecodeAddCategoryRequest,
		transport.EncodeAddCategoryResponse,
		Opts[2],
	))

	r.Methods("GET").Path(constants.ALL_CATEGORIES).Handler(httptransport.NewServer(
		endpoints.GetAllCategoriesEndpoint,
		transport.DecodeGetAllCategoriesRequest,
		transport.EncodeGetAllCategoriesResponse,
		Opts[2],
	))

	// r.Methods("GET").Path(constants.EMPLOYEE).Handler(httptransport.NewServer(
	// 	endpoints.GetAllEmployeesEndpoint,
	// 	transport.DecodeGetAllEmpRequest,
	// 	transport.EncodeGetAllEmpResponse,
	// 	Opts[0],
	// ))

	// r.Methods("GET").Path(constants.EMPLOYEE_ID).Handler(httptransport.NewServer(
	// 	endpoints.GetEmployeeByIdEndpoint,
	// 	transport.DecodeGetEmpByIdRequest,
	// 	transport.EncodeGetAllEmpResponse,
	// 	Opts[0],
	// ))

	fs := http.FileServer(http.Dir("./docs/"))
	r.PathPrefix(constants.DOC_PATH).Handler(http.StripPrefix(constants.DOC_PATH, fs))

	return r
}

func commonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.RequestURI, "docs") {
			w.Header().Set("Content-Type", "text/html")
		} else {
			w.Header().Set("Content-Type", "application/json")
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		// w.Header().Set("Access-Control-Expose-Headers", "Message,RowsInResult,Status,TimeStamp,Content-Disposition")

		next.ServeHTTP(w, r)
	})
}

//CreateNewServer is a function to return server
func CreateNewServer(c config.ConfigurationDetails, handler http.Handler) http.Server {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
	})

	newHandler := cors.Handler(handler)
	return http.Server{
		Addr:    ":" + c.ServerPort,
		Handler: newHandler,
	}
}
