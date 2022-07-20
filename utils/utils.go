/*
Package utils has the functions for date convertion and array functions
*/
package utils

// FindItemsInList exported function to search an item in array. returns true if found else false
func FindItemsInList(items []string, searchItem string) bool {
	for _, str := range items {
		if str == searchItem {
			return true
		}
	}
	return false
}
