# changecase

changecase is a port of npm package [change-case](https://github.com/blakeembrey/change-case).
It provides helpful various text conversion functions.

# Usage

```go
import "github.com/ku/go-change-case"
changecase.Camel("go-change-case")  // returns "goChangeCase"
```

It contains following methods as [original change-case](https://github.com/blakeembrey/change-case) does.

- [`camel`](#camel)
     - `goChangeCase`
- [`snake`](#snake)
    - `go_change_case`
- [`pascal`](#pascal)
    - `GoChangeCase`
- [`param`](#param)
    - `go-change-case`
- [`dot`](#dot)
     - `go.change.case`
- [`constant`](#constant)
     - `GO_CHANGE_CASE`
- [`header`](#header)
     - `Go-Change-Case`
- [`path`](#path)
    - `go/change/case`
- [`sentence`](#sentence)
    - `Go change case`
- [`title`](#title)
    - `Go Change Case`
- [`no`](#no)
    - `go change case`
- [`lower`](#lower)
    - alias for `strings.ToLower()`
- [`upper`](#upper)
    - alias for `strings.ToUpper()`
- [`lcFirst`](#lcFirst)
    - `go-change-case`
- [`ucFirst`](#ucFirst)
    - `Go-change-case`
- [`swap`](#swap)
    - `GO-CHANGE-CASE`
- [`isLower`](#isLower)
    - Returns true if all the letters in given string are in lower case.
- [`isUpper`](#isUpper)
    - Returns true if all the letters in given string are in upper case.
