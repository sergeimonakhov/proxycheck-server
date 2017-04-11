package main

import (
	"fmt"
	"go-proxycheck/config"
	"go-proxycheck/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := config.NewDB("postgres://proxy:proxy@localhost/proxy?sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
	}
	env := &config.Env{DB: db}

	router := httprouter.New()
	router.GET("/api/v1/proxy", models.AllProxy(env))
	router.GET("/api/v1/country", models.AllCountry(env))
	router.GET("/api/v1/country/:id", models.FilterCountry(env))
	router.GET("/api/v1/proxy/:id", models.FilterProxy(env))
	router.POST("/api/v1/proxy/:id", models.UpdateProxyStatus(env))
	router.POST("/api/v1/addproxy", models.AddProxy(env))
	http.ListenAndServe(":3000", router)
}
