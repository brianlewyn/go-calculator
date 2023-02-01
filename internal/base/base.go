package base

// Type is the data type of the result
type Type string

const (
	Int08   = Type("i8")
	Int16   = Type("i16")
	Int32   = Type("i32")
	Int64   = Type("i64")
	Float32 = Type("f32")
	Float64 = Type("f64")
)

type List struct {
	Original  []string
	Temporary []string
	Start     int
	End       int
	// *doubly.Doubly[string]
}

type Result struct {
	Kind    Type
	Int08   int8
	Int16   int16
	Int32   int32
	Int64   int64
	Float32 float32
	Float64 float64
}
