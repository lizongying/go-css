# go-css

简单的css选择器 simple css selector

[document](https://pkg.go.dev/github.com/lizongying/go-css)

## Install

```
go get github.com/lizongying/go-css
```

## Usage

* get attr `@attr`
* For more usage, please refer to the test
  [selector_test](./selector/selector_test.go)


```go
package main

import (
	"fmt"
	"github.com/lizongying/go-css/selector"
)

func main() {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := selector.NewSelectorFromStr(html)
	fmt.Println(x)
}
```