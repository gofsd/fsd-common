package types

import "fmt"

const (
	ENOTEXIST uint16 = iota + 1
	EEXIST
	EUNKNOWN
)

type CustomErrors map[uint16]string

var ErMap CustomErrors = CustomErrors{}

func (err *CustomErrors) AddError(code uint16, message string) {
	if len(message) > 30 {
		panic("Error message can't be longer than 30 symbols")
	}
	(*err)[code] = message
}

func (err *CustomErrors) GetError(code uint16, args ...any) error {
	if value, ok := (*err)[code]; ok {
		return fmt.Errorf(value, args...)
	}
	value, _ := (*err)[EUNKNOWN]
	return fmt.Errorf(value, args)
}
func init() {
	ErMap.AddError(ENOTEXIST, "File not found")
	ErMap.AddError(EUNKNOWN, "Unknown error with code: %d and arguments: %#v")
}
