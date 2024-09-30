package tests

import (
	"testing"
	"transcribe_and_detect_speech/dto"
	"transcribe_and_detect_speech/filters"
)

func TestPipelineWithMultiFilter(t *testing.T) {
	// Create Filters
	transcribeFilter := filters.TranscribeFilter{}

	// Create Pipeline
	pipeline := dto.NewPipeline[[]byte, []byte]()
	pipeline.Add(transcribeFilter)

	in := []byte("Hello World")
	out := pipeline.Process(in)

	println(out)
}
