package css

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"os"
	"strconv"
	"strings"
)

type Css struct {
	node *goquery.Selection
}

// GetNode get node
func (x *Css) GetNode() (node *goquery.Selection) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	node = x.node
	return
}

// FindNodeMany find nodes
func (x *Css) FindNodeMany(str string) (xs []*Css) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	ns := x.node.Find(str)
	for i := range ns.Nodes {
		xs = append(xs, &Css{
			node: ns.Eq(i),
		})
	}
	return
}

// FindNodeOne find node
func (x *Css) FindNodeOne(str string) (x1 *Css) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	n := x.node.Find(str).First()
	if n == nil {
		return
	}
	x1 = &Css{
		node: n,
	}
	return
}

// FindNodeOneOr find node
func (x *Css) FindNodeOneOr(str string) (x1 *Css) {
	if x == nil {
		x1 = &Css{}
		return
	}
	if x.node == nil {
		x1 = &Css{}
		return
	}
	n := x.node.Find(str).First()
	if n == nil {
		x1 = &Css{}
		return
	}
	x1 = &Css{
		node: n,
	}
	return
}

// FindStrMany find a string list
func (x *Css) FindStrMany(css string) (list []string) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	index := strings.LastIndex(css, "@")
	attr := ""
	if index > -1 {
		attr = css[index+1:]
		css = css[:index]
	}
	var ns *goquery.Selection
	if index == 0 {
		ns = x.node.Find(fmt.Sprintf(`[%s]`, attr))
	} else {
		ns = x.node.Find(css[:index])
	}
	for i := range ns.Nodes {
		n := ns.Eq(i)
		str := ""
		if attr != "" {
			str, _ = n.Attr(attr)
		} else {
			str = n.Text()
		}
		str = strings.TrimSpace(str)
		list = append(list, str)
	}
	return
}

// FindStrOne find a string
func (x *Css) FindStrOne(css string) (str string) {
	if x == nil {
		return
	}
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	index := strings.LastIndex(css, "@")
	attr := ""
	if index > -1 {
		attr = css[index+1:]
		css = css[:index]
	}
	var n *goquery.Selection
	if index == 0 {
		n = x.node.Find(fmt.Sprintf(`[%s]`, attr)).First()
	} else {
		n = x.node.Find(css).First()
	}
	if n == nil {
		return
	}
	if attr != "" {
		str, _ = n.Attr(attr)
	} else {
		str = n.Text()
	}
	str = strings.TrimSpace(str)
	return
}

// FindStrOneOr find a string, will return a default string if you find nothing
func (x *Css) FindStrOneOr(css string, or string) (str string) {
	if x == nil {
		str = or
		return
	}
	if x.node == nil {
		str = or
		return
	}
	index := strings.LastIndex(css, "@")
	attr := ""
	if index > -1 {
		attr = css[index+1:]
		css = css[:index]
	}
	var n *goquery.Selection
	if index == 0 {
		n = x.node.Find(fmt.Sprintf(`[%s]`, attr)).First()
	} else {
		n = x.node.Find(css).First()
	}
	if n == nil {
		str = or
		return
	}
	if attr != "" {
		str, _ = n.Attr(attr)
	} else {
		str = n.Text()
	}
	str = strings.TrimSpace(str)
	if str != "" {
		return
	}
	str = or
	return
}

// FindIntMany find int list
func (x *Css) FindIntMany(css string) (list []int) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	index := strings.LastIndex(css, "@")
	attr := ""
	if index > -1 {
		attr = css[index+1:]
		css = css[:index]
	}
	var ns *goquery.Selection
	if index == 0 {
		ns = x.node.Find(fmt.Sprintf(`[%s]`, attr))
	} else {
		ns = x.node.Find(css[:index])
	}
	for i := range ns.Nodes {
		n := ns.Eq(i)
		str := ""
		if attr != "" {
			str, _ = n.Attr(attr)
		} else {
			str = n.Text()
		}
		str = strings.TrimSpace(str)
		in, _ := strconv.Atoi(str)
		list = append(list, in)
	}
	return
}

// FindIntOne find int
func (x *Css) FindIntOne(css string) (i int) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	index := strings.LastIndex(css, "@")
	attr := ""
	if index > -1 {
		attr = css[index+1:]
		css = css[:index]
	}
	var n *goquery.Selection
	if index == 0 {
		n = x.node.Find(fmt.Sprintf(`[%s]`, attr)).First()
	} else {
		n = x.node.Find(css).First()
	}
	if n == nil {
		return
	}
	str := ""
	if attr != "" {
		str, _ = n.Attr(attr)
	} else {
		str = n.Text()
	}
	str = strings.TrimSpace(str)
	if str != "" {
		i, _ = strconv.Atoi(str)
		return
	}
	return
}

// FindIntOneOr find  int, will return a default int if you find nothing
func (x *Css) FindIntOneOr(css string, or int) (i int) {
	if x == nil {
		i = or
		return
	}
	if x.node == nil {
		i = or
		return
	}
	index := strings.LastIndex(css, "@")
	attr := ""
	if index > -1 {
		attr = css[index+1:]
		css = css[:index]
	}
	var n *goquery.Selection
	if index == 0 {
		n = x.node.Find(fmt.Sprintf(`[%s]`, attr)).First()
	} else {
		n = x.node.Find(css).First()
	}
	if n == nil {
		i = or
		return
	}
	str := ""
	if attr != "" {
		str, _ = n.Attr(attr)
	} else {
		str = n.Text()
	}
	str = strings.TrimSpace(str)
	if str != "" {
		i, _ = strconv.Atoi(str)
		return
	}
	i = or
	return
}

// NewCssFromStr css init
func NewCssFromStr(s string) (css *Css, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return
	}
	css = &Css{
		node: doc.Selection,
	}
	return
}

// NewCssFromBytes css init
func NewCssFromBytes(b []byte) (css *Css, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return
	}
	css = &Css{
		node: doc.Selection,
	}
	return
}

// NewCssFromReader css init
func NewCssFromReader(i io.Reader) (css *Css, err error) {
	doc, err := goquery.NewDocumentFromReader(i)
	if err != nil {
		return
	}

	css = &Css{
		node: doc.Selection,
	}
	return
}

// NewCssFromFile css init
func NewCssFromFile(f string) (css *Css, err error) {
	file, err := os.Open(f)
	if err != nil {
		return
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return
	}
	css = &Css{
		node: doc.Selection,
	}
	return
}

// OutHtml return html
func (x *Css) OutHtml(_ bool) (str string) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	str, _ = x.node.Html()
	return
}
