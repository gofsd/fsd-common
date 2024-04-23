package types

var (
//_ ICrud = &Default{}
)

type Default struct {
	CommonID  `mapstructure:",squash"`
	RootFlags `mapstructure:",squash"`
}
