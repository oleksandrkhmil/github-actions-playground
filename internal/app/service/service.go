package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/oleksandrkhmil/github-actions-playground/internal/config"
	"github.com/oleksandrkhmil/github-actions-playground/internal/domain/blog"
	"github.com/oleksandrkhmil/github-actions-playground/internal/repository/inmemory"
	"github.com/oleksandrkhmil/github-actions-playground/internal/server"
)

func Run(ctx context.Context, c config.Config) (func() error, error) {
	blogRepository := inmemory.NewBlogRepository(time.Now)

	blogService := blog.NewService(blogRepository)

	blogHandler := server.NewBlogHandler(blogService)

	httpServer := server.NewServer(c.ServerPort, blogHandler)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Unexpected server error", "error_message", err.Error())
		}
	}()

	return func() error {
		if err := httpServer.Close(); err != nil {
			return fmt.Errorf("close http server: %w", err)
		}

		return nil
	}, nil
}
