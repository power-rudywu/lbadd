package command

//go:generate stringer -type=Type

type Type uint16

const (
	Unknown Type = iota
	Select
	Create
	Insert
	Drop
)
