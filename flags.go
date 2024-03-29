package types

type RootFlags struct {
	Action   uint8  `json:"action" mapstructure:"action"`
	Output   uint8  `json:"output" mapstructure:"output"`
	UserID   uint32 `json:"user" mapstructure:"user"`
	EntityID uint32 `json:"entity" mapstructure:"entity"`
}
