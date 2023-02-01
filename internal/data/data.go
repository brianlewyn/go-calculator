package data

type Kind int8

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
