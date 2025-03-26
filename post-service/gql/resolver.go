package gql

import (
	"context"
	"post-service/db"
	"post-service/models"
	"time"
)

type Resolver struct{}

func (r *Resolver) CreatePost(ctx context.Context, title string, content string, authorID string) (*models.Post, error) {
	post := &models.Post{
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
	}
	db.DB.Create(post)
	return post, nil
}

func (r *Resolver) CreateComment(ctx context.Context, postID string, content string, authorID string) (*models.Comment, error) {
	comment := &models.Comment{
		PostID:    postID,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
	}
	db.DB.Create(comment)
	return comment, nil
}
