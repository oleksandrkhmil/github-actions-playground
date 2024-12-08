package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"
	"github.com/oleksandrkhmil/github-actions-playground/internal/server"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type blogHandlerMocks struct {
	blogService *MockblogService
}

func newBlogHandler(t *testing.T) (server.BlogHandler, blogHandlerMocks) {
	ctrl := gomock.NewController(t)
	blogService := NewMockblogService(ctrl)
	blogHandler := server.NewBlogHandler(blogService)
	return blogHandler, blogHandlerMocks{blogService: blogService}
}

func TestBlogHandler_Create(t *testing.T) {
	now, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	require.NoError(t, err)

	blogHandler, mocks := newBlogHandler(t)

	requestModel := server.Post{
		Title:   "Any title",
		Tags:    []string{"Personal"},
		Content: "Any content",
	}

	requestData, err := json.Marshal(requestModel)
	require.NoError(t, err)

	httpRequest := httptest.NewRequest(http.MethodPost, "https://example.com/api/v1/posts", bytes.NewReader(requestData))

	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/posts", blogHandler.Create)

	expectedServiceModel := blog.Post{
		Title:   "Any title",
		Tags:    []blog.Tag{{Title: "Personal"}},
		Content: "Any content",
	}

	expectedResponse := server.Post{
		ID:        1,
		Title:     "Any title",
		Tags:      []string{"Personal"},
		Content:   "Any content",
		CreatedAt: "2006-01-02T15:04:05Z",
	}

	mocks.blogService.
		EXPECT().
		Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, p blog.Post) (blog.Post, error) {
			assert.Equal(t, expectedServiceModel, p)
			p.ID = 1
			p.CreatedAt = now
			return p, nil
		})

	mux.ServeHTTP(recorder, httpRequest)

	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

	responseData, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	var responseModel server.Post
	require.NoError(t, json.Unmarshal(responseData, &responseModel))

	assert.Equal(t, expectedResponse, responseModel)
}

func TestBlogHandler_GetAll(t *testing.T) {
	now, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	require.NoError(t, err)

	blogHandler, mocks := newBlogHandler(t)

	httpRequest := httptest.NewRequest(http.MethodGet, "https://example.com/api/v1/posts", http.NoBody)

	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/posts", blogHandler.GetAll)

	serviceResult := []blog.Post{
		{
			ID:        1,
			Title:     "Post 1",
			Tags:      []blog.Tag{{Title: "Personal"}},
			Content:   "Post content 1",
			CreatedAt: now,
		},
		{
			ID:        2,
			Title:     "Post 2",
			Tags:      []blog.Tag{{Title: "Personal"}},
			Content:   "Post content 2",
			CreatedAt: now,
		},
	}

	expectedResponse := []server.Post{
		{
			ID:        1,
			Title:     "Post 1",
			Tags:      []string{"Personal"},
			Content:   "Post content 1",
			CreatedAt: "2006-01-02T15:04:05Z",
		},
		{
			ID:        2,
			Title:     "Post 2",
			Tags:      []string{"Personal"},
			Content:   "Post content 2",
			CreatedAt: "2006-01-02T15:04:05Z",
		},
	}

	mocks.blogService.
		EXPECT().
		GetAll(gomock.Any()).
		Return(serviceResult, nil)

	mux.ServeHTTP(recorder, httpRequest)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	responseData, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	var responseModel []server.Post
	require.NoError(t, json.Unmarshal(responseData, &responseModel))

	assert.Equal(t, expectedResponse, responseModel)
}

func TestBlogHandler_GetByID(t *testing.T) {
	now, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	require.NoError(t, err)

	blogHandler, mocks := newBlogHandler(t)

	httpRequest := httptest.NewRequest(http.MethodGet, "https://example.com/api/v1/posts/1", http.NoBody)

	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/posts/{id}", blogHandler.GetByID)

	serviceResult := blog.Post{
		ID:        1,
		Title:     "Post 1",
		Tags:      []blog.Tag{{Title: "Personal"}},
		Content:   "Post content 1",
		CreatedAt: now,
	}

	expectedResponse := server.Post{
		ID:        1,
		Title:     "Post 1",
		Tags:      []string{"Personal"},
		Content:   "Post content 1",
		CreatedAt: "2006-01-02T15:04:05Z",
	}

	mocks.blogService.
		EXPECT().
		GetByID(gomock.Any(), int64(1)).
		Return(serviceResult, nil)

	mux.ServeHTTP(recorder, httpRequest)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	responseData, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	t.Log(string(responseData))

	var responseModel server.Post
	require.NoError(t, json.Unmarshal(responseData, &responseModel))

	assert.Equal(t, expectedResponse, responseModel)
}
