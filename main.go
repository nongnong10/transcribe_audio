package main

import (
	"transcribe_and_detect_speech/dto"
	"transcribe_and_detect_speech/filters"
)

func main() {
	// Create Filters
	extractAudioFilter := filters.ExtractAudioFilter{}
	transcribeFilter := filters.TranscribeFilter{}

	// Create Pipeline
	pipeline := dto.NewPipeline[[]byte, []byte]()
	pipeline.Add(extractAudioFilter)
	pipeline.Add(transcribeFilter)

	// Create request
	filePath := "/home/huy/pipe/transcribe_and_detect_speech/assets/video/video.mp4"
	in := []byte(filePath)
	out := pipeline.Process(in)

	println(out)
}
