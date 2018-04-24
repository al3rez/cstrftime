# ⌛️cstrftime [![GoDoc](https://godoc.org/github.com/cooldrip/cstrftime?status.svg)](https://godoc.org/github.com/cooldrip/cstrftime) [![codecov](https://codecov.io/gh/cooldrip/cstrftime/branch/master/graph/badge.svg)](https://codecov.io/gh/cooldrip/cstrftime)

The missing strftime in Go.

## Installation

```
> go get github.com/cooldrip/cstrftime
```

## Example

```go
package main

import "github.com/cooldrip/cstrftime"

func main() {
	t := time.Now()
	fmt.Println(cstrftime.Format("%d", t)) // 25
	fmt.Println(cstrftime.Format("%M", t)) // 54
	// etc.
}
```

## Directives

### Date

| Directive | Example                  |
| --------- | ------------------------ |
| `%a`      | `Sun`                    |
| `%A`      | `Sunday`                 |
| `%w`      | `0`..`6` _(Sunday is 0)_ |
| `%y`      | `13`                     |
| `%Y`      | `2013`                   |
| `%b`      | `Jan`                    |
| `%B`      | `January`                |
| `%m`      | `01`..`12`               |
| `%d`      | `01`..`31`               |
| `%e`      | `1`..`31`                |

### Time

| Directive | Example      |
| --------- | ------------ |
| `%l`      | `1`          |
| `%H`      | `00`..`23`   |
| `%I`      | `01`..`12`   |
| `%M`      | `00`..`59`   |
| `%S`      | `00`..`60`   |
| `%p`      | `AM`         |
| `%Z`      | `+08`        |
| `%j`      | `001`..`366` |
| `%%`      | `%`          |
