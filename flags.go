package types

type RootFlags struct {
	Action, Command, Output uint8
	UserID, EntityID        uint32
	Flag                    []string
}
