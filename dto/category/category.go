package category

type Pagination struct {
	Search  string
	Start   *int64
	Length  *int64
	OrderBy int
}

type UpdateCategory struct {
	Id       string
	Category string
}
