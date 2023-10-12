package tisp

type flag uint8

const (
	_ flag = iota
	flagNil
	flagBool
	flagInt
	flagInt8
	flagInt16
	flagInt32
	flagInt64
	flagUint
	flagUint8
	flagUint16
	flagUint32
	flagUint64
	flagFloat32
	flagFloat64
	flagString
	flagMap
	flagSlice
	flagBreak
)
