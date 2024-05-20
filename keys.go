// keys.go
package correlationcontexthandler

type CorrelationId string
type TraceId string
type FlowId string

const (
	CORRELATION_ID CorrelationId = "CorrelationId"
	TRACE_ID       TraceId       = "TraceId"

	FLOW_ID FlowId = "FlowId"
)
