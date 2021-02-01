package types

import (
	"io"
	"strings"

	"github.com/lytic-health/js-polyfill-server/version"
)

type Browser int

const (
	Android Browser = iota
	BB
	Chrome
	Edge
	EdgeMobile
	Firefox
	IOSChrome
	IOSSafari
	IE
	IEMobile
	Opera
	OperaMini
	Safari
	FirefoxMobile
	SamsungMobile
)

type SupportMatrix struct {
	AndroidNeeded       version.Range
	BBNeeded            version.Range
	ChromeNeeded        version.Range
	EdgeNeeded          version.Range
	EdgeMobileNeeded    version.Range
	FirefoxNeeded       version.Range
	IOSChromeNeeded     version.Range
	IOSSafariNeeded     version.Range
	IENeeded            version.Range
	IEMobileNeeded      version.Range
	OperaNeeded         version.Range
	OperaMiniNeeded     version.Range
	SafariNeeded        version.Range
	FirefoxMobileNeeded version.Range
	SamsungMobileNeeded version.Range
}

type Polyfill struct {
	code    string
	deps    []*Polyfill
	support SupportMatrix
}

func NewPolyfill(code string, deps []*Polyfill, support SupportMatrix) *Polyfill {
	return &Polyfill{code, deps, support}
}

func (p Polyfill) Needed(browser Browser, version version.Version) bool {
	v := p.support
	switch browser {
	case Android:
		return v.AndroidNeeded(version)
	case BB:
		return v.BBNeeded(version)
	case Chrome:
		return v.ChromeNeeded(version)
	case Edge:
		return v.EdgeNeeded(version)
	case EdgeMobile:
		return v.EdgeMobileNeeded(version)
	case Firefox:
		return v.FirefoxNeeded(version)
	case IOSChrome:
		return v.IOSChromeNeeded(version)
	case IOSSafari:
		return v.IOSSafariNeeded(version)
	case IE:
		return v.IENeeded(version)
	case IEMobile:
		return v.IEMobileNeeded(version)
	case Opera:
		return v.OperaNeeded(version)
	case OperaMini:
		return v.OperaMiniNeeded(version)
	case Safari:
		return v.SafariNeeded(version)
	case FirefoxMobile:
		return v.FirefoxMobileNeeded(version)
	case SamsungMobile:
		return v.SamsungMobileNeeded(version)
	default:
		// TODO(configure unknown browser)
		return true
	}
}

func (p Polyfill) ForEachDep(fn func(dep *Polyfill)) {
	for _, d := range p.deps {
		fn(d)
	}
}

func (p Polyfill) Code() io.Reader {
	return strings.NewReader(p.code)
}
