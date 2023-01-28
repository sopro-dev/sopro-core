package resampler

import "fmt"

// Println prints the items if the transcoder is in verbose mode
func (rs *Resampler) Println(items ...any) {
	if rs.Verbose {
		fmt.Println(items...)
	}
}
