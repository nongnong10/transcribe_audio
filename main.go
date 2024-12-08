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
	fileFilter := filters.FileFilter{}

	// Create Pipeline
	pipeline := dto.NewPipeline[[]byte, []byte]()
	pipeline.Add(extractAudioFilter, 1)
	pipeline.Add(transcribeFilter, 1)
	pipeline.Add(fileFilter, 1)

	// Create request
	in := []byte(cfg.Files.VideoFiles[0])
	out := pipeline.Process(in)

	println(out)
}
