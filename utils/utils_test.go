// +build utils_test or all

package utils

import (
	"fmt"
	"log"
	"testing"
	"time"

	constants "metrics/constants"
	model "metrics/model"

	"github.com/stretchr/testify/suite"
)

// This is our suite
type UtilsSuite struct {
	suite.Suite
}

// Test FindItemsInList
func (suite *UtilsSuite) TestFindItemsInList() {
	log.Println("inside TestFindItemsInList")
	inp := []string{"val1", "val2"}
	flag := FindItemsInList(inp, "val1")
	suite.True(flag)

	inp = []string{"val1", "val2"}
	flag = FindItemsInList(inp, "val3")
	suite.False(flag)
}

// This gets run automatically by `go test`
func TestUtilsSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(UtilsSuite))
}
