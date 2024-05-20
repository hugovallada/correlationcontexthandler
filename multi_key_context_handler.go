// multi_key_context_handler.go
package correlationcontexthandler

import (
	"context"
	"fmt"
	"log/slog"
)

type MultiKeyContextHandler struct {
	slog.Handler
	keys []any
}

func New(keys []any, handler slog.Handler) MultiKeyContextHandler {
	return MultiKeyContextHandler{
		handler, keys,
	}
}

func (h MultiKeyContextHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, key := range h.keys {
		if value, ok := ctx.Value(key).(string); ok {
			record.Add(fmt.Sprintf("%v", key), slog.StringValue(value))
		}
	}
	return h.Handler.Handle(ctx, record)
}
