package server

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/casbin/casbin"
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
		httptransport.ServerErrorEncoder(transport.AuthErrorEncoder),
	}

	// swagger:route GET /api/health miscellaneous health
	// Health of the application
	//
	// responses:
	//  500: InternalServerErrorResponse
	//  200: HealthResponse
	r.Methods(http.MethodGet).Path(constants.STATUS_PATH).Handler(httptransport.NewServer(
		endpoints.StatusEndpoint,
		transport.DecodeStatusRequest,
		transport.EncodeStatusResponse,
		Opts[0],
	))

	// swagger:route POST /api/user admin UserRequest
	// Create new user
	//
	// Security:
	// 	- Bearer: []
	//
	// securityDefinitions:
	//   Bearer:
	//     type: apiKey
	//     name: Authorization
	//     in: header
	//
	// requests:
	// responses:
	//  409: ConflictErrorResponse
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  401: UnAuthenticatedAccessResponse
	// 	403: UnAuthorizedAccessResponse
	//  200: SuccessResponse
	r.Methods(http.MethodPost).Path(constants.USER).Handler(httptransport.NewServer(
		endpoints.AddUserEndpoint,
		transport.DecodeAddUserRequest,
		transport.EncodeAddUserResponse,
		Opts[0],
	))

	// swagger:route DELETE /api/user/{Username} admin DeleteUserRequest
	// Delete user
	//
	// Security:
	// 	- Bearer: []
	//
	// securityDefinitions:
	//   Bearer:
	//     type: apiKey
	//     name: Authorization
	//     in: header
	//
	// requests:
	// responses:
	//  404: NotFoundErrorResponse
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  401: UnAuthorizedAccessResponse
	//  200: UserDeleteSuccessResponse
	r.Methods(http.MethodDelete).Path(constants.USER_NAME).Handler(httptransport.NewServer(
		endpoints.DeleteUserEndpoint,
		transport.DecodeDeleteUserRequest,
		transport.EncodeDeleteUserResponse,
		Opts[0],
	))

	// swagger:route GET /api/users admin listUsers
	// Get all available users
	//
	// Security:
	// 	- Bearer: []
	//
	// securityDefinitions:
	//   Bearer:
	//     type: apiKey
	//     name: Authorization
	//     in: header
	//
	// responses:
	//  500: InternalServerErrorResponse
	//  404: NotFoundEmptyErrorResponse
	//  401: UnAuthorizedAccessResponse
	//  400: BadRequestErrorResponse
	//  200: UsersResponse
	r.Methods(http.MethodGet).Path(constants.ALL_USERS).Handler(httptransport.NewServer(
		endpoints.GetAllUsersEndpoint,
		transport.DecodeGetAllUsersRequest,
		transport.EncodeGetAllUsersResponse,
		Opts[0],
	))

	// swagger:route PUT /api/user/{Username}/revoke admin RevokeUserRequest
	// Revoke user
	//
	//     Security:
	//     - bearer
	//
	//     SecurityDefinitions:
	//     bearer:
	//          type: apiKey
	//          name: Authorization
	//          in: header
	//
	// requests:
	// responses:
	//  404: NotFoundErrorResponse
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  401: UnAuthorizedAccessResponse
	//  200: UserRevokeSuccessResponse
	r.Methods(http.MethodPut).Path(constants.REVOKE_USER).Handler(httptransport.NewServer(
		endpoints.RevokeUserEndpoint,
		transport.DecodeDeleteUserRequest,
		transport.EncodeDeleteUserResponse,
		Opts[0],
	))

	// swagger:route PUT /api/user/{Username}/password user ResetUserPasswordRequest
	// Revoke user
	//
	// Security:
	// 	- Bearer: []
	//
	// securityDefinitions:
	//   Bearer:
	//     type: apiKey
	//     name: Authorization
	//     in: header
	//
	// responses:
	//  404: NotFoundErrorResponse
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  401: UnAuthorizedAccessResponse
	//  200: ResetUserPasswordRequest
	r.Methods(http.MethodPut).Path(constants.RESET_PASSWORD).Handler(httptransport.NewServer(
		endpoints.ResetPasswordEndpoint,
		transport.DecodePasswordResetRequest,
		transport.EncodePasswordResetRequest,
		Opts[0],
	))

	// swagger:route POST /api/user/login user LoginRequest
	// Login with username and password
	//
	// requests:
	// responses:
	//  401: UnAuthorizedResponse
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  200: LoginResponse
	r.Methods(http.MethodPost).Path(constants.LOGIN).Handler(httptransport.NewServer(
		endpoints.LoginEndpoint,
		transport.DecodeAuthRequest,
		transport.EncodeAuthResponse,
		Opts[3],
	))

	// swagger:route POST /api/question/{Category} questions QuestionRequest
	// Add new question based on category
	//
	// Security:
	// - bearer
	// SecurityDefinitions:
	// bearer:
	//      type: apiKey
	//      name: Authorization
	//      in: header
	// responses:
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  409: ConflictErrorResponse
	//  200: SuccessResponse
	r.Methods(http.MethodPost).Path(constants.QUESTION).Handler(httptransport.NewServer(
		endpoints.AddQuestionEndpoint,
		transport.DecodeAddQuestionRequest,
		transport.EncodeAddQuestionResponse,
		Opts[1],
	))

	// swagger:route POST /api/questions/{Category} questions QuestionsRequest
	// Fetch all categories
	//
	// Security:
	// - bearer
	// SecurityDefinitions:
	// bearer:
	//      type: apiKey
	//      name: Authorization
	//      in: header
	// responses:
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  200: SuccessResponse
	r.Methods(http.MethodPost).Path(constants.ALL_QUESTIONS).Handler(httptransport.NewServer(
		endpoints.AddMultipleQuestionEndpoint,
		transport.DecodeAddMutlipleQuestionsRequest,
		transport.EncodeAddMultipleQuestionsResponse,
		Opts[1],
	))

	// swagger:route GET /api/questions/{Category} questions GetQuestionRequest
	// Add new question based on category
	//
	// Security:
	// - bearer
	// SecurityDefinitions:
	// bearer:
	//      type: apiKey
	//      name: Authorization
	//      in: header
	// responses:
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  409: ConflictErrorResponse
	//  200: QuestionsResponse
	r.Methods(http.MethodGet).Path(constants.ALL_QUESTIONS).Handler(httptransport.NewServer(
		endpoints.GetAllQuestionsEndpoint,
		transport.DecodeGetAllQuestionsRequest,
		transport.EncodeGetAllQuestionsResponse,
		Opts[1],
	))

	// swagger:route POST /api/category category CategoryRequest
	// Add new category
	//
	// Security:
	// - bearer
	// SecurityDefinitions:
	// bearer:
	//      type: apiKey
	//      name: Authorization
	//      in: header
	// responses:
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  409: ConflictErrorResponse
	//  200: SuccessResponse
	r.Methods(http.MethodPost).Path(constants.CATEGORY).Handler(httptransport.NewServer(
		endpoints.AddCategoryEndpoint,
		transport.DecodeAddCategoryRequest,
		transport.EncodeAddCategoryResponse,
		Opts[2],
	))

	// swagger:route GET /api/categories category listCategories
	// Fetch all categories
	//
	// Security:
	// - bearer
	// SecurityDefinitions:
	// bearer:
	//      type: apiKey
	//      name: Authorization
	//      in: header
	// responses:
	//  500: InternalServerErrorResponse
	//  200: CategoriesResponse
	r.Methods(http.MethodGet).Path(constants.ALL_CATEGORIES).Handler(httptransport.NewServer(
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

		next.ServeHTTP(w, r)
		// Authorizer()
	})
}

func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// role := session.GetString(r, "role")
			// if err != nil {
			//     writeError(http.StatusInternalServerError, "ERROR", w, err)
			//     return
			// }

			// if role == "" {
			// role = "anonymous"
			// }
			// setup casbin auth rules
			authEnforcer, err := casbin.NewEnforcerSafe("./configuration/conf/auth_model.conf", "./configuration/conf/policy.csv")
			if err != nil {
				log.Fatal(err)
			}

			log.Println("authEnforcer >> ", authEnforcer)

			log.Println("authorizer >>>>>")
			role := "admin"

			// if it's a member, check if the user still exists
			// if role == "member" {
			// 	uid, err := session.GetInt(r, "userID")
			// 	if err != nil {
			// 		writeError(http.StatusInternalServerError, "ERROR", w, err)
			// 		return
			// 	}
			// 	exists := users.Exists(uid)
			// 	if !exists {
			// 		writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
			// 		return
			// 	}
			// }

			// casbin rule enforcing
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				// writeError(http.StatusInternalServerError, "ERROR", w, err)
				log.Println("1 >> ", err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				// writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
				log.Println("2 >> ", err)
				return
			}
		}

		return http.HandlerFunc(fn)
	}
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
