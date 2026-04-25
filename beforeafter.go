package deferhelper

import (
	"io"
)

// BeforeAfter, if called in the following manner
// from a [defer] statement at the very beginning of a function,
//
//		func Example() {
//		  defer deferhelper.BeforeAfter(before(), after())()
//	    // do something
//		}
//
// will call 'before' function before executing the function code
// and 'after' function after executing the function code.
//
// [defer]: https://go.dev/ref/spec#Defer_statements
func BeforeAfter(before, after func()) func() {
	before()
	return func() {
		after()
	}
}

// WriteBeforeAfter is like [BeforeAfter] but writes 'before' and
// 'after' slices to 'w' before and after executing the function code.
// It panics if a write to 'w' fails.
func WriteBeforeAfter(w io.Writer, before, after []byte) func() {
	if w == nil {
		panic("deferhelper: nil writer")
	}
	if len(before) == 0 && len(after) == 0 {
		return func() {}
	}
	return BeforeAfter(
		func() {
			if len(before) > 0 {
				if _, err := w.Write(before); err != nil {
					panic(err)
				}
			}
		},
		func() {
			if len(after) > 0 {
				if _, err := w.Write(after); err != nil {
					panic(err)
				}
			}
		},
	)
}

// WriteStringBeforeAfter is like [WriteBeforeAfter] but accepts strings.
// It panics if a write to 'w' fails.
func WriteStringBeforeAfter(w io.Writer, before, after string) func() {
	return WriteBeforeAfter(w, []byte(before), []byte(after))
}
