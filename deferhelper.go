package deferhelper

import (
	"bufio"
	"io"

	"github.com/solsw/errorhelper"
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
func WriteBeforeAfter(w io.Writer, before, after []byte) func() {
	if len(before) == 0 && len(after) == 0 {
		return func() {}
	}
	return BeforeAfter(
		func() {
			if len(before) > 0 {
				_ = errorhelper.Must(w.Write(before))
			}
		},
		func() {
			if len(after) > 0 {
				_ = errorhelper.Must(w.Write(after))
			}
		},
	)
}

// WriteStringBeforeAfter is like [BeforeAfter] but writes 'before' and
// 'after' strings to 'w' before and after executing the function code.
func WriteStringBeforeAfter(w io.Writer, before, after string) func() {
	if before == "" && after == "" {
		return func() {}
	}
	bw := bufio.NewWriter(w)
	return BeforeAfter(
		func() {
			if before != "" {
				_ = errorhelper.Must(bw.WriteString(before))
				errorhelper.Must0(bw.Flush())
			}
		},
		func() {
			if after != "" {
				_ = errorhelper.Must(bw.WriteString(after))
				errorhelper.Must0(bw.Flush())
			}
		},
	)
}
