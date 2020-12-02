package polyfill

import (
	"io"
	"strings"
)

type Browser string

const (
	Firefox = Browser("Firefox")
	Chrome  = Browser("Chrome")
	Safari  = Browser("Safari")
	IE      = Browser("IE")
)

type LastVersionBeforeSupport struct {
	Firefox int
	Chrome  int
	Safari  int
}

type Polyfill struct {
	code     string
	deps     []*Polyfill
	versions LastVersionBeforeSupport
}

func (p Polyfill) Needed(browser Browser, majorVersion int) bool {
	v := p.versions
	switch browser {
	case Firefox:
		return v.Firefox == 0 || majorVersion <= v.Firefox
	case Chrome:
		return v.Chrome == 0 || majorVersion <= v.Chrome
	case Safari:
		return v.Safari == 0 || majorVersion <= v.Safari
	default:
		return true
	}
}

func (p Polyfill) Code() io.Reader {
	return strings.NewReader(p.code)
}
