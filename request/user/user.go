package user

// User Request
// swagger:model
type UserRequest struct {
	// the username for this user
	// required: true
	Username string `json:"username" validate:"required,min=5"`
	// the password for this user
	// required: true
	// min length: 8
	Password string `json:"password" validate:"required,min=8"`
	// the email for this user
	// required: true
	// example: user@provider.com
	Email string `json:"email" validate:"required,email"`
	// the roleId for this user
	// required: true
	// example: 62d64f2142dac7953ac4ff32
	Role string `json:"roleId" validate:"required"`
}

// User Request
// swagger:parameters UserRequest
type UserRequestSwagger struct {
	// in:body
	Body UserRequest
}

// swagger:parameters DeleteUserRequest
type DeleteUserRequestSwagger struct {
	// in: path
	// required: true
	// example: username
	Username string
}

// swagger:parameters RevokeUserRequest
type RevokeUserRequestSwagger struct {
	// in: path
	// required: true
	// example: username
	Username string
}

// swagger:model
type PasswordReset struct {
	// the old password for this user
	// required: true
	OldPassword string `json:"oldPassword" validate:"required"`
	// the new password for this user
	// required: true
	NewPassword string `json:"newPassword" validate:"required"`
}

// swagger:parameters ResetPasswordRequest
type ResetPasswordRequestSwagger struct {
	// in: path
	// required: true
	// example: admin
	Username string
	// in: body
	Body PasswordReset
}

// swagger:parameters ResetUserPasswordRequest
type ResetUserPasswordRequestSwagger struct {
	// in: body
	Body PasswordReset
}
