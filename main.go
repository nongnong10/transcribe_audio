package main

import (
	"transcribe_and_detect_speech/config"
	"transcribe_and_detect_speech/dto"
	"transcribe_and_detect_speech/filters"
)

func main() {
	// Load config
	cfg := config.Load()

	// Create Filters
	extractAudioFilter := filters.ExtractAudioFilter{}
	transcribeFilter := filters.TranscribeFilter{}

	// Create Pipeline
	pipeline := dto.NewPipeline[[]byte, []byte]()
	pipeline.Add(extractAudioFilter)
	pipeline.Add(transcribeFilter)

	// Create request
	in := []byte(cfg.Files.VideoFile)
	out := pipeline.Process(in)

	println(out)
}
