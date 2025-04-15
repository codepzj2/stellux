package web

type PostRequest struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	IsPublished bool     `json:"is_published"`
	IsTop       bool     `json:"is_top"`
	Thumbnail   string   `json:"thumbnail"`
}
