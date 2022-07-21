package constants

const (
	PAGE       = "page"
	PAGE_SIZE  = "pageSize"
	SEARCH     = "search"
	SORT_BY    = "sortBy"
	ORDER_BY   = "orderBy"
	ASCENDING  = "asc"
	DESCENDING = "desc"
)

var (
	// get all users
	ALLOWED_QUERY_PARAMS      = []string{PAGE, PAGE_SIZE, SEARCH, SORT_BY, ORDER_BY}
	REQUIRED_QUERY_PARAMS     = []string{PAGE, PAGE_SIZE, SEARCH}
	ALLOWED_ORDER_BY_LITERALS = []string{ASCENDING, DESCENDING}
	ALLOWED_SORT_BY_LITERALS  = []string{"username", "email", "createdAt", "updatedAt"}
)

// error messages
var ValidationErrors = map[string]string{
	"ErrInvalidQueryParam":   "Invalid Query param : ",
	"ErrMandatoryQueryParam": "Mandatory Query param missing : ",
	"ErrQueryParamOccurence": "Multiple occurence of query param is not allowed : ",
	"ErrInvalidData":         "Invalid data provided. Provide number type data which is greater than '0' : ",
	"ErrInvalidSortBy":       "Invalid order by literal provided. Supported only (username|email|createdAt|updatedAt) : ",
	"ErrInvalidOrderBy":      "Invalid order by literal provided. Supported only (asc|desc) : ",
	"ErrMissingSortBy":       "'orderBy' require 'sortBy' query param : ",
	"ErrMissingOrderBy":      "'sortBy' require 'orderBy' query param : ",
}
