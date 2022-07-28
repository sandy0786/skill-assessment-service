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
	// example: testuser
	Username string
	// in: body
	Body struct {
		// the new password for this user
		// required: true
		Password string `json:"Password" validate:"required,min=8"`
	}
}

// swagger:parameters ResetUserPasswordRequest
type ResetUserPasswordRequestSwagger struct {
	// in: body
	Body PasswordReset
}

// swagger:parameters GetAllUsersRequest
type AllUsersRequestSwagger struct {
	// Provide page number
	// in: query
	// required: true
	// example: 1
	Page int `json:"page"`
	// Provide page size
	// in: query
	// required: true
	// example: 10
	PageSize int `json:"pageSize"`
	// Provide search literal
	// in: query
	// required: true
	// example: admin
	Search string `json:"search"`
	// Provide sort by field (username|email|createdAt|updatedAt)
	// in: query
	// required: false
	// example: username
	SortBy string `json:"sortBy"`
	// provide order by asc|desc
	// in: query
	// required: false
	// example: asc
	OrderBy string `json:"orderBy"`
}
