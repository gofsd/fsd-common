package types

type IID interface {
	SetID(uint64)
	GetID() uint64
	SetKey([]byte)
	GetKey() []byte
	Combine(uint32, uint32)
}

type ICrud interface {
	IID
}

type IJustCrud interface {
	JustCreate(ICrud) error
	JustRead(ICrud) error
	JustUpdate(ICrud) error
	JustDelete(ICrud) error
}

type IRootFlags interface {
	GetAction() uint8
	GetEntityID() uint32
	GetOutput() uint8
	GetUserID() uint32
	GetTestID() uint32
	GetLogByID() uint64
	SetAction(uint8)
	SetEntityID(uint32)
	SetOutput(uint8)
	SetUserID(uint32)
	SetTestID(uint32)
	SetLogRequestID(uint64)
}
