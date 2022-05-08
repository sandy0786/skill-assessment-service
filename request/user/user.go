package user

type UserRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required"`
	Role     string `validate:"required"`
}
