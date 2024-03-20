package types

import (
	"encoding/binary"
	"encoding/json"
)

type Pair struct {
	K uint64 `json:"k" validate:"min=1,max=18446744073709551615"`
	V string `json:"v" validate:"min=1,max=255"`
}

func (pair *Pair) SetID(id uint64) {
	pair.K = id
}

func (Pair *Pair) SetKey(s []byte) {
	Pair.K = binary.BigEndian.Uint64(s)
}

func (Pair *Pair) GetKey() []byte {
	k := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(k, Pair.K)
	return k
}

func (pair *Pair) FromJson(s []byte) error {
	return json.Unmarshal(s, pair)
}

func (pair *Pair) FromString(s string) error {
	(*pair).V = s
	return nil
}

func (pair *Pair) FromGob(s []byte) error {
	return nil
}

func (pair *Pair) Json() ([]byte, error) {

	return json.Marshal(pair)
}

func (pair *Pair) Gob() ([]byte, error) {
	return []byte(""), nil
}

func (pair *Pair) String() string {
	return ""
}

func (pair *Pair) GetValue() any {
	return pair
}

func (pair *Pair) Update() error {
	return nil
}

func (Pair *Pair) ToPretifiedJson() (b []byte, e error) {
	b, e = json.MarshalIndent(Pair.V, "", "  ")

	return b, e
}

func (Pair *Pair) ToJson() (b []byte, e error) {
	b, e = json.Marshal(Pair)
	return b, e
}

func (Pair *Pair) JustUpdate() error {
	return nil
}

func (Pair *Pair) FromSlice(k []byte) (e error) {
	e = json.Unmarshal(k, Pair.V)

	return
}

func (Pair *Pair) Tes() []byte {
	k := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(k, Pair.K)
	return k
}
