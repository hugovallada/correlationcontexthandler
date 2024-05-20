package config

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hugovallada/slog-multikey-context-handler/internal/keys"
)

type CorrelationContextHandler struct {
	slog.Handler
}

func (h CorrelationContextHandler) Handle(ctx context.Context, record slog.Record) error {
	if correlationId, ok := ctx.Value(keys.CORRELATION_ID).(string); ok {
		record.Add(fmt.Sprintf("%v", keys.CORRELATION_ID), slog.StringValue(correlationId))
	}
	return h.Handler.Handle(ctx, record)
}
