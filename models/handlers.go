package models

import (
	"fmt"
	"go-proxycheck/config"
	"net/http"
)

//ProxyIndex get all proxy
func ProxyIndex(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := AllProxy(env.DB)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
		}
	})
}
