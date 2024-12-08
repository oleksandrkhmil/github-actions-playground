package inmemory

import (
	"cmp"
	"context"
	"slices"
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
	defer r.mutex.Unlock()

	p.ID = int64(len(r.list)) + 1
	p.CreatedAt = r.now()

	r.list = append(r.list, p)

	return p, nil
}

func (r *BlogRepository) GetAll(context.Context) ([]blog.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.list, nil
}

func (r *BlogRepository) GetByID(_ context.Context, id int64) (blog.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	found, ok := slices.BinarySearchFunc(r.list, id, func(p blog.Post, id int64) int {
		return cmp.Compare(p.ID, id)
	})
	if !ok {
		return blog.Post{}, blog.ErrNotFound
	}

	return r.list[found], nil
}
