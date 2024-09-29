package main

import (
	"transcribe_and_detect_speech/dto"
	"transcribe_and_detect_speech/filters"
)

func main() {
	// Create Filters
	transcribeFilter := filters.TranscribeFilter{}

	// Create Pipeline
	pipeline := dto.NewPipeline()
	pipeline.Add(transcribeFilter)

	// Create request
	filePath := "/home/huy/pipe/transcribe_and_detect_speech/assets/audio/sports_injuries.mp3"
	in := []byte(filePath)
	out := pipeline.Process(in)

	println(out)
}
