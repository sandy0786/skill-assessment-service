package user

type UserResponse struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Role     string `validate:"required"`
}
