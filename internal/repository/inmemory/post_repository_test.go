package inmemory_test

import (
	"context"
	"testing"
	"time"

	"github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"
	"github.com/oleksandrkhmil/github-actions-playground/internal/repository/inmemory"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlogRepository(t *testing.T) {
	now, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	require.NoError(t, err)

	blogRepository := inmemory.NewBlogRepository(func() time.Time {
		return now
	})

	post := blog.Post{
		Title:   "Any title",
		Tags:    []blog.Tag{{Title: "Personal"}},
		Content: "Any content",
	}

	expectedPost := blog.Post{
		ID:        1,
		Title:     "Any title",
		Tags:      []blog.Tag{{Title: "Personal"}},
		Content:   "Any content",
		CreatedAt: now,
	}

	result, err := blogRepository.Create(context.Background(), post)
	require.NoError(t, err)
	assert.Equal(t, expectedPost, result)

	resultList, err := blogRepository.GetAll(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []blog.Post{expectedPost}, resultList)

	resultGetByID, err := blogRepository.GetByID(context.Background(), expectedPost.ID)
	require.NoError(t, err)
	assert.Equal(t, expectedPost, resultGetByID)
}
