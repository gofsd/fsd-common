package types

import "encoding/json"

type Command struct {
	Name  []string `json:"name" validate:"required,min=1,max=9,dive,min=2,max=16"`
	Args  []string `json:"args" validate:"omitempty,min=0,max=9,dive,min=1,max=512"`
	Flags []KV     `json:"flags" validate:"omitempty,min=1,max=16"`
}

func (cmd *Command) SetFlag(key, value string) *Command {
	cmd.Flags = append(cmd.Flags, KV{
		K: key,
		V: value,
	})
	return cmd
}

func (cmd *Command) SetArg(value string) *Command {
	cmd.Args = append(cmd.Args, value)
	return cmd

}

func (cmd *Command) SetArgs(args []string) *Command {
	cmd.Args = args
	return cmd
}

func (cmd *Command) SetName(name string) *Command {
	cmd.Name = append(cmd.Name, name)
	return cmd
}

type KV struct {
	K string `json:"key" validate:"omitempty,min=1,max=16"`
	V string `json:"value" validate:"omitempty,min=0,max=32"`
}

type Error struct {
	Code    uint16
	Message string   `json:"error,omitempty" validate:"required"`
	Command *Command `json:"command" validate:"required"`
}

func (e *Error) Error() string {
	var data []byte
	json.Unmarshal(data, e)
	return string(data)
}

type Cmd func(Command) (string, error)

type CommandResponse struct {
	Id       uint64 `json:"id"`
	Command  Command
	Result   any
	Duration int64
}

func (cr *CommandResponse) String() string {
	return ""
}

func (cr *CommandResponse) Json() ([]byte, error) {

	return json.Marshal(cr)
}

func (cr *CommandResponse) Gob() ([]byte, error) {
	return []byte(""), nil
}

func (cr *CommandResponse) GetValue() any {
	return cr
}

func (cr *CommandResponse) Update() error {
	return nil
}

func (cr *CommandResponse) GetKey() []byte {
	return nil
}
func (cr *CommandResponse) FromJson(s []byte) error {
	return json.Unmarshal(s, cr)
}

func (cr *CommandResponse) FromString(s string) error {
	return nil
}
func (cr *CommandResponse) FromGob(s []byte) error {
	return nil
}

func (cr *CommandResponse) SetID(id uint64) {
	cr.Id = id
}

func (cr *CommandResponse) SetKey(key []byte) {
}

func (cr *Command) String() string {
	return ""
}

func (cr *Command) Json() ([]byte, error) {

	return json.Marshal(cr)
}

func (cr *Command) Gob() ([]byte, error) {
	return []byte(""), nil
}

func (cr *Command) GetValue() any {
	return cr
}

func (cr *Command) Update() error {
	return nil
}

func (cr *Command) GetKey() []byte {
	return nil
}
func (cr *Command) FromJson(s []byte) error {
	return json.Unmarshal(s, cr)
}

func (cr *Command) FromString(s string) error {
	return nil
}
func (cr *Command) FromGob(s []byte) error {
	return nil
}

func (cr *Command) SetID(id uint64) {
}

func (cr *Command) SetKey(key []byte) {
}
