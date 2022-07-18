package constants

// api endpoints
const (
	STATUS_PATH = "/api/health" // GET

	// users
	USER                = "/api/user"                     // POST
	USER_NAME           = "/api/user/{username}"          // GET|DELETE|PUT
	ALL_USERS           = "/api/users"                    // GET
	REVOKE_USER         = "/api/user/{username}/revoke"   // PUT
	RESET_PASSWORD      = "/api/user/{username}/password" // PUT
	RESET_PASSWORD_USER = "/api/user/password/reset"      // PUT

	//role
	All_ROLES = "/api/user/roles" // GET

	// update roles
	QUESTION       = "/api/question/{category}"  // POST
	ALL_QUESTIONS  = "/api/questions/{category}" // GET|POST
	CATEGORY       = "/api/category"             // POST
	ALL_CATEGORIES = "/api/categories"           // GET
	DOC_PATH       = "/docs"                     // GET

	// auth
	LOGIN         = "/api/user/login"    // POST
	LOGOUT        = "/api/user/logout"   // POST
	REFRESH_TOKEN = "/api/token/refresh" // POST
)
