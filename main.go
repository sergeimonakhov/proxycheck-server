package main

import (
	"flag"
	"fmt"
	"go-proxycheck/config"
	"go-proxycheck/models"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/parnurzeal/gorequest"
)

type getWithproxy struct {
	proxy string
	url   string
	env   *config.Env
}

func (g *getWithproxy) getproxy() {
	httpProxy := fmt.Sprintf("http://%s", g.proxy)
	str := strings.Split(g.proxy, ":")
	ip := str[0]
	bks := models.ExistIP(g.env.DB, ip)
	//
	if bks == false {
		request := gorequest.New().Proxy(httpProxy).Timeout(2 * time.Second)
		timeStart := time.Now()
		resp, _, err := request.Get(g.url).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError, http.StatusRequestTimeout).End()
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("GOOD: ", g.proxy)
			country := ipToCountry(ip)
			respone := time.Since(timeStart)
			//add to bd
			models.AddToBase(g.env.DB, g.proxy, country, respone)
			//
		} else {
			fmt.Println("BAD: ", g.proxy)
		}
	}
}

func ipToCountry(ip string) string {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		fmt.Printf("Could not open GeoIP database\n")
		os.Exit(1)
	}
	defer db.Close()
	country, _ := db.Country(net.ParseIP(ip))
	return country.Country.IsoCode
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	var (
		url    = flag.String("url", "https://m.vk.com", "")
		fileIn = flag.String("in", "proxylist.txt", "full path to proxy file")
		server = flag.Bool("server", false, "run httprest api on 3000 port")
		treds  = flag.Int("treds", 50, "number of treds")
	)

	flag.Parse()

	db, err := config.NewDB("postgres://proxy:proxy@localhost/proxy?sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	env := &config.Env{DB: db}

	if *server == true {
		router := httprouter.New()
		router.GET("/proxys", models.ProxyIndex(env))
		router.GET("/countrys", models.CountryIndex(env))
		http.ListenAndServe(":3000", router)
	} else {
		content, _ := ioutil.ReadFile(*fileIn)
		proxys := strings.Split(string(content), "\n")

		workers := *treds

		wg := new(sync.WaitGroup)
		in := make(chan string, 2*workers)

		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for proxy := range in {
					gp := getWithproxy{
						proxy: proxy,
						url:   *url,
						env:   env,
					}
					gp.getproxy()
				}
			}()
		}

		for _, proxy := range proxys {
			if proxy != "" {
				in <- proxy
			}
		}
		close(in)
		wg.Wait()
	}
}
