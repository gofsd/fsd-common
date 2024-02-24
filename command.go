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
