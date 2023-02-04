package data

type Kind uint8

// Kind is the data type of the result
const (
	_ = Kind(iota)
	Int08
	Int16
	Int32
	Int64
	Float32
	Float64
)

// DigitLimit is the limit of digits of a float64 and dot
const DigitLimit uint16 = 617

// It is the limit of digits of a data type
const (
	F64 uint16 = 308
	F32 uint8  = 38
	I64 uint8  = 18
	I32 uint8  = 10
	I16 uint8  = 5
	I08 uint8  = 3
)
