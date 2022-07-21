package common

import (
	"strconv"
)

// function types
type ValidatorIsNumberType func(string) (int, bool, string)

func IsQueryParamNumber(queryParam string) (int, bool, string) {
	val, err := strconv.Atoi(queryParam)
	if err != nil {
		return val, false, "ErrInvalidData"
	}
	return val, true, ""
}
