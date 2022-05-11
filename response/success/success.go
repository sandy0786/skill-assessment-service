package success

type SuccessResponse struct {
	TimeStamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}
