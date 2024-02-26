package types

type Command struct {
	Name  []string `json:"name" validate:"required,min=1,max=9,dive,min=2,max=16"`
	Args  []string `json:"args" validate:"omitempty,min=0,max=9,dive,min=1,max=64"`
	Flags []KV     `json:"flags" validate:"omitempty,min=0,max=9"`
}

type KV struct {
	K string `json:"key" validate:"omitempty,min=0,max=9"`
	V string `json:"value" validate:"omitempty,min=0,max=256"`
}

type Pair struct {
	K uint16 `json:"i" validate:"min=1,max=65535"`
	V string `json:"v" validate:"min=1,max=255"`
}

type Error struct {
	Error   string   `json:"error,omitempty" validate:"required"`
	Command *Command `json:"command" validate:"required"`
}

type Cmd func(Command) (string, error)

type CommandResponse struct {
	Id      uint64 `json:"id"`
	Command Command
	Result  any
}

func (cr *CommandResponse) String() string {
	return ""
}

func (cr *CommandResponse) Json() ([]byte, error) {
	return []byte(""), nil
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
	return nil
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
