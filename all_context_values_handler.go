// all_context_values_handler.go
package correlationcontexthandler

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

type AllContextValuesHandler struct {
	slog.Handler
}

func NewAllContextValueHandler(slog slog.Handler) AllContextValuesHandler {
	return AllContextValuesHandler{
		slog,
	}
}

func (h AllContextValuesHandler) Handle(ctx context.Context, record slog.Record) error {
	contextValues := getContextValues(ctx)
	for key, value := range contextValues {
		record.Add(fmt.Sprintf("%v", key), slog.StringValue(value))
	}
	return h.Handler.Handle(ctx, record)
}

func getContextValues(ctx context.Context) map[string]string {
	ctxAsString := fmt.Sprintf("%s", ctx)
	stringAsList := strings.Split(ctxAsString, ".WithValue")
	listAsString := fmt.Sprintf("%s", strings.Split(strings.Join(stringAsList, ","), ",")[1:])
	stringAsList = strings.Split(listAsString, ")")
	listAsString = strings.Trim(strings.Trim(strings.Join(stringAsList, ","), "["), "]")
	stringAsList = strings.Split(strings.Trim(listAsString, " "), ",")
	values := make(map[string]string)
	for _, item := range stringAsList {
		keyValuePair := strings.Split(fmt.Sprintf("%v", strings.Split(item, ".")[1:]), "val")
		if len(keyValuePair) == 2 {
			values[strings.TrimSpace(strings.Trim(keyValuePair[0], "["))] = strings.TrimSpace(strings.Trim(keyValuePair[1], "]"))
		}
	}
	return values
}
