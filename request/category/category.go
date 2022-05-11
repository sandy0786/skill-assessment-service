package category

type CategoryRequest struct {
	CategoryName   string `json:"categoryName" validate:"required"`
	CollectionName string `json:"collectionName" validate:"required"`
	Author         string `json:"author" validate:"required"`
}
