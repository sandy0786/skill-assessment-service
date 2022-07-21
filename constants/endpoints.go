package constants

// api endpoints
const (

	// misceleneous
	STATUS_PATH = "/api/health" // GET
	DOC_PATH    = "/docs"       // GET

	// users
	USER                = "/api/user"                     // POST
	USER_NAME           = "/api/user/{username}"          // GET|DELETE|PUT
	ALL_USERS           = "/api/users"                    // GET
	REVOKE_USER         = "/api/user/{username}/revoke"   // PUT
	RESET_PASSWORD      = "/api/user/{username}/password" // PUT
	RESET_PASSWORD_USER = "/api/user/password/reset"      // PUT

	// role
	All_ROLES = "/api/user/roles" // GET

	// questions
	QUESTION          = "/api/{category}/question"  // POST
	ALL_QUESTIONS     = "/api/{category}/questions" // GET|POST|PUT
	GET_QUESTION_TYPE = "/api/question/types"       // GET

	// Categories
	CATEGORY       = "/api/category"      // POST
	CATEGORY_ID    = "/api/category/{id}" // DELETE|PUT
	ALL_CATEGORIES = "/api/categories"    // GET

	// auth
	LOGIN         = "/api/user/login"    // POST
	LOGOUT        = "/api/user/logout"   // POST
	REFRESH_TOKEN = "/api/token/refresh" // POST
)
