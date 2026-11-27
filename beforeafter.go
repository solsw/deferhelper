package deferhelper

import (
	"io"
)

// BeforeAfter calls before immediately and returns after. Use it in a [defer]
// statement to run after when the enclosing function exits:
//
//	defer deferhelper.BeforeAfter(before, after)()
//
// When deferred this way, after can call recover directly to handle a panic.
//
// [defer]: https://go.dev/ref/spec#Defer_statements
func BeforeAfter(before, after func()) func() {
	before()
	return after
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
