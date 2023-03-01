package test

import (
	"github.com/lizongying/go-css/css"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

// go test -v test/css_test.go

// TestCssFromStr go test -v test/css_test.go -run TestCssFromStr
func TestCssFromStr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewCssFromStr(html)

	assert.Equal(t, x.GetNode().Get(0).FirstChild.Data, "html")
}

// TestCssFromReader go test -v test/css_test.go -run TestCssFromReader
func TestCssFromReader(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewCssFromReader(strings.NewReader(html))

	assert.Equal(t, x.GetNode().Get(0).FirstChild.Data, "html")
}

// TestCssFromFile go test -v test/css_test.go -run TestCssFromFile
func TestCssFromFile(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	file, _ := os.CreateTemp(os.TempDir(), "")
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Log(err)
		}
	}(file.Name())
	_, _ = file.WriteString(html)
	x, _ := css.NewCssFromFile(file.Name())

	assert.Equal(t, x.GetNode().Get(0).FirstChild.Data, "html")
}

// TestCssFindStrMany go test -v test/css_test.go -run TestCssFindStrMany
func TestCssFindStrMany(t *testing.T) {
	html := []byte(`<html class="123">....<div class="789">....</div><div class="456">....</div></html>`)
	x, _ := css.NewCssFromBytes(html)

	li := x.FindStrMany(`div@class`)
	t.Log(li)
	assert.Equal(t, li, []string{"789", "456"})

	li = x.FindStrMany(`div@classs`)
	t.Log(li)
	assert.Equal(t, li, []string{"", ""})
}

// TestCssFindStrOne go test -v test/css_test.go -run TestCssFindStrOne
func TestCssFindStrOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewCssFromStr(html)

	str := x.FindStrOne(`div@class`)
	t.Log(str)
	assert.Equal(t, str, "789")

	str = x.FindStrOne(`div@classs`)
	t.Log(str)
	assert.Equal(t, str, "")
}

// TestCssFindStrOneOr go test -v test/css_test.go -run TestCssFindStrOneOr
func TestCssFindStrOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewCssFromStr(html)

	str := x.FindStrOneOr(`div@class`, "999")
	t.Log(str)
	assert.Equal(t, str, "789")

	str = x.FindStrOneOr(`div@classs`, "999")
	t.Log(str)
	assert.Equal(t, str, "999")
}

// TestCssFindIntMany go test -v test/css_test.go -run TestCssFindIntMany
func TestCssFindIntMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewCssFromStr(html)

	li := x.FindIntMany(`div@class`)
	t.Log(li)
	assert.Equal(t, li, []int{789, 456})

	li = x.FindIntMany(`div@classs`)
	t.Log(li)
	assert.Equal(t, li, []int{0, 0})
}

// TestCssFindIntOne go test -v test/css_test.go -run TestCssFindIntOne
func TestCssFindIntOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewCssFromStr(html)

	i := x.FindIntOne(`div@class`)
	t.Log(i)
	assert.Equal(t, i, 789)

	i = x.FindIntOne(`div@classs`)
	t.Log(i)
	assert.Equal(t, i, 0)
}

// TestCssFindIntOneOr go test -v test/css_test.go -run TestCssFindIntOneOr
func TestCssFindIntOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := css.NewCssFromStr(html)

	i := x.FindIntOneOr(`div@class`, 999)
	t.Log(i)
	assert.Equal(t, i, 789)

	i = x.FindIntOneOr(`div@classs`, 999)
	t.Log(i)
	assert.Equal(t, i, 999)
}

// TestCssFindNodeManyFindStrMany go test -v test/css_test.go -run TestCssFindNodeManyFindStrMany
func TestCssFindNodeManyFindStrMany(t *testing.T) {
	html := `<div class="abc">....<div class="789">....</div><div class="456">....</div></div><div class="abc">....<div class="789">....</div><div class="jqk">....</div></div>`
	x, _ := css.NewCssFromStr(html)

	for k, i := range x.FindNodeMany(`div.abc`) {
		li := i.FindStrMany(`div@class`)
		if k == 0 {
			assert.Equal(t, li, []string{"789", "456"})
		}
		if k == 1 {
			assert.Equal(t, li, []string{"789", "jqk"})
		}
	}
}

// TestCssFindNodeOneFindStrMany go test -v test/css_test.go -run TestCssFindNodeOneFindStrMany
func TestCssFindNodeOneFindStrMany(t *testing.T) {
	html := `<div class="123">....<div class="789">....</div><div class="456">....</div></div><div class="123">....<div class="789">....</div><div class="456">....</div></div>`
	x, _ := css.NewCssFromStr(html)

	li := x.FindNodeOne(`div[class="123"]`).FindStrMany(`div@class`)
	t.Log(li)
	assert.Equal(t, li, []string{"789", "456"})

	li = x.FindNodeOneOr(`div[class="1234"]`).FindStrMany(`div@class`)
	t.Log(li)
	assert.Equal(t, li, []string(nil))
}

// TestCssOutHtml go test -v test/css_test.go -run TestCssOutHtml
func TestCssOutHtml(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="jqk">...</div></html>`
	x, _ := css.NewCssFromStr(html)

	s := x.FindNodeOne(`div`).OutHtml(true)
	t.Log(s)

	assert.Equal(t, s, "....")

	s = x.FindNodeOne(`div.jqk`).OutHtml(false)
	t.Log(s)

	assert.Equal(t, s, "...")
}
