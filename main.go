package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	flagPort = flag.String("port", "5000", "http service address")
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		// Use flag if no environment variable
		port = *flagPort
	}
	return ":" + port, nil
}

func init() {
	flag.Parse()
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		enc := json.NewEncoder(w)
		ret := struct {
			Question string `json:"question"`
		}{
			Question: getQuestion(),
		}
		enc.Encode(ret)
	})

	log.Println("Service listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
