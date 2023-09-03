# go-css

golang实现的简单css选择器，使用更简单。simple css selector

[go-css](https://github.com/lizongying/go-css)

[document](https://pkg.go.dev/github.com/lizongying/go-css)

[english](./README.md)

## 安装

```
go get -u github.com/lizongying/go-css@latest
```

## 用法

* 获取属性 `@attr`
* 有关更多用法，请参考测试文件。
  [selector_test](./css/selector_test.go)

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-css/css"
)

func main() {
	html := `<html class="abc">....<div class="def">....</div><div class="gkl">123</div></html>`
	x, _ := css.NewSelectorFromStr(html)

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