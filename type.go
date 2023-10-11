package tisp

type Type uint8

const (
	Nil Type = iota + 1
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Map
	Slice
	Break
)
