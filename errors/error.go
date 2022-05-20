package errors

// swagger:model
// GlobalError is an exported type
type GlobalError struct {
	// example: 2022-05-20 16:59:05
	TimeStamp string `json:"timestamp"`
	// example: 500
	Status int `json:"status"`
	// example: Something went wrong
	Message string `json:"message"`
}

func (g GlobalError) Error() string {
	return g.Message
}
