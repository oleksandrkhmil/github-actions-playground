package blog

import (
	"context"
	"fmt"
)

//go:generate mockgen -source=$GOFILE -destination=service_mock_test.go -package=${GOPACKAGE}_test -typed=true

type repository interface {
	Create(context.Context, Post) (Post, error)
	GetAll(context.Context) ([]Post, error)
	GetByID(context.Context, int64) (Post, error)
}

type Service struct {
	repository repository
}

func NewService(repository repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) Create(ctx context.Context, post Post) (Post, error) {
	if err := post.Validate(); err != nil {
		return Post{}, fmt.Errorf("validate: %w", err)
	}

	post, err := s.repository.Create(ctx, post)
	if err != nil {
		return Post{}, fmt.Errorf("create: %w", err)
	}

	return post, nil
}

func (s Service) GetAll(ctx context.Context) ([]Post, error) {
	posts, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all: %w", err)
	}

	return posts, nil
}

func (s Service) GetByID(ctx context.Context, id int64) (Post, error) {
	post, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return Post{}, fmt.Errorf("get by id: %w", err)
	}

	return post, nil
}
