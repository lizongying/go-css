package selector

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

// go test -v selector/*.go

// TestSelectorFromStr go test -v selector/*.go -run TestSelectorFromStr
func TestSelectorFromStr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	assert.Equal(t, "html", x.GetNode().Get(0).FirstChild.Data)
}

// TestSelectorFromReader go test -v selector/*.go -run TestSelectorFromReader
func TestSelectorFromReader(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromReader(strings.NewReader(html))

	assert.Equal(t, "html", x.GetNode().Get(0).FirstChild.Data)
}

// TestSelectorFromFile go test -v selector/*.go -run TestSelectorFromFile
func TestSelectorFromFile(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	file, _ := os.CreateTemp(os.TempDir(), "")
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Log(err)
		}
	}(file.Name())
	_, _ = file.WriteString(html)
	x, _ := NewSelectorFromFile(file.Name())

	assert.Equal(t, "html", x.GetNode().Get(0).FirstChild.Data)
}

// TestSelectorFindStrMany go test -v selector/*.go -run TestSelectorFindStrMany
func TestSelectorFindStrMany(t *testing.T) {
	html := []byte(`<html class="123">....<div class="789">....</div><div class="456">....</div></html>`)
	x, _ := NewSelectorFromBytes(html)

	li := x.FindStrMany(`div@class`)
	t.Log(li)
	assert.Equal(t, []string{"789", "456"}, li)

	li = x.FindStrMany(`div@classs`)
	t.Log(li)
	assert.Equal(t, []string(nil), li)
}

// TestSelectorFindStrOne go test -v selector/*.go -run TestSelectorFindStrOne
func TestSelectorFindStrOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	str := x.FindStrOne(`div@class`)
	t.Log(str)
	assert.Equal(t, "789", str)

	str = x.FindStrOne(`div@classs`)
	t.Log(str)
	assert.Equal(t, "", str)
}

// TestSelectorFindStrOneOr go test -v selector/*.go -run TestSelectorFindStrOneOr
func TestSelectorFindStrOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	str := x.FindStrOneOr(`div@class`, "999")
	t.Log(str)
	assert.Equal(t, "789", str)

	str = x.FindStrOneOr(`div@classs`, "999")
	t.Log(str)
	assert.Equal(t, "999", str)
}

// TestSelectorFindIntMany go test -v selector/*.go -run TestSelectorFindIntMany
func TestSelectorFindIntMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	li := x.FindIntMany(`div@class`)
	t.Log(li)
	assert.Equal(t, []int{789, 456}, li)

	li = x.FindIntMany(`div@classs`)
	t.Log(li)
	assert.Equal(t, []int(nil), li)
}

// TestSelectorFindIntOne go test -v selector/*.go -run TestSelectorFindIntOne
func TestSelectorFindIntOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	i := x.FindIntOne(`div@class`)
	t.Log(i)
	assert.Equal(t, 789, i)

	i = x.FindIntOne(`div@classs`)
	t.Log(i)
	assert.Equal(t, 0, i)
}

// TestSelectorFindIntOneOr go test -v selector/*.go -run TestSelectorFindIntOneOr
func TestSelectorFindIntOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	i := x.FindIntOneOr(`div@class`, 999)
	t.Log(i)
	assert.Equal(t, 789, i)

	i = x.FindIntOneOr(`div@classs`, 999)
	t.Log(i)
	assert.Equal(t, 999, i)
}

// TestSelectorFindNodeManyFindStrMany go test -v selector/*.go -run TestSelectorFindNodeManyFindStrMany
func TestSelectorFindNodeManyFindStrMany(t *testing.T) {
	html := `<div class="abc">1<div class="789">2</div><div class="456">3</div></div><div class="abc">4<div classs="789">5</div><div classs="jqk">6</div></div>`
	x, _ := NewSelectorFromStr(html)

	//for k, i := range x.FindNodeMany(`div.abc div`) {
	//	li := i.FindStrMany(`@class`)
	//	if k == 0 {
	//		assert.Equal(t, []string{"789"}, li)
	//	}
	//	if k == 1 {
	//		assert.Equal(t, []string{"456"}, li)
	//	}
	//}
	for k, i := range x.FindNodeMany(`div.abc div`) {
		s := i.FindStrOne(`@classs`)
		if k == 0 {
			assert.Equal(t, "", s)
		}
		if k == 2 {
			t.Log("s", s, i.OutHtml(false))
			assert.Equal(t, "789", s)
		}
	}
}

// TestSelectorFindNodeOneFindStrMany go test -v selector/*.go -run TestSelectorFindNodeOneFindStrMany
func TestSelectorFindNodeOneFindStrMany(t *testing.T) {
	html := `<div class="123">....<div class="789">....</div><div class="456">....</div></div><div class="123">....<div class="789">....</div><div class="457">....</div></div>`
	x, _ := NewSelectorFromStr(html)

	li := x.FindNodeOne(`div[class="123"]`).FindStrMany(`div@class`)
	t.Log(li)
	assert.Equal(t, []string{"789", "456"}, li)

	li = x.FindNodeOneOr(`div[class="123"]`).FindStrMany(`@class`)
	t.Log(li)
	assert.Equal(t, []string{"123"}, li)

	li = x.FindNodeMany(`div[class="123"]`)[1].FindStrMany(`@class`)
	t.Log(li)
	assert.Equal(t, []string{"123"}, li)
}

// TestSelectorOutHtml go test -v selector/*.go -run TestSelectorOutHtml
func TestSelectorOutHtml(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="jqk">...</div></html>`
	x, _ := NewSelectorFromStr(html)

	s := x.FindNodeOne(`div`).OutHtml(false)
	t.Log(s)

	assert.Equal(t, "....", s)

	s = x.FindNodeOne(`div.jqk`).OutHtml(true)
	t.Log(s)

	assert.Equal(t, "....<div class=\"789\">....</div><div class=\"jqk\">...</div>", s)
}
