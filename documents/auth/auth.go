package auth

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email,unique"`
	Role     string `bson:"role"`
	Active   bool   `bson:"active"`
}
