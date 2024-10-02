package blog_test

import (
	"context"
	"testing"
	"time"

	"github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type serviceMocks struct {
	repository *Mockrepository
}

func newService(t *testing.T) (blog.Service, serviceMocks) {
	ctrl := gomock.NewController(t)
	repository := NewMockrepository(ctrl)
	service := blog.NewService(repository)
	return service, serviceMocks{repository: repository}
}

func TestService_Create(t *testing.T) {
	t.Run("It should return error if received invalid request", func(t *testing.T) {
		service, _ := newService(t)

		post := blog.Post{
			Content: "Missing title",
		}

		_, err := service.Create(context.Background(), post)
		assert.Error(t, err)
	})

	t.Run("It should create post", func(t *testing.T) {
		service, mocks := newService(t)

		post := blog.Post{
			Title:     "Any title",
			Tags:      []blog.Tag{{Title: "Personal"}},
			Content:   "Any content",
			CreatedAt: time.Now(),
		}

		expectedResult := post
		expectedResult.ID = 1

		mocks.repository.
			EXPECT().
			Create(gomock.Any(), post).
			DoAndReturn(func(ctx context.Context, p blog.Post) (blog.Post, error) {
				p.ID = 1
				return p, nil
			})

		result, err := service.Create(context.Background(), post)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})
}

func TestService_GetAll(t *testing.T) {
	service, mocks := newService(t)

	expectedList := []blog.Post{
		{
			Title:     "Any title 1",
			Tags:      []blog.Tag{{Title: "Personal"}},
			Content:   "Any content 1",
			CreatedAt: time.Now(),
		},
		{
			Title:     "Any title 2",
			Tags:      []blog.Tag{{Title: "Personal"}},
			Content:   "Any content 2",
			CreatedAt: time.Now(),
		},
	}

	mocks.repository.EXPECT().GetAll(gomock.Any()).Return(expectedList, nil)

	list, err := service.GetAll(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, expectedList, list)
}
