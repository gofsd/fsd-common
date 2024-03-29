package types

type TryOptions struct {
	Default  `mapstructure:",squash"`
	Apply    string
	Destroy  string
	Name     string
	Checks   []Check
	Equals   []Equal
	Output   string
	Triggers map[string]bool
}

type Equal struct {
	Statement string
	Got       string
	Want      string
}

type Check struct {
	Statement string
	Pass      bool
}

func (cmd *TryOptions) SetCheck(statements []string, pass []bool) {
	for i, v := range statements {
		cmd.Checks = append(cmd.Checks, Check{
			Statement: v,
			Pass:      pass[i],
		})
	}
}

func (cmd *TryOptions) SetEqual(statements, gots, wants []string) {
	for i, v := range statements {
		cmd.Equals = append(cmd.Equals, Equal{
			Statement: v,
			Got:       gots[i],
			Want:      wants[i],
		})
	}
}
