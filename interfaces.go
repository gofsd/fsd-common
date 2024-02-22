package client

type ICrud interface {
	SetID(uint64)
	SetKey([]byte)
	GetKey() []byte
	FromJson([]byte) error
	FromGob([]byte) error
	FromString([]byte) error
	Json() []byte
	Gob() []byte
	String() string
}
