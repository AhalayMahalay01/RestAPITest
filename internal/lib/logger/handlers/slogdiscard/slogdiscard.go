package slogdiscard

import (
	"context"

	"golang.org/x/exp/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct{}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// Ігнорує запис журналу
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// Повертає обробник, не має атрибутів для збереження
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	// Повертає обробник, не має групи для збереження
}

func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// Завжди повертає false, запис журналу ігнорується
	return false
}
