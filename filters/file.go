package filters

import (
	"fmt"
	"os"
	"sync"
	"transcribe_and_detect_speech/config"
)

type FileFilter struct{}

func (filter FileFilter) dump_to_file(in []byte) []byte {
	cfg := config.Load()
	fileName := cfg.Files.TextFiles[0]

	// Create the file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return nil
	}
	defer file.Close() // Ensure the file is closed when the function ends

	// Write the content to the file
	_, err = file.Write(in)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return nil
	}

	return []byte(fileName)
}

// Process simply returns the input
func (filter FileFilter) Process(in chan []byte, numWorkers int) chan []byte {
	out := make(chan []byte)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for val := range in {
				tmp := filter.dump_to_file(val)
				fmt.Printf("Goroutine %d - FileFilter - Process output: %v \n", workerID, string(tmp))
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
