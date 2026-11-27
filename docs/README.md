# deferhelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/deferhelper.svg)](https://pkg.go.dev/github.com/solsw/deferhelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/deferhelper)

Package `deferhelper` contains helpers for Go's
[defer](https://go.dev/ref/spec#Defer_statements) statement. Use them to pair
setup with cleanup or to trace function entry and exit.

## Installation

```sh
go get github.com/solsw/deferhelper
```

```go
import "github.com/solsw/deferhelper"
```

## How it works

When execution reaches the `defer` statement below, `BeforeAfter` calls
`before` immediately and returns `after`. The trailing `()` makes `after()`
the deferred call.

```go
func work(before, after func()) {
    defer deferhelper.BeforeAfter(before, after)()
    // do something
}
```

`after` runs when `work` returns or a panic unwinds it. Because `after` is
itself deferred, it can call `recover()` directly to handle the panic.

## Functions

| Function | Signature | Purpose |
| --- | --- | --- |
| [`BeforeAfter`](https://github.com/solsw/deferhelper/blob/main/beforeafter.go) | `BeforeAfter(before, after func()) func()` | run `before` now, `after` on function exit |
| [`WriteBeforeAfter`](https://github.com/solsw/deferhelper/blob/main/beforeafter.go) | `WriteBeforeAfter(w io.Writer, before, after []byte) func()` | write `before`/`after` byte slices to `w` around the body |
| [`WriteStringBeforeAfter`](https://github.com/solsw/deferhelper/blob/main/beforeafter.go) | `WriteStringBeforeAfter(w io.Writer, before, after string) func()` | like `WriteBeforeAfter` but takes strings |

`WriteBeforeAfter` and `WriteStringBeforeAfter` both panic if `w` is `nil`
or a write fails. Empty `before`/`after` values
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
