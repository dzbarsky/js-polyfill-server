// Note: Benchmark written as external package to ensure code is usable
package polyfill_test

import (
	"io/ioutil"
	"testing"

	polyfill "github.com/lytic-health/js-polyfill-server"
	"github.com/lytic-health/js-polyfill-server/types"
)

var polyfills = []*types.Polyfill{
	polyfill.Array_prototype_forEach,
	polyfill.Array_prototype_includes,
	polyfill.Element_prototype_remove,
	polyfill.Fetch,
	polyfill.Object_entries,
	polyfill.Object_keys,
	polyfill.Object_values,
	polyfill.Promise,
	polyfill.String_prototype_repeat,
	polyfill.String_prototype_startsWith,
}

func TestMobileChrome(t *testing.T) {
	c := polyfill.New()
	reader := c.Polyfills(
		"Mozilla/5.0 (Linux; Android 11; Pixel 2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Mobile Safari/537.36",
		polyfills,
	)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}
	if len(data) != 0 {
		t.Error("Should have no polyfills")
	}
}

func TestMobileFirefox(t *testing.T) {
	c := polyfill.New()
	reader := c.Polyfills(
		"Mozilla/5.0 (Android 11; Mobile; rv:84.0) Gecko/84.0 Firefox/84.0",
		polyfills,
	)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("Should have polyfills for Element.prototype.remove")
	}
}

var r []byte

func BenchmarkPolyfills(b *testing.B) {
	c := polyfill.New()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := c.Polyfills(
			"Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)",
			[]*types.Polyfill{
				polyfill.Array_prototype_forEach,
				polyfill.Array_prototype_includes,
				polyfill.Element_prototype_remove,
				polyfill.Fetch,
				polyfill.Object_entries,
				polyfill.Object_keys,
				polyfill.Object_values,
				polyfill.Promise,
				polyfill.String_prototype_repeat,
				polyfill.String_prototype_startsWith,
			},
		)
		r, _ = ioutil.ReadAll(reader)
	}
}
