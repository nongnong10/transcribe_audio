package filters

type ExtractAudioFilter struct{}

// Process simply returns the input
func (fiter ExtractAudioFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		for val := range in {
			out <- val
		}
		close(out)
	}()
	return out
}
