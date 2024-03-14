package types

type ICrud interface {
	SetID(uint64)
	SetKey([]byte)
	GetKey() []byte
	FromJson([]byte) error
	FromString(string) error
	FromGob([]byte) error
	Json() ([]byte, error)
	Gob() ([]byte, error)
	String() string
}

type IJustCrud interface {
	JustCreate(ICrud) error
	JustRead(ICrud) error
	JustUpdate(ICrud) error
	JustDelete(ICrud) error
}
