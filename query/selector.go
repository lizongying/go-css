package query

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"os"
	"strconv"
	"strings"
)

type Selector struct {
	node *goquery.Selection
}

// GetNode get node
func (s *Selector) GetNode() (node *goquery.Selection) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	node = s.node
	return
}

func (s *Selector) Remove(css string) (node *goquery.Selection) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	s.node.RemoveFiltered(css)
	return s.node
}

// FindNodeMany find nodes
func (s *Selector) FindNodeMany(css string) (selectors []*Selector) {
	if s == nil {
		return
	}
	if s.node == nil {
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
		ns = s.node
	} else {
		ns = s.node.Find(css)
	}
	for i := range ns.Nodes {
		n := ns.Eq(i)
		if index > -1 {
			if _, exists := n.Attr(attr); !exists {
				continue
			}
		}
		selectors = append(selectors, &Selector{
			node: n,
		})
	}
	return
}

// FindNodeOne find node or nil
func (s *Selector) FindNodeOne(css string) (selector *Selector) {
	if s == nil {
		return
	}
	if s.node == nil {
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
		for i := range s.node.Nodes {
			if _, exists := s.node.Eq(i).Attr(attr); exists {
				n = s.node.Eq(i)
				break
			}
		}
	} else {
		n = s.node.Find(css).First()
	}
	if n == nil {
		return
	}
	selector = &Selector{
		node: n,
	}
	return
}

// FindNodeOneOr find node
func (s *Selector) FindNodeOneOr(css string) (selector *Selector) {
	if s == nil {
		selector = &Selector{}
		return
	}
	if s.node == nil {
		selector = &Selector{}
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
		for i := range s.node.Nodes {
			if _, exists := s.node.Eq(i).Attr(attr); exists {
				n = s.node.Eq(i)
				break
			}
		}
	} else {
		n = s.node.Find(css).First()
	}
	if n == nil {
		selector = &Selector{}
		return
	}
	selector = &Selector{
		node: n,
	}
	return
}

// FindStrMany find a string list
func (s *Selector) FindStrMany(css string) (list []string) {
	if s == nil {
		return
	}
	if s.node == nil {
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
		ns = s.node
	} else {
		ns = s.node.Find(css)
	}
	for i := range ns.Nodes {
		n := ns.Eq(i)
		if index > -1 {
			if _, exists := n.Attr(attr); !exists {
				continue
			}
		}
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
func (s *Selector) FindStrOne(css string) (str string) {
	if s == nil {
		return
	}
	if s.node == nil {
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
		for i := range s.node.Nodes {
			if _, exists := s.node.Eq(i).Attr(attr); exists {
				n = s.node.Eq(i)
				break
			}
		}
	} else {
		n = s.node.Find(css).First()
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
func (s *Selector) FindStrOneOr(css string, or string) (str string) {
	if s == nil {
		str = or
		return
	}
	if s.node == nil {
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
		for i := range s.node.Nodes {
			if _, exists := s.node.Eq(i).Attr(attr); exists {
				n = s.node.Eq(i)
				break
			}
		}
	} else {
		n = s.node.Find(css).First()
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
func (s *Selector) FindIntMany(css string) (list []int) {
	if s == nil {
		return
	}
	if s.node == nil {
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
		ns = s.node
	} else {
		ns = s.node.Find(css)
	}
	for i := range ns.Nodes {
		n := ns.Eq(i)
		if index > -1 {
			if _, exists := n.Attr(attr); !exists {
				continue
			}
		}
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
func (s *Selector) FindIntOne(css string) (i int) {
	if s == nil {
		return
	}
	if s.node == nil {
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
		for ii := range s.node.Nodes {
			if _, exists := s.node.Eq(ii).Attr(attr); exists {
				n = s.node.Eq(ii)
				break
			}
		}
	} else {
		n = s.node.Find(css).First()
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

// FindIntOneOr find one int, will return a default int if you find nothing
func (s *Selector) FindIntOneOr(css string, or int) (i int) {
	if s == nil {
		i = or
		return
	}
	if s.node == nil {
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
		for ii := range s.node.Nodes {
			if _, exists := s.node.Eq(ii).Attr(attr); exists {
				n = s.node.Eq(ii)
				break
			}
		}
	} else {
		n = s.node.Find(css).First()
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

// NewSelectorFromStr selector from string
func NewSelectorFromStr(s string) (selector *Selector, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return
	}
	selector = &Selector{
		node: doc.Selection,
	}
	return
}

// NewSelectorFromBytes selector from bytes
func NewSelectorFromBytes(b []byte) (selector *Selector, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return
	}
	selector = &Selector{
		node: doc.Selection,
	}
	return
}

// NewSelectorFromReader selector from reader
func NewSelectorFromReader(i io.Reader) (selector *Selector, err error) {
	doc, err := goquery.NewDocumentFromReader(i)
	if err != nil {
		return
	}

	selector = &Selector{
		node: doc.Selection,
	}
	return
}

// NewSelectorFromFile selector from file
func NewSelectorFromFile(f string) (selector *Selector, err error) {
	file, err := os.Open(f)
	if err != nil {
		return
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return
	}
	selector = &Selector{
		node: doc.Selection,
	}
	return
}

// OutHtml return html
func (s *Selector) OutHtml(self bool) (str string) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	if self {
		str, _ = s.node.Parent().Html()
	} else {
		str, _ = s.node.Html()
	}
	return
}
