package common

import (
	"strconv"
)

// function types
type ValidatorIsNumberType func(string) (bool, string)

func IsQueryParamNumber(queryParam string) (bool, string) {
	_, err := strconv.Atoi(queryParam)
	if err != nil {
		return false, "ErrInvalidData"
	}
	return true, ""
}
