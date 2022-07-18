package role

// import "go.mongodb.org/mongo-driver/bson/primitive"

// User Response
// swagger:model
type Role struct {
	// example: admin
	Role string `json:"role"`
}

// List of Roles
// swagger:response RolesResponse
type UsersResponse struct {
	// in: body
	Body []Role
}
