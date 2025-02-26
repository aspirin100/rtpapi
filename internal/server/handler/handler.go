package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Generator interface {
	GenerateMultiplier() float32
}

type Handler struct {
	generator Generator
	server    *http.Server
}

func New(generator Generator) *Handler {
	handler := &Handler{
		generator: generator,
	}

	router := gin.Default()

	router.GET("/get", handler.Generate)

	server := http.Server{
		Addr:    "localhost:64333",
		Handler: router,
	}

	handler.server = &server

	return handler
}

func (h *Handler) Start() error {
	err := h.server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start http server: %w", err)
	}

	return nil
}

func (h *Handler) Shutdown(ctx context.Context) error {
	err := h.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to stop http server: %w", err)
	}

	return nil
}

func (h *Handler) Generate(ctx *gin.Context) {
	mp := h.generator.GenerateMultiplier()

	resp := struct {
		Result *float32 `json:"result"`
	}{
		Result: &mp,
	}

	ctx.JSON(http.StatusOK, resp)
}
