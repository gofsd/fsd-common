package types

var _ IRootFlags = &RootFlags{}

type RootFlags struct {
	Action, Output           uint8
	TestID, UserID, EntityID uint32
	LogByRequestID           uint64
}

func (rf *RootFlags) GetUint8(key string) *uint8 {
	if key == "action" {
		return &rf.Action
	} else {
		return &rf.Output
	}
}

func (rf *RootFlags) GetUint32(key string) *uint32 {
	if key == "test" {
		return &rf.TestID
	} else if key == "user" {
		return &rf.UserID
	} else {
		return &rf.EntityID
	}
}

func (rf *RootFlags) GetUint64(key string) *uint64 {
	return &rf.LogByRequestID
}
