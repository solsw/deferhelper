package deferhelper

import (
	"fmt"
	"os"
	"testing"
)

func TestBeforeAfterRecover(t *testing.T) {
	const want = "panic in body"
	defer func() {
		if escaped := recover(); escaped != nil {
			t.Errorf("panic escaped after callback: %v", escaped)
		}
	}()
	defer BeforeAfter(func() {}, func() {
		if got := recover(); got != want {
			t.Errorf("after recovered %v, want %q", got, want)
		}
	})()
	panic(want)
}

func ExampleBeforeAfter() {
	defer BeforeAfter(func() { fmt.Println("before") }, func() { fmt.Println("after") })()
	fmt.Println("in the middle")
	// Output:
	// before
	// in the middle
	// after
}

func ExampleWriteStringBeforeAfter() {
	defer WriteStringBeforeAfter(os.Stdout, "", "")()
	defer WriteStringBeforeAfter(os.Stdout, "only before\n", "")()
	defer WriteStringBeforeAfter(os.Stdout, "before\n", "after\n")()
	defer WriteStringBeforeAfter(os.Stdout, "", "only after\n")()
	fmt.Println("in the middle")
	// Output:
	// only before
	// before
	// in the middle
	// only after
	// after
}
