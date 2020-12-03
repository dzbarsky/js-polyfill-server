// Note: Benchmark written as external package to ensure code is usable
package polyfill_test

import (
	"io/ioutil"
	"testing"

	polyfill "js-polyfill-server"
	"js-polyfill-server/types"
)

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
