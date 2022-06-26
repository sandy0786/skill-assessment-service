package auth

// Login Response
// swagger:model
type LoginResponse struct {
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjU2MjQxNDE4fQ.Apq5YRkhS5UWr7bYgMKihH-GfcdDygKr777zdU5YWmI
	Token string `json:"token"`
}

// swagger:response LoginResponse
type UsersResponse struct {
	// in: body
	Body LoginResponse
}
