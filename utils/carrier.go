package utils

import "google.golang.org/grpc/metadata"

// MDTraceCarrier is an implementation for Tracing carrier, it will hold traceID, etc
type MDTraceCarrier metadata.MD

// Set part of the Carrier interface
func (md MDTraceCarrier) Set(key, value string) { _ = "STUB: not implemented"; return }

// ForeachKey part of the Carrier interface
func (md MDTraceCarrier) ForeachKey(handler func(key, value string) error) error {
	_ = "STUB: not implemented"
	return nil
}
