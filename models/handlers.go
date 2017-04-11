package models

import (
	"encoding/json"
	"go-proxycheck/config"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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

//FilterCountry get
func FilterCountry(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, _ := strconv.Atoi(p.ByName("id"))

		if r.Method != get {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := FilterCountryReq(env.DB, id)
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

//FilterProxy get
func FilterProxy(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, _ := strconv.Atoi(p.ByName("id"))

		if r.Method != get {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := FilterProxyReq(env.DB, id)
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

//UpdateProxyStatus post json
func UpdateProxyStatus(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, _ := strconv.Atoi(p.ByName("id"))

		if r.Method != "POST" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		err := UpdateStatus(env.DB, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
	}
}
