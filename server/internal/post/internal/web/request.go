package web

type PostRequest struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	IsPublish   bool     `json:"is_publish"`
	IsTop       bool     `json:"is_top"`
	Thumbnail   string   `json:"thumbnail" binding:"required"`
}
