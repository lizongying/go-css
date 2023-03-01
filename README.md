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
  [css_test](./test/css_test.go)


```go
package main

import (
	"fmt"
	"github.com/lizongying/go-css/css"
)

func main() {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewXpathFromStr(html)
	fmt.Println(x)
}
```