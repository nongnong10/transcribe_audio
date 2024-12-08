package tests

import (
	"testing"
	"transcribe_and_detect_speech/dto"
	"transcribe_and_detect_speech/filters"
)

func TestPipelineWithMultiFilter(t *testing.T) {
	// Create Filters
	transcribeFilter := filters.TranscribeFilter{}
	extractAudioFilter := filters.ExtractAudioFilter{}

	// Create Pipeline
	pipeline := dto.NewPipeline[[]byte, []byte]()
	pipeline.Add(extractAudioFilter, 1)
	pipeline.Add(transcribeFilter, 3)

	// Create request
	filePath := []string{"/home/huy/pipe/transcribe_and_detect_speech/assets/video/video.mp4", "/home/huy/pipe/transcribe_and_detect_speech/assets/video/video1.mp4", "/home/huy/pipe/transcribe_and_detect_speech/assets/video/video2.mp4"}

	for _, path := range filePath {
		go func() {
			in := []byte(path)
			out := pipeline.Process(in)
			println(out)
		}()
	}
}
