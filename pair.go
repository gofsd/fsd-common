package types

type Pair struct {
	RootFlags
	CommonID[string]
	V string `json:"v" validate:"min=1,max=255"`
}
