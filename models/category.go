package models

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Id            string `json:"id"`
	CategoryTitle string `json:"category_title"`
	ParentID      string `json:"parent_id"`
	CategoryImage string `json:"image"`
}

type Category struct {
	Id            string `json:"id"`
	CategoryTitle string `json:"category_title"`
	CategoryImage string `json:"image"`
	ParentID      string `json:"parent_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type UpdateCategory struct {
	Id            string `json:"id"`
	CategoryTitle string `json:"category_title"`
	ParentID      string `json:"parent_id"`
}

type GetListCategoryRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type GetListCategoryResponse struct {
	Count      int         `json:"count"`
	Categories []*Category `json:"categories"`
}
