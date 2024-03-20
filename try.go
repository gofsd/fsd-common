package types

import (
	"encoding/binary"
	"encoding/json"
)

type TryOptions struct {
	ID       uint64
	Apply    string
	Destroy  string
	Name     string
	Checks   []Check
	Equals   []Equal
	Output   string
	Triggers map[string]bool
}

type Equal struct {
	Statement string
	Got       string
	Want      string
}

type Check struct {
	Statement string
	Pass      bool
}

func (cmd *TryOptions) SetCheck(statements []string, pass []bool) {
	for i, v := range statements {
		cmd.Checks = append(cmd.Checks, Check{
			Statement: v,
			Pass:      pass[i],
		})
	}
}

func (cmd *TryOptions) SetEqual(statements, gots, wants []string) {
	for i, v := range statements {
		cmd.Equals = append(cmd.Equals, Equal{
			Statement: v,
			Got:       gots[i],
			Want:      wants[i],
		})
	}
}

func (try *TryOptions) SetID(id uint64) {
	try.ID = id
}

func (try *TryOptions) SetKey(s []byte) {

}

func (try *TryOptions) GetKey() []byte {
	k := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(k, try.ID)
	return k
}

func (try *TryOptions) FromJson(s []byte) error {
	return json.Unmarshal(s, try)
}

func (try *TryOptions) FromString(s string) error {
	return nil
}

func (try *TryOptions) FromGob(s []byte) error {
	return nil
}

func (try *TryOptions) Json() ([]byte, error) {

	return json.Marshal(try)
}

func (try *TryOptions) Gob() ([]byte, error) {
	return []byte(""), nil
}

func (try *TryOptions) String() string {
	return ""
}
