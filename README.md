# go-query

golang实现的简单query选择器，使用更简单。simple query selector

[go-query](https://github.com/lizongying/go-query)
[document](https://pkg.go.dev/github.com/lizongying/go-query)

## Install

```
go get github.com/lizongying/go-query
```

## Usage

* get attr `@attr`
* For more usage, please refer to the test
  [selector_test](./query/selector_test.go)

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-query/query"
)

func main() {
	html := `<html class="abc">....<div class="def">....</div><div class="gkl">123</div></html>`
	x, _ := query.NewSelectorFromStr(html)

	s := x.FindStrOne(`.def`)
	//....
	fmt.Println(s)

	i := x.FindIntOneOr(`.gkl`, 111)
	//123
	fmt.Println(i)

	i = x.FindIntOneOr(`.mn`, 111)
	//111
	fmt.Println(i)

	sl := x.FindStrMany(`div`)
	//[.... 123]
	fmt.Println(sl)

	s = x.FindNodeOne(`.abc`).FindStrOne(`.gkl`)
	//123
	fmt.Println(s)
}

```