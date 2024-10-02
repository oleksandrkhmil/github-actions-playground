package blog

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Tag struct {
	ID        int64
	Title     string `validate:"required,max=100"`
	CreatedAt time.Time
}

func (t Tag) Validate() error {
	if err := validate.Struct(t); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

type Post struct {
	ID        int64
	Title     string `validate:"required,max=100"`
	Tags      []Tag
	Content   string `validate:"required,max=1000"`
	CreatedAt time.Time
}

func (p Post) Validate() error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}
