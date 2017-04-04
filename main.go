package main

import (
	"go-proxycheck/config"
	"go-proxycheck/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := config.NewDB("postgres://proxy:proxy@localhost/proxy?sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	env := &config.Env{DB: db}

	router := httprouter.New()
	router.GET("/proxys", models.ProxyIndex(env))
	http.ListenAndServe(":3000", router)
}
