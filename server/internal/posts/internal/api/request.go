package api

import "server/internal/posts/internal/domain"

type PostsReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func toPosts(postsReq PostsReq) *domain.Posts {
	return &domain.Posts{
		Title:   postsReq.Title,
		Content: postsReq.Content,
		Author:  postsReq.Author,
	}
}
