package deferhelper

import (
	"fmt"
	"math"
	"os"
	"time"
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

func testWithDuration(d *time.Duration) {
	defer DurationBeforeAfter(d)()
	time.Sleep(1 * time.Second)
}

func ExampleDurationBeforeAfter() {
	var d time.Duration
	testWithDuration(&d)
	fmt.Println(math.Round(d.Seconds()))
	// Output:
	// 1
}
