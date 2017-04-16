package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/D1abloRUS/proxycheck-server/config"
	"github.com/D1abloRUS/proxycheck-server/models"

	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
)

type postgres struct {
	User     string
	Password string
	Host     string
	DBname   string
	Port     int `default:"5432"`
}

func main() {
	var p postgres

	err := envconfig.Process("postgres", &p)
	if err != nil {
		fmt.Println(err.Error())
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.User, p.Password, p.DBname)

	db, err := config.NewDB(psqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	env := &config.Env{DB: db}
	fmt.Println("proxycheck-server started normaly")

	router := httprouter.New()
	router.GET("/api/v1/proxy", models.AllProxy(env))
	router.GET("/api/v1/country", models.AllCountry(env))
	router.GET("/api/v1/country/:id", models.FilterCountry(env))
	router.GET("/api/v1/proxy/:id", models.FilterProxy(env))
	router.POST("/api/v1/proxy/:id", models.UpdateProxyStatus(env))
	router.POST("/api/v1/addproxy", models.AddProxy(env))
	http.ListenAndServe(":3000", router)
}
