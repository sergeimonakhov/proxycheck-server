package models

import (
	"encoding/json"
	"go-proxycheck/config"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*func getID(w http.ResponseWriter, ps httprouter.Params) (string, bool) {
	id := ps.ByName("id")
	if err != nil {
		w.WriteHeader(400)
		return "0", false
	}
	return id, true
}*/
const (
	get = "GET"
)

//AllProxy get
func AllProxy(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		if r.Method != get {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		bks, err := AllProxyReq(env.DB)
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

//AllCountry get
func AllCountry(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		if r.Method != get {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := AllCountryReq(env.DB)
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

/*//IDindex get
func IDindex(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		if r.Method != get {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := InfoID(env.DB, id)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err = json.NewEncoder(w).Encode(bks); err != nil {
			w.WriteHeader(500)
		}
	}
}*/
