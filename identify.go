package types

import "sync/atomic"

var (
	TotalRequestID, RequestID, ID atomic.Uint64
)
