package types

import "encoding/json"

type Command struct {
	Name  []string `json:"name" validate:"required,min=1,max=9,dive,min=2,max=16"`
	Args  []string `json:"args" validate:"omitempty,min=0,max=9,dive,min=1,max=512"`
	Flags []KV     `json:"flags" validate:"omitempty,min=1,max=16"`
}

type Error struct {
	Default
	Code    uint16
	Message string   `json:"error,omitempty" validate:"required"`
	Command *Command `json:"command" validate:"required"`
}

type Testing struct {
	Default
}

func (cmd *Command) SetFlag(key, value string) *Command {
	cmd.Flags = append(cmd.Flags, KV{
		Key:   key,
		Value: value,
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
	Key   string `json:"k" validate:"omitempty,min=1,max=16"`
	Value string `json:"v" validate:"omitempty,min=0,max=32"`
}

func (e *Error) Error() string {
	var data []byte
	json.Unmarshal(data, e)
	return string(data)
}

type Cmd func(Command) (string, error)

type CommandResponse struct {
	Default
	CommonID
	Command  Command
	Result   any
	Duration int64
}
