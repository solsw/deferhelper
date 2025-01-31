package deferhelper

import (
	"fmt"
	"os"
)

func ExampleBeforeAfter() {
	defer BeforeAfter(func() { fmt.Println("before") }, func() { fmt.Println("after") })()
	fmt.Println("in the middle")
	// Output:
	// before
	// in the middle
	// after
}

func ExampleWriteStringBeforeAfter() {
	defer WriteStringBeforeAfter(os.Stdout, "before\n", "after\n")()
	fmt.Println("in the middle")
	// Output:
	// before
	// in the middle
	// after
}
