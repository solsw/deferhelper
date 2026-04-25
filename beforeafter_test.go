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
