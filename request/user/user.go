package user

// User Request
// swagger:model
type UserRequest struct {
	// the username for this user
	// required: true
	Username string `json:"username" validate:"required"`
	// the password for this user
	// required: true
	// min length: 8
	Password string `json:"password" validate:"required"`
	// the email for this user
	// required: true
	// example: user@provider.com
	Email string `json:"email" validate:"required"`
	// the role for this user
	// required: true
	// example: manager
	Role string `json:"role" validate:"required"`
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

// swagger:parameters ResetUserPasswordRequest
type ResetPasswordRequestSwagger struct {
	// in: path
	// required: true
	// example: admin
	Username string
	// in: body
	Body PasswordReset
}
