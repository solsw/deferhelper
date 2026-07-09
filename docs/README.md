# deferhelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/deferhelper.svg)](https://pkg.go.dev/github.com/solsw/deferhelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/deferhelper)

Package `deferhelper` contains helpers for Go's
[defer](https://go.dev/ref/spec#Defer_statements) statement. They let a single
deferred call run one action **before** the enclosing function body and another
**after** it returns — a convenient pattern for tracing, logging and setup/teardown.

## Installation

```sh
go get github.com/solsw/deferhelper
```

```go
import "github.com/solsw/deferhelper"
```

## How it works

`defer` evaluates its arguments immediately but runs the deferred call when the
function returns. [`BeforeAfter`](https://github.com/solsw/deferhelper/blob/main/beforeafter.go)
exploits this: it takes a `before` and an `after` function, calls `before` right
away (while the arguments are being evaluated), and returns a closure that calls
`after`. Deferring that closure — and calling it — schedules `after` for
function exit:

```go
func Example() {
    defer deferhelper.BeforeAfter(before, after)()
    // do something
}
```

Note the trailing `()`: `BeforeAfter` returns the function that `defer` will run.

## Functions

| Function | Signature | Purpose |
| --- | --- | --- |
| [`BeforeAfter`](https://github.com/solsw/deferhelper/blob/main/beforeafter.go) | `BeforeAfter(before, after func()) func()` | run `before` now, `after` on function exit |
| [`WriteBeforeAfter`](https://github.com/solsw/deferhelper/blob/main/beforeafter.go) | `WriteBeforeAfter(w io.Writer, before, after []byte) func()` | write `before`/`after` byte slices to `w` around the body |
| [`WriteStringBeforeAfter`](https://github.com/solsw/deferhelper/blob/main/beforeafter.go) | `WriteStringBeforeAfter(w io.Writer, before, after string) func()` | like `WriteBeforeAfter` but takes strings |

`WriteBeforeAfter` and `WriteStringBeforeAfter` panic if the write to `w` fails,
and `WriteBeforeAfter` panics if `w` is `nil`. Empty `before`/`after` values
produce no write.

## Usage

### Arbitrary before/after actions

```go
func Example() {
    defer deferhelper.BeforeAfter(
        func() { fmt.Println("before") },
        func() { fmt.Println("after") },
    )()
    fmt.Println("in the middle")
    // Output:
    // before
    // in the middle
    // after
}
```

### Writing markers to an io.Writer

```go
func Example() {
    defer deferhelper.WriteStringBeforeAfter(os.Stdout, "before\n", "after\n")()
    fmt.Println("in the middle")
    // Output:
    // before
    // in the middle
    // after
}
```

Either side may be empty, so the same call can emit only a prefix or only a suffix:

```go
defer deferhelper.WriteStringBeforeAfter(os.Stdout, "only before\n", "")()
defer deferhelper.WriteStringBeforeAfter(os.Stdout, "", "only after\n")()
```

### Tracing function entry and exit

Combined with a caller-name helper, this makes a compact entry/exit tracer:

```go
func work() {
    defer deferhelper.WriteStringBeforeAfter(os.Stdout,
        runtimehelper.CallerName()+" start\n",
        runtimehelper.CallerName()+" finish\n")()
    fmt.Println("work()")
}
```
