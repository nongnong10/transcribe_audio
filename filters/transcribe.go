package filters

import (
	"context"
	"fmt"
	aai "github.com/AssemblyAI/assemblyai-go-sdk"
	"os"
	"transcribe_and_detect_speech/config"
)

type TranscribeFilter struct{}

func (filter TranscribeFilter) transcribe(in []byte) []byte {
	cfg := config.Load()
	client := aai.NewClient(cfg.APIKey)

	params := &aai.TranscriptOptionalParams{
		SpeakerLabels: aai.Bool(true),
	}

	fmt.Printf("TranscribeFilter - transcribe input: %v \n", string(in))
	f, err := os.Open(string(in))

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	transcript, err := client.Transcripts.TranscribeFromReader(context.Background(), f, params)

	fmt.Println(*transcript.Text)

	for _, utterance := range transcript.Utterances {
		fmt.Printf("Speaker %v: %v\n", *utterance.Speaker, *utterance.Text)
	}

	return []byte(*transcript.Text)
}

// Process simply returns the input
func (filter TranscribeFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		for val := range in {
			tmp := filter.transcribe(val)
			fmt.Printf("TranscribeFilter - Process output: %v \n", string(tmp))
			out <- tmp
		}
		close(out)
	}()
	return out
}
