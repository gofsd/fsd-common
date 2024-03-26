package types

type RootFlags struct {
	Action, Output     uint8
	TestID, UserID, ID uint32
	LogByRequestID     uint64
}
