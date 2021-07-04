package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var portPtr *int
var badPercentPtr *float64
var timeoutPtr *int

func main() {
	portPtr = flag.Int("p", 8899, "Which port to listen on")
	badPercentPtr = flag.Float64("e", 50, "Percentage of the requests that should return a server error")
	timeoutPtr = flag.Int("t", 0, "Maximum amount of time to randomly stall (milliseconds)")
	flag.Parse()

	h := &handler{}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *portPtr),
		Handler: h,
	}
	log.Println("Starting Flaky Server on " + server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving request " + r.Method + " " + r.URL.String())

	if *timeoutPtr > 0 {
		time.Sleep(time.Duration(rand.Intn(*timeoutPtr)) * time.Millisecond)
	}

	if rand.Float64()*100 < *badPercentPtr {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Flaky Server had an error"))
		if err != nil {
			log.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Flaky Server succeeded"))
		if err != nil {
			log.Println(err)
		}
	}
}
