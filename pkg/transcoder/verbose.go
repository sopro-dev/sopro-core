package transcoder

import "fmt"

// Println prints the items if the transcoder is in verbose mode
func (t *Transcoder) Println(items ...any) {
	if t.Verbose {
		fmt.Println(items...)
	}
}
