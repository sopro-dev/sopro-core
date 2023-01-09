package transcoder

import (
	"fmt"
)

func (t *Transcoder) Println(items ...any) {
	if t.Verbose {
		fmt.Println(items...)
	}
}
