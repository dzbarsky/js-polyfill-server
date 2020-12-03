package polyfill

import (
	"errors"
	"io"
	"strings"

	"github.com/ua-parser/uap-go/uaparser"

	"js-polyfill-server/types"
	"js-polyfill-server/version"
)

func New() *Composer {
	return &Composer{
		uaParser: uaparser.NewFromSaved(),
	}
}

type Composer struct {
	uaParser *uaparser.Parser
}

func (c *Composer) parse(userAgent string) (browser types.Browser, v version.Version, err error) {
	// TODO(can we stick to just pulling out major/minor/patch and avoid this slow parse?)
	ua := c.uaParser.ParseUserAgent(userAgent)
	switch ua.Family {
	case "Firefox":
		browser = types.Firefox
	case "Chrome":
		browser = types.Chrome
	case "Safari":
		browser = types.Safari
	case "IE":
		browser = types.IE
	default:
		err = errors.New("Unrecognized browser")
		return
	}

	v, err = version.Parse(strings.Join([]string{ua.Major, ua.Minor}, "."))
	return
}

func (c *Composer) PolyfillsByName(userAgent string, polyfills []string) (io.Reader, error) {
	structs := make([]*types.Polyfill, len(polyfills))
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
	p *types.Polyfill,
	taken map[*types.Polyfill]struct{},
	codes []io.Reader,
	needed func(p *types.Polyfill) bool,
) []io.Reader {
	if _, ok := taken[p]; ok {
		return codes
	}
	taken[p] = struct{}{}
	p.ForEachDep(func(d *types.Polyfill) {
		codes = includeWithDeps(d, taken, codes, needed)
	})
	if needed(p) {
		codes = append(codes, p.Code())
	}
	return codes
}

func (c *Composer) Polyfills(userAgent string, polyfills []*types.Polyfill) io.Reader {
	// TODO(dave): make this configurable
	alwaysPolyfill := false
	browser, version, err := c.parse(userAgent)
	if err != nil {
		alwaysPolyfill = true
	}

	var codes []io.Reader

	taken := map[*types.Polyfill]struct{}{}

	for _, poly := range polyfills {
		codes = includeWithDeps(poly, taken, codes,
			func(p *types.Polyfill) bool {
				return alwaysPolyfill || p.Needed(browser, version)
			})
	}

	return io.MultiReader(codes...)
}
