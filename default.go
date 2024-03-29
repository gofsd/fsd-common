package types

var (
	_ ICrud = &Default{}
)

type Default struct {
	CommonID
	RootFlags
}
