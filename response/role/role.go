package role

// User Response
// swagger:model
type Role struct {
	// example: 62d6435f333f27963c162a05
	ID string `json:"id"`
	// example: admin
	Role string `json:"role"`
}

// List of Roles
// swagger:response RolesResponse
type UsersResponse struct {
	// in: body
	Body []Role
}
