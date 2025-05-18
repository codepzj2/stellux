package web

type DocumentRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	ParentID   string `json:"parent_id"`
	DocumentID string `json:"document_id"`
}