package dto

// Pipeline keeps references to the start and end of pipeline
type Pipeline struct {
	head chan []byte
	tail chan []byte
}

// NewPipeline
func NewPipeline() *Pipeline {
	return &Pipeline{}
}

// Add adds a new pipeline step
func (p *Pipeline) Add(filter Filter) {
	// Case 1: Pipeline is empty
	if p.tail == nil {
		p.head = make(chan []byte)
		p.tail = filter.Process(p.head)
	} else {
		// Case 2: Pipeline is not empty, continue to add and process
		p.tail = filter.Process(p.tail)
	}
}

// Process executes the pipeline
func (p *Pipeline) Process(in []byte) (out []byte) {
	// Case 1: Pipeline is empty
	if p.head == nil {
		return in
	}
	p.head <- in
	close(p.head)
	return <-p.tail
}
