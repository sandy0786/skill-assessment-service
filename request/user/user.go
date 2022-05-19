package user

// User Request
// swagger:parameters UserRequest
type UserRequest struct {
	// in:body
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Role     string `json:"role" validate:"required"`
}
