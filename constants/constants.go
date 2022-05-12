package constants

// api endpoints
const (
	STATUS_PATH    = "/api/health"              // GET
	USER           = "/api/user"                // POST
	USER_ID        = "/api/user/{user_name}"    // GET
	ALL_USERS      = "/api/users"               // GET
	QUESTION       = "/api/question/{category}" // POST {category}
	ALL_QUESTIONS  = "/api/questions"           // GET|POST {category}
	CATEGORY       = "/api/category"            // POST
	ALL_CATEGORIES = "/api/categories"          // GET
	DOC_PATH       = "/api/swagger/docs"        // GET
)
