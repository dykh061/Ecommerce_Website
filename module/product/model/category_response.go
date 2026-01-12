package productmodel

// CategoryListItem for category tree response
type CategoryListItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}
