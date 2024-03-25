package types

import (
	"encoding/binary"
	"encoding/json"
)

var (
	_ ICrud = &Default{}
)

type Default struct {
	ID     uint64 `json:"id"`
	CMD    Command
	Output string `json:"output" validate:"max=512"`
	Error  Error
}

func (try *Default) SetID(id uint64) {
	try.ID = id
}

func (try *Default) SetKey(s []byte) {

}

func (try *Default) GetKey() []byte {
	k := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(k, try.ID)
	return k
}

func (try *Default) FromJson(s []byte) error {
	return json.Unmarshal(s, try)
}

func (try *Default) FromString(s string) error {
	return nil
}

func (try *Default) FromGob(s []byte) error {
	return nil
}

func (try *Default) Json() ([]byte, error) {

	return json.Marshal(try)
}

func (try *Default) Gob() ([]byte, error) {
	return []byte(""), nil
}

func (try *Default) String() string {
	return ""
}
