package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"log"
	"net/http"
	"strings"
)

var (
	bind   = "127.0.0.1:8080"
	header = "X-Real-IP"
)

func main() {
	flag.StringVar(&bind, "b", "127.0.0.1:8080", "Binding IP Address for the HTTP Server")
	flag.StringVar(&header, "h", "X-Real-IP", "What Header will be used to retrieve the Clients IP Address")
	flag.Parse()
	log.Printf("Starting Server and listening on %v, retrieving IP's from Header %v", bind, header)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	var resp []byte
	var err error
	switch format {
	case "json":
		remoteAddr := strings.Split(r.Header.Get(header), ",")[0]
		ip := IP{IP: remoteAddr}
		resp, err = json.Marshal(ip)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		break
	case "xml":
		remoteAddr := strings.Split(r.Header.Get(header), ",")[0]
		ip := IP{IP: remoteAddr}
		resp, err = xml.Marshal(ip)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/xml")
		break
	default:
		remoteAddr := strings.Split(r.Header.Get(header), ",")[0]
		err := r.Body.Close()
		if err != nil {
			panic(err)
		}
		resp = []byte(remoteAddr)
		w.Header().Set("Content-Type", "text/plain")
		break
	}
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Encountered Error: ")
		log.Println(err)
	}
}

type IP struct {
	IP string `json:"ip"`
}
