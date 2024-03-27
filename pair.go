package types

type Pair struct {
	RootFlags
	ID
	V string `json:"v" validate:"min=1,max=255"`
}
