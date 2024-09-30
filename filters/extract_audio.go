package filters

import (
	"fmt"
	"os"
	"os/exec"
)

type ExtractAudioFilter struct{}

func (filter ExtractAudioFilter) extractAudio(in []byte) []byte {
	videoFilePath := string(in)

	nameAudio := "audio"
	audioFilePath := "/home/huy/pipe/transcribe_and_detect_speech/assets/audio/" + nameAudio + ".mp3"

	// Remove file if exists
	if _, err := os.Stat(audioFilePath); err == nil {
		err := os.Remove(audioFilePath)
		if err != nil {
			fmt.Printf("Error deleting file: %w", err)
		}
		fmt.Println("File deleted:", audioFilePath)
	}

	// Run the command
	cmd := exec.Command("ffmpeg", "-i", videoFilePath, "-q:a", "0", "-map", "a", audioFilePath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error extracting audio: %w", err)
		return nil
	}

	return []byte(audioFilePath)
}

// Process simply returns the input
func (fiter ExtractAudioFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		for val := range in {
			tmp := fiter.extractAudio(val)
			fmt.Printf("ExtractAudioFilter - Process output: %v \n", string(tmp))
			out <- tmp
		}
		close(out)
	}()
	return out
}
