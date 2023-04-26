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
