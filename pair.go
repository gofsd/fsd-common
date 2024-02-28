package types

import "encoding/json"

type Pair struct {
	K uint64 `json:"k" validate:"min=1,max=18446744073709551615"`
	V string `json:"v" validate:"min=1,max=255"`
}

func (pair *Pair) String() string {
	return ""
}

func (pair *Pair) Json() ([]byte, error) {
	return json.Marshal(pair)
}

func (pair *Pair) Gob() ([]byte, error) {
	return []byte(""), nil
}

func (pair *Pair) GetValue() any {
	return pair
}

func (pair *Pair) Update() error {
	return nil
}

func (pair *Pair) GetKey() []byte {
	return nil
}
func (pair *Pair) FromJson(s []byte) error {
	return json.Unmarshal(s, *pair)
}

func (pair *Pair) FromString(s string) error {
	return nil
}
func (pair *Pair) FromGob(s []byte) error {
	return nil
}

func (pair *Pair) SetID(id uint64) {
	pair.K = id
}

func (pair *Pair) SetKey(key []byte) {
}
