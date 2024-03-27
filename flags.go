package types

var _ IRootFlags = RootFlags{}

type RootFlags struct {
	Action, Output           uint8
	TestID, UserID, EntityID uint32
	LogByRequestID           uint64
}

func (rf RootFlags) GetAction() uint8 {
	return rf.Action
}

func (rf RootFlags) GetOutput() uint8 {
	return rf.Action
}
func (rf RootFlags) GetUserID() uint32 {
	return rf.EntityID
}

func (rf RootFlags) GetEntityID() uint32 {
	return rf.EntityID
}

func (rf RootFlags) GetLogByID() uint64 {
	return rf.LogByRequestID
}

func (rf RootFlags) GetTestID() uint32 {
	return rf.TestID
}

func (rf *RootFlags) SetAction(action uint8) {
	rf.Action = action
}

func (rf *RootFlags) SetOutput(output uint8) {
	rf.Output = output
}
func (rf *RootFlags) SetUserID(userID uint32) {
	rf.UserID = userID
}

func (rf *RootFlags) SetEntityID(entityID uint32) {
	rf.EntityID = entityID
}

func (rf *RootFlags) SetLogRequestID(logRequestID uint64) {
	rf.LogByRequestID = logRequestID
}

func (rf *RootFlags) SetTestID(testID uint32) {

	rf.TestID = testID
}
