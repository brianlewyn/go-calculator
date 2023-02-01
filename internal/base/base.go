package base

import "github.com/brianlewyn/go-calculator/internal/data"

// List
type List struct {
	Original  []string // *doubly.Doubly[string]
	Temporary []string // *doubly.Doubly[string]
	Start     int
	End       int
}

// Result
type Result struct {
	Kind    data.Kind
	Int08   int8
	Int16   int16
	Int32   int32
	Int64   int64
	Float32 float32
	Float64 float64
}
