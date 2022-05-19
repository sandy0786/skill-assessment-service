package success

// Success response
// swagger:response SuccessResponse
type SuccessResponse struct {
	TimeStamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}

// swagger:response HealthResponse
type Health struct {
	Status string `json:"status"`
}
