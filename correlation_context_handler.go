// correlation_context_handler.go
package correlationcontexthandler

import (
	"context"
	"fmt"
	"log/slog"
)

type CorrelationContextHandler struct {
	slog.Handler
}

func (h CorrelationContextHandler) Handle(ctx context.Context, record slog.Record) error {
	if correlationId, ok := ctx.Value(CORRELATION_ID).(string); ok {
		record.Add(fmt.Sprintf("%v", CORRELATION_ID), slog.StringValue(correlationId))
	}
	return h.Handler.Handle(ctx, record)
}
