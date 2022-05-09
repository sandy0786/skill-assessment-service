package user

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// User success created
type UserSuccessResponse struct {
	TimeStamp string `json:"timestamp"`
	Status    int          `json:staus"`
	Message   string `json:"message"`
}
