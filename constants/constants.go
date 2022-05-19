package constants

// api endpoints
const (
	STATUS_PATH    = "/api/health"               // GET
	USER           = "/api/user"                 // POST
	USER_ID        = "/api/user/{user_name}"     // GET
	ALL_USERS      = "/api/users"                // GET
	QUESTION       = "/api/question/{category}"  // POST
	ALL_QUESTIONS  = "/api/questions/{category}" // GET|POST
	CATEGORY       = "/api/category"             // POST
	ALL_CATEGORIES = "/api/categories"           // GET
	DOC_PATH       = "/docs"                     // GET
)
