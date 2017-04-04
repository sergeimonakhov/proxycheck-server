package models

import (
	"encoding/json"
	"go-proxycheck/config"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Proxy table
type Proxy struct {
	id      int
	Proxy   string
	country string
	respone string
}

//ProxyIndex get all proxy
func ProxyIndex(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var str string

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		if len(r.URL.RawQuery) > 0 {
			str = r.URL.Query().Get("country")
			if str == "" {
				w.WriteHeader(400)
				return
			}
		}

		bks, err := CountryProxy(env.DB, str)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err = json.NewEncoder(w).Encode(bks); err != nil {
			w.WriteHeader(500)
		}
	}
}
