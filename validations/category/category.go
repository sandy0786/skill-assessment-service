package category

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	constants "github.com/sandy0786/skill-assessment-service/constants"
	globalErr "github.com/sandy0786/skill-assessment-service/errors"
	utils "github.com/sandy0786/skill-assessment-service/utils"
	commonValidations "github.com/sandy0786/skill-assessment-service/validations/common"
)

// function types
type validatorQueryparams func(map[string][]string, []string) (bool, string, string)
type validatorError func(string, string) globalErr.GlobalError

type CategoryValidator interface {
	ValidateGetAllUsersRequest(*http.Request) (bool, error)
}

type categoryValidator struct {
	CategoryValidator
	ValidationMandatoryQueryparams validatorQueryparams
	ValidationQueryparamsOccurence validatorQueryparams
	ValidationQueryParamsCheck     validatorQueryparams
	ValidationError                validatorError
	ValidatorNumberType            commonValidations.ValidatorIsNumberType
}

func NewCategoryValidator() *categoryValidator {
	return &categoryValidator{
		ValidationMandatoryQueryparams: validateMandatoryQueryparams,
		ValidationQueryparamsOccurence: validateQueryparamsOccurences,
		ValidationQueryParamsCheck:     queryParamsCheck,
		ValidationError:                validationError,
		ValidatorNumberType:            commonValidations.IsQueryParamNumber,
	}
}

func (u *categoryValidator) ValidateGetAllUsersRequest(r *http.Request) (bool, error) {
	log.Println("Validator : Inside ValidateGetAllUsersRequest")
	// verify query params
	query := r.URL.Query()

	// check for query params which are not allowed
	flag, errMsg, qp := u.ValidationQueryParamsCheck(query, constants.ALLOWED_QUERY_PARAMS_ALL_CATEGORIES)
	if !flag {
		return false, u.ValidationError(errMsg, qp)
	}

	// validate mandatory query params available or not
	flag, errMsg, qp = u.ValidationMandatoryQueryparams(query, constants.REQUIRED_QUERY_PARAMS)
	if !flag {
		return false, u.ValidationError(errMsg, qp)
	}

	// validate allowed query params available or not
	flag, errMsg, qp = u.ValidationQueryparamsOccurence(query, constants.ALLOWED_QUERY_PARAMS)
	if !flag {
		return false, u.ValidationError(errMsg, qp)
	}

	// validate page number
	val, flag, errMsg := u.ValidatorNumberType(query[constants.PAGE][0])
	if !flag {
		return false, u.ValidationError(errMsg, constants.PAGE)
	}

	// Check for zero value
	if val == 0 {
		return false, u.ValidationError("ErrInvalidData", constants.PAGE)
	}

	// validate page size
	val, flag, errMsg = u.ValidatorNumberType(query[constants.PAGE_SIZE][0])
	if !flag {
		return false, u.ValidationError(errMsg, constants.PAGE_SIZE)
	}

	// Check for zero value
	if val == 0 {
		return false, u.ValidationError("ErrInvalidData", constants.PAGE_SIZE)
	}

	// validate orderBy
	if len(query[constants.ORDER_BY]) > 0 {
		orderBy := query[constants.ORDER_BY][0]
		flag = utils.FindItemsInList(constants.ALLOWED_ORDER_BY_LITERALS, strings.ToLower(orderBy))
		if !flag {
			return false, u.ValidationError("ErrInvalidOrderBy", constants.ORDER_BY)
		}
	}

	return true, nil
}

// queryParamsCheck is to check for invalid query params
func queryParamsCheck(query map[string][]string, allowedQueryParams []string) (bool, string, string) {
	log.Println("validations: queryParamsCheck: inside queryParamsCheck()")
	for key, _ := range query {
		if !utils.FindItemsInList(allowedQueryParams, key) {
			return false, "ErrInvalidQueryParam", key
		}
	}
	return true, "", ""
}

// validateMandatoryQueryparams validates mandatory query params
func validateMandatoryQueryparams(query map[string][]string, mandatoryQueryParams []string) (bool, string, string) {
	log.Println("validations: validateMandatoryQueryparams: inside validateMandatoryQueryparams()")

	for _, val := range mandatoryQueryParams {
		if len(query[val]) == 0 {
			return false, "ErrMandatoryQueryParam", val
		}
	}
	return true, "", ""
}

// validateQueryparamsOccurences validates query params occurences
func validateQueryparamsOccurences(query map[string][]string, allowedQueryParams []string) (bool, string, string) {
	log.Println("validations: validateQueryparamsOccurences: inside validateQueryparamsOccurences()")

	for _, val := range allowedQueryParams {
		if len(query[val]) > 1 {
			return false, "ErrQueryParamOccurence", val
		}
	}
	return true, "", ""
}

// validationError creates global error and returns
func validationError(errorKey string, val string) globalErr.GlobalError {
	return globalErr.GlobalError{
		Message:   fmt.Sprintf("%v", constants.ValidationErrors[errorKey]) + "'" + val + "'",
		TimeStamp: time.Now().UTC().String(),
		Status:    http.StatusBadRequest,
	}
}
