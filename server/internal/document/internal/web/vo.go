package web

import "github.com/codepzj/stellux/server/internal/pkg/apiwrap"

type DocumentVO struct {
	ID         string         `json:"id"`
	CreatedAt  string         `json:"created_at"`
	UpdatedAt  string         `json:"updated_at"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	ParentID   apiwrap.BsonID `json:"parent_id,omitzero"`
	DocumentID apiwrap.BsonID `json:"document_id,omitzero"`
}
