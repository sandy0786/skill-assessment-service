package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/casbin/casbin"
	"github.com/rs/cors"
	config "github.com/sandy0786/skill-assessment-service/configuration"
	constants "github.com/sandy0786/skill-assessment-service/constants"
	endpoint "github.com/sandy0786/skill-assessment-service/endpoint"
	globalErr "github.com/sandy0786/skill-assessment-service/errors"
	jwtP "github.com/sandy0786/skill-assessment-service/jwt"
	transport "github.com/sandy0786/skill-assessment-service/transport"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
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
		httptransport.ServerErrorEncoder(transport.RoleErrorEncoder),
		httptransport.ServerErrorEncoder(transport.QuestionTypeErrorEncoder),
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

	// swagger:route GET /api/users admin GetAllUsersRequest
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
	//  401: UnAuthorizedAccessResponse
	//  400: BadRequestErrorResponse
	//  200: UsersResponse
	r.Methods(http.MethodGet).Path(constants.ALL_USERS).Handler(httptransport.NewServer(
		endpoints.GetAllUsersEndpoint,
		transport.DecodeGetAllUsersRequest,
		transport.EncodeGetAllUsersResponse,
		Opts[0],
	))

	// swagger:route GET /api/user/roles admin listRoles
	// Get all roles
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
	//  401: UnAuthorizedAccessResponse
	//  200: RolesResponse
	r.Methods(http.MethodGet).Path(constants.All_ROLES).Handler(httptransport.NewServer(
		endpoints.GetAllRolesEndpoint,
		transport.DecodeGetAllRolesRequest,
		transport.EncodeGetAllRolesResponse,
		Opts[4],
	))

	// swagger:route PUT /api/user/{Username}/revoke admin RevokeUserRequest
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

	// swagger:route PUT /api/user/{Username}/password admin ResetPasswordRequest
	// Reset user password- only for admin
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
	//  200: ResetUserPasswordResponse
	r.Methods(http.MethodPut).Path(constants.RESET_PASSWORD).Handler(httptransport.NewServer(
		endpoints.ResetPasswordEndpoint,
		transport.DecodePasswordResetRequest,
		transport.EncodePasswordResetRequest,
		Opts[0],
	))

	// swagger:route PUT /api/user/password/reset user ResetUserPasswordRequest
	// Reset user password
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
	//  200: LoginResponse
	r.Methods(http.MethodPut).Path(constants.RESET_PASSWORD_USER).Handler(httptransport.NewServer(
		endpoints.ResetUserPasswordEndpoint,
		transport.DecodeUserPasswordResetRequest,
		transport.EncodePasswordResetRequest,
		Opts[0],
	))

	// swagger:route POST /api/user/login auth LoginRequest
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

	// swagger:route POST /api/token/refresh auth RefreshTokenRequest
	// Refresh token with existing token
	//
	// requests:
	// responses:
	//  500: InternalServerErrorResponse
	//  400: InvalidTokenResponse
	//  200: LoginResponse
	r.Methods(http.MethodPost).Path(constants.REFRESH_TOKEN).Handler(httptransport.NewServer(
		endpoints.RefreshTokenEndpoint,
		transport.DecodeRefreshTokenRequest,
		transport.EncodeRefreshTokenRequest,
		Opts[3],
	))

	// swagger:route POST /api/{Category}/question questions QuestionRequest
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

	// swagger:route POST /api/{Category}/questions questions QuestionsRequest
	// Fetch all categories
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
	//  400: BadRequestErrorResponse
	//  200: SuccessResponse
	r.Methods(http.MethodPost).Path(constants.ALL_QUESTIONS).Handler(httptransport.NewServer(
		endpoints.AddMultipleQuestionEndpoint,
		transport.DecodeAddMutlipleQuestionsRequest,
		transport.EncodeAddMultipleQuestionsResponse,
		Opts[1],
	))

	// swagger:route GET /api/{Category}/questions questions GetQuestionRequest
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

	// swagger:route GET /api/question/types questions listQuestionTypes
	// Get all question types
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
	//  401: UnAuthorizedAccessResponse
	//  200: QuestionTypeResponse
	r.Methods(http.MethodGet).Path(constants.GET_QUESTION_TYPE).Handler(httptransport.NewServer(
		endpoints.GetAllQuestionTypesEndpoint,
		transport.DecodeGetAllQuestionTypesRequest,
		transport.EncodeGetAllQuestionTypesResponse,
		Opts[5],
	))

	// swagger:route POST /api/category category CategoryRequest
	// Add new category
	//
	// responses:
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  409: ConflictErrorResponse
	//  201: DataSavedSuccessResponse
	r.Methods(http.MethodPost).Path(constants.CATEGORY).Handler(httptransport.NewServer(
		endpoints.AddCategoryEndpoint,
		transport.DecodeAddCategoryRequest,
		transport.EncodeAddCategoryResponse,
		Opts[2],
	))

	// swagger:route PUT /api/category/{id} category UpdateCategoryRequestId
	// Update existing category name
	//
	// responses:
	//  500: InternalServerErrorResponse
	//  400: BadRequestErrorResponse
	//  409: UpdateCategoryConflictResponse
	//  404: UpdateCategoryNotFoundResponse
	//  200: UpdateCategorySuccessResponse
	r.Methods(http.MethodPut).Path(constants.CATEGORY_ID).Handler(httptransport.NewServer(
		endpoints.UpdateCategoryEndpoint,
		transport.DecodeUpdateCategoryRequest,
		transport.EncodeUpdateCategoryResponse,
		Opts[2],
	))

	// swagger:route DELETE /api/category/{id} category DeleteCategoryRequestId
	// Update existing category name
	//
	// responses:
	//  500: InternalServerErrorResponse
	//  400: DeleteCategoryBadRequestResponse
	//  404: UpdateCategoryNotFoundResponse
	//  200: DeleteCategorySuccessResponse
	r.Methods(http.MethodDelete).Path(constants.CATEGORY_ID).Handler(httptransport.NewServer(
		endpoints.DeleteCategoryEndpoint,
		transport.DecodeDeleteCategoryRequest,
		transport.EncodeDeleteCategoryResponse,
		Opts[2],
	))

	// swagger:route GET /api/categories category GetAllCategoriesRequest
	// Fetch all categories
	//
	// responses:
	//  200: CategoriesResponse
	//  400: BadRequestErrorResponse
	//  401: UnAuthorizedAccessResponse
	//  500: InternalServerErrorResponse
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,authorization,Access-Control-Allow-Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

		err := Authorizer(w, r)
		if err != (globalErr.GlobalError{}) {
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func Authorizer(w http.ResponseWriter, r *http.Request) globalErr.GlobalError {

	var claims *jwtP.Claims
	var tokenErr globalErr.GlobalError
	if strings.Contains(r.RequestURI, "/login") || strings.Contains(r.RequestURI, "/health") || strings.Contains(r.RequestURI, "/docs") {
		claims = &jwtP.Claims{
			Role: "anonymous",
		}
	} else {
		// verify token
		claims, tokenErr = jwtP.VerifyToken(r)
		if tokenErr != (globalErr.GlobalError{}) {
			return tokenErr
		}

	}

	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcerSafe("./configuration/conf/auth/auth_model.conf", "./configuration/conf/auth/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// casbin rule enforcing
	res, err := authEnforcer.EnforceSafe(claims.Role, r.URL.Path, r.Method)
	if err != nil {
		internalServerError := globalErr.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusInternalServerError,
			Message:   "Something went wrong",
		}
		return internalServerError
	}
	if res {
		// next.ServeHTTP(w, r)
	} else {
		gerr := globalErr.GlobalError{
			TimeStamp: time.Now().UTC().String()[0:19],
			Status:    http.StatusForbidden,
			Message:   "User does not have access",
		}
		return gerr
	}
	return globalErr.GlobalError{}
}

//CreateNewServer is a function to return server
func CreateNewServer(c config.ConfigurationDetails, handler http.Handler) http.Server {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{c.Cors.Origin},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		// ExposedHeaders: []string{"Content-Type,authorization,Access-Control-Allow-Origin"},
		AllowCredentials: true,
		Debug:            c.Cors.Debug,
	})

	newHandler := cors.Handler(handler)
	return http.Server{
		Addr:    ":" + c.ServerPort,
		Handler: newHandler,
	}
}
