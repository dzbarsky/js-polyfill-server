package http

import (
	"io"
	"net/http"
	"strings"

	polyfill "github.com/lytic-health/js-polyfill-server"
)

func Handler() func(w http.ResponseWriter, r *http.Request) {
	c := polyfill.New()
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		out, err := c.PolyfillsByName(r.UserAgent(),
			strings.Split(r.Form.Get("features"), ","))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = io.Copy(w, out)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
