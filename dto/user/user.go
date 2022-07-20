package user

type UserDTO struct {
	Username    string
	OldPassword string
	NewPassword string
}

type UserPaginationDTO struct {
	Search  string
	Start   *int64
	Length  *int64
	SortBy  string
	OrderBy int
}
