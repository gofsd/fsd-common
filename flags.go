package types

var _ IRootFlags = &RootFlags{}

type RootFlags struct {
	Action, Output           uint8
	TestID, UserID, EntityID uint32
	LogByRequestID           uint64
}

func (rf RootFlags) GetUint8(key string) uint8 {
	if key == "action" {
		return rf.Action
	} else if key == "output" {
		return rf.Output
	}
	return 0
}

func (rf RootFlags) GetUint32(key string) uint32 {
	if key == "test" {
		return rf.TestID
	} else if key == "user" {
		return rf.UserID
	} else if key == "entity-id" {
		return rf.EntityID
	}
	return 0
}

func (rf RootFlags) GetUint64(key string) uint64 {
	if key == "log" {
		return rf.LogByRequestID
	}
	return 0
}

func (rf *RootFlags) SetUint8(key string, value uint8) (res uint8) {
	if key == "action" {
		rf.Action = value
		return rf.Action
	} else if key == "output" {
		rf.Output = value
		return rf.Output
	}
	return res
}

func (rf *RootFlags) SetUint32(key string, value uint32) (res uint32) {
	if key == "test" {
		rf.TestID = value
		return rf.TestID
	} else if key == "user" {
		rf.UserID = value
		return rf.UserID
	} else if key == "entity-id" {
		rf.EntityID = value
		return rf.EntityID
	}
	return res
}

func (rf *RootFlags) SetUint64(key string, value uint64) (res uint64) {
	if key == "log" {
		rf.LogByRequestID = value
		return rf.LogByRequestID
	}
	return res
}
