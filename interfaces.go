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
	GetUint8(string) uint8
	GetUint32(string) uint32
	GetUint64(string) uint64
}
