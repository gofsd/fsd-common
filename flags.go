package types

type RootFlags struct {
	Action   uint8  `json:"action"`
	Output   uint8  `json:"output"`
	UserID   uint32 `json:"user"`
	EntityID uint32 `json:"entity"`
}
