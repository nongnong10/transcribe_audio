package tests

import (
	"bytes"
	"testing"
	"transcribe_and_detect_speech/filters"
)

func assertBytes(t *testing.T, got, want []byte) {
	t.Helper()
	if bytes.Compare(got, want) != 0 {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestTranscribeFilterProcess(t *testing.T) {
	// Create Filters
	transcribeFilter := filters.TranscribeFilter{}

	// File path
	filePath := "/home/huy/pipe/transcribe_and_detect_speech/assets/audio/sports_injuries.mp3"

	in := make(chan []byte)
	go func() {
		in <- []byte(filePath)
		close(in)
	}()
	got := <-transcribeFilter.Process(in)
	println(got)
}

func TestExtractAudioFilterProcess(t *testing.T) {
	// Create Filters
	extractAudioFilter := filters.ExtractAudioFilter{}

	// File path
	filePath := "/home/huy/pipe/transcribe_and_detect_speech/assets/video/video.mp4"

	in := make(chan []byte)
	go func() {
		in <- []byte(filePath)
		close(in)
	}()
	got := <-extractAudioFilter.Process(in)
	println(got)
}
