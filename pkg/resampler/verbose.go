package resampler

import "fmt"

// Println prints the items if the transcoder is in verbose mode
<<<<<<< HEAD
func (rs *Resampler) Println(items ...any) {
	if rs.Verbose {
=======
func (t *Resampler) Println(items ...any) {
	if t.Verbose {
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		fmt.Println(items...)
	}
}
