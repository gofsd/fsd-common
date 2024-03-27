package types

type Pair struct {
	K uint64 `json:"k" validate:"min=1,max=18446744073709551615"`
	V string `json:"v" validate:"min=1,max=255"`
}
