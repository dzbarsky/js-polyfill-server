package polyfill

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ua-parser/uap-go/uaparser"

	"github.com/lytic-health/js-polyfill-server/types"
	"github.com/lytic-health/js-polyfill-server/version"
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
	case "Firefox Mobile":
		browser = types.FirefoxMobile
	case "Chrome", "Chrome Mobile", "Chrome Mobile WebView":
		browser = types.Chrome
	case "Safari":
		browser = types.Safari
	case "Mobile Safari", "Chrome Mobile iOS", "Firefox iOS":
		browser = types.IOSSafari
	case "IE":
		browser = types.IE
	case "Edge":
		browser = types.Edge
	case "Edge Mobile":
		browser = types.EdgeMobile
	default:
		fmt.Println("Unrecognized browser family: ", ua.Family)
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
	if needed(p) {
		p.ForEachDep(func(d *types.Polyfill) {
			codes = includeWithDeps(d, taken, codes, needed)
		})
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

func name(p *types.Polyfill) string {
	for k, v := range polyfillsByName {
		if v == p {
			return k
		}
	}
	panic("unknown polyfill")
}
