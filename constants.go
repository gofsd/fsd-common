package types

const (
	CREATE uint8 = iota + 1
	READ
	UPDATE
	DELETE
)

const (
	SERVER uint8 = iota + 1
	CLIENT
)
