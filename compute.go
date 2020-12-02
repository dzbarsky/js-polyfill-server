package polyfill

import (
	"errors"
	"io"
	"strconv"

	"github.com/ua-parser/uap-go/uaparser"
)

func New() *Composer {
	return &Composer{
		uaParser: uaparser.NewFromSaved(),
	}
}

type Composer struct {
	uaParser *uaparser.Parser
}

func (c *Composer) parse(userAgent string) (browser Browser, majorVersion int, err error) {
	ua := c.uaParser.ParseUserAgent(userAgent)
	switch ua.Family {
	case "Firefox":
		browser = Firefox
	case "Chrome":
		browser = Chrome
	case "Safari":
		browser = Safari
	case "IE":
		browser = IE
	default:
		err = errors.New("Unrecognized browser")
		return
	}

	majorVersion, err = strconv.Atoi(ua.Major)
	return
}

func (c *Composer) PolyfillsByName(userAgent string, polyfills []string) (io.Reader, error) {
	structs := make([]*Polyfill, len(polyfills))
	for i, p := range polyfills {
		s, ok := polyfillsByName[p]
		if !ok {
			return nil, errors.New("Unknown polyfill " + p)
		}
		structs[i] = s
	}

	return c.Polyfills(userAgent, structs), nil
}

func includeWithDeps(
	p *Polyfill,
	taken map[*Polyfill]struct{},
	codes []io.Reader,
	needed func(p *Polyfill) bool,
) []io.Reader {
	if _, ok := taken[p]; ok {
		return codes
	}
	taken[p] = struct{}{}
	for _, d := range p.deps {
		codes = includeWithDeps(d, taken, codes, needed)
	}
	if needed(p) {
		codes = append(codes, p.Code())
	}
	return codes
}

func (c *Composer) Polyfills(userAgent string, polyfills []*Polyfill) io.Reader {
	alwaysPolyfill := false
	browser, version, err := c.parse(userAgent)
	if err != nil {
		alwaysPolyfill = true
	}

	var codes []io.Reader

	taken := map[*Polyfill]struct{}{}

	for _, poly := range polyfills {
		codes = includeWithDeps(poly, taken, codes,
			func(p *Polyfill) bool {
				return alwaysPolyfill || p.Needed(browser, version)
			})
	}

	return io.MultiReader(codes...)
}
