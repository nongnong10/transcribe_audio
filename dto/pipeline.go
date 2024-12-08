package dto

// Pipeline struct supporting different input and output types for each step
type Pipeline[In any, Out any] struct {
	head chan In
	tail chan Out
}

// NewPipeline creates a new pipeline instance
func NewPipeline[In any, Out any]() *Pipeline[In, Out] {
	return &Pipeline[In, Out]{}
}

// Add adds a new pipeline step that transforms input from one type to another
func (p *Pipeline[In, Out]) Add(filter Filter[In, Out], numWorkers int) {
	// Case 1: Pipeline is empty
	if p.tail == nil {
		p.head = make(chan In)
		p.tail = filter.Process(p.head, numWorkers)
	} else {
		// Case 2: Pipeline is not empty, continue to add and process
		// Update the pipeline by adding a new filter step
		p.tail = filter.Process(any(p.tail).(chan In), numWorkers)
	}
}

// Process executes the pipeline and transforms the input through all steps
func (p *Pipeline[In, Out]) Process(in In) Out {
	// Case 1: Pipeline is empty
	if p.head == nil {
		return any(in).(Out) // Trả về đầu vào nếu không có bước nào trong pipeline
	}
	p.head <- in
	close(p.head)
	return <-p.tail
}
