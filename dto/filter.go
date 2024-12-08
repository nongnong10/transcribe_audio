package dto

// Filter interface using generics for input and output types
type Filter[In any, Out any] interface {
	Process(in chan In, numWorkers int) chan Out
}
