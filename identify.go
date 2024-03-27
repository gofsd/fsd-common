package types

import "sync/atomic"

var (
	TotalRequestID, RequestID atomic.Uint64
)
