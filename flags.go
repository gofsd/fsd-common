package types

type RootFlags struct {
	Action   uint8  `mapstructure:"action"`
	Output   uint8  `mapstructure:"output"`
	UserID   uint32 `mapstructure:"user"`
	EntityID uint32 `mapstructure:"entity"`
}
