package main

import (
	"flag"
	"fmt"
	"github.com/julienmarie/bandit-server/assets"
	"github.com/julienmarie/bandit-server/handlers"
	"github.com/julienmarie/bandit-server/repository"
	"github.com/julienmarie/bandit-server/strategies"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var port int = 3000

var logger = log.New(os.Stdout, "bandit-server: ", log.Ldate|log.Ltime)

func init() {
	flag.IntVar(&port, "port", 3000, "http port")
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	var repo repository.Repository

	repo = repository.NewMemory()

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		var favicon, _ = assets.Asset("favicon.ico")
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(favicon)
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		var robots, _ = assets.Asset("robots.txt")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(robots)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "ok")
	})

	http.HandleFunc("/ucb1", handlers.NewHttpHandler(strategies.NewUCB1(), repo))

	logger.Print("Listening on port ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		logger.Fatal(err)
	}
}
