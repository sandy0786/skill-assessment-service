package errors

//GlobalError is an exported type
type GlobalError struct {
	Result    []map[string]interface{} `json:"result"`
	TimeStamp string                   `json:"timestamp"`
	Status    int                      `json:"status"`
	Message   string                   `json:"message"`
}

func (g GlobalError) Error() string {
	return g.Message
}
