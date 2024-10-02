package inmemory

import (
	"context"
	"sync"
	"time"

	"github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"
)

type BlogRepository struct {
	mutex sync.RWMutex
	list  []blog.Post
	now   func() time.Time
}

func NewBlogRepository(now func() time.Time) *BlogRepository {
	return &BlogRepository{
		mutex: sync.RWMutex{},
		list:  make([]blog.Post, 0),
		now:   now,
	}
}

func (r *BlogRepository) Create(_ context.Context, p blog.Post) (blog.Post, error) {
	r.mutex.Lock()
	p.ID = int64(len(r.list)) + 1
	p.CreatedAt = r.now()
	r.list = append(r.list, p)
	r.mutex.Unlock()
	return p, nil
}

func (r *BlogRepository) GetAll(context.Context) ([]blog.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.list, nil
}
