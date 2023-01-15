package resampler

import "fmt"

// Println prints the items if the transcoder is in verbose mode
func (t *Resampler) Println(items ...any) {
	if t.Verbose {
		fmt.Println(items...)
	}
}
