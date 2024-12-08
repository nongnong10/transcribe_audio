package filters

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"transcribe_and_detect_speech/config"
)

type ExtractAudioFilter struct{}

func (filter ExtractAudioFilter) extractAudio(in []byte) []byte {
	videoFilePath := string(in)
	basePath, _ := filepath.Abs("./assets/video/")
	videoFilePath = filepath.Join(basePath, videoFilePath)

	cfg := config.Load()
	audioFilePath := cfg.Files.AudioFiles[0]
	basePath, _ = filepath.Abs("./assets/audio/")
	audioFilePath = filepath.Join(basePath, audioFilePath)

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
func (filter ExtractAudioFilter) Process(in chan []byte, numWorkers int) chan []byte {
	out := make(chan []byte)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for val := range in {
				// Xử lý extract audio
				tmp := filter.extractAudio(val)
				fmt.Printf("Goroutine %d - ExtractAudioFilter - Process output: %v \n", workerID, string(tmp))
				out <- tmp
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
