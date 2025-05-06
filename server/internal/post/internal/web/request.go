package web

type PostRequest struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	CategoryID  string   `json:"category_id"`
	TagsID      []string `json:"tags_id"`
	IsPublish   bool     `json:"is_publish"`
	IsTop       bool     `json:"is_top"`
	Thumbnail   string   `json:"thumbnail" binding:"required"`
}

type PostIDRequest struct {
	ID string `uri:"id" binding:"required"`
}
