package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"
)

//go:generate mockgen -source=$GOFILE -destination=blog_handler_mock_test.go -package=${GOPACKAGE}_test -typed=true

type blogService interface {
	Create(context.Context, blog.Post) (blog.Post, error)
	GetAll(context.Context) ([]blog.Post, error)
	GetByID(context.Context, int64) (blog.Post, error)
}

type BlogHandler struct {
	blogService blogService
}

func NewBlogHandler(blogService blogService) BlogHandler {
	return BlogHandler{blogService: blogService}
}

func (h BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	var apiRequest Post
	if err := json.NewDecoder(r.Body).Decode(&apiRequest); err != nil {
		respond(r.Context(), w, http.StatusBadRequest, newErrorResponse(CodeInvalidRequest, err.Error()))
		return
	}

	domainRequest, err := apiRequest.ToDomain()
	if err != nil {
		respond(r.Context(), w, http.StatusBadRequest, newErrorResponse(CodeInvalidRequest, err.Error()))
		return
	}

	response, err := h.blogService.Create(r.Context(), domainRequest)
	if err != nil {
		respond(r.Context(), w, http.StatusBadRequest, newErrorResponse(CodeInvalidRequest, err.Error()))
	}

	respond(r.Context(), w, http.StatusCreated, NewPost(response))
}

func (h BlogHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.blogService.GetAll(r.Context())
	if err != nil {
		respond(r.Context(), w, http.StatusBadRequest, newErrorResponse(CodeInvalidRequest, err.Error()))
		return
	}

	respond(r.Context(), w, http.StatusOK, NewPosts(posts))
}

func (h BlogHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	parsedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		respond(r.Context(), w, http.StatusBadRequest, newErrorResponse(CodeInvalidRequest, err.Error()))
		return
	}

	post, err := h.blogService.GetByID(r.Context(), parsedID)
	if errors.Is(err, blog.ErrNotFound) {
		respond(r.Context(), w, http.StatusNotFound, newErrorResponse(CodeInvalidRequest, err.Error()))
		return
	}
	if err != nil {
		respond(r.Context(), w, http.StatusBadRequest, newErrorResponse(CodeInvalidRequest, err.Error()))
		return
	}

	respond(r.Context(), w, http.StatusOK, NewPost(post))
}
