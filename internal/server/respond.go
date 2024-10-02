package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func respond(ctx context.Context, w http.ResponseWriter, status int, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to marshal response", "error_message", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to write response", "error_message", err.Error())
		return
	}
}
