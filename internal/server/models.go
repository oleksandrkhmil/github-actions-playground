package server

import (
	"fmt"
	"time"

	"github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"
)

type Code string

const (
	CodeInvalidRequest Code = "invalid_request"
)

type ErrResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(code Code, message string) ErrResponse {
	return ErrResponse{Code: code, Message: message}
}

type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	Content   string   `json:"content"`
	CreatedAt string   `json:"created_at,omitempty"`
}

func NewPost(d blog.Post) Post {
	tags := make([]string, len(d.Tags))
	for i, v := range d.Tags {
		tags[i] = v.Title
	}

	return Post{
		ID:        d.ID,
		Title:     d.Title,
		Tags:      tags,
		Content:   d.Content,
		CreatedAt: d.CreatedAt.Format(time.RFC3339),
	}
}

func NewPosts(sl []blog.Post) []Post {
	result := make([]Post, len(sl))
	for i, v := range sl {
		result[i] = NewPost(v)
	}

	return result
}

func (p Post) ToDomain() (_ blog.Post, err error) {
	tags := make([]blog.Tag, len(p.Tags))
	for i, v := range p.Tags {
		tags[i] = blog.Tag{Title: v}
	}

	var createdAt time.Time
	if p.CreatedAt != "" {
		createdAt, err = time.Parse(time.RFC3339, p.CreatedAt)
		if err != nil {
			return blog.Post{}, fmt.Errorf("parse created at: %w", err)
		}
	}

	return blog.Post{
		ID:        p.ID,
		Title:     p.Title,
		Tags:      tags,
		Content:   p.Content,
		CreatedAt: createdAt,
	}, nil
}
