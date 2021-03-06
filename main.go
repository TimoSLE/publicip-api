package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"log"
	"net/http"
	"strings"
)

//GitVersion is the current used git version
var GitVersion = ""

//GitBranch represents which Branch is currently running
var GitBranch = ""

var (
	bind   = "127.0.0.1:8080"
	header = "X-Real-IP"
)

func main() {
	log.Printf("PublicIP-API Version %s, Git Branch %s", GitVersion, GitBranch)
	flag.StringVar(&bind, "b", "127.0.0.1:8080", "Binding IP Address for the HTTP Server")
	flag.StringVar(&header, "h", "X-Real-IP", "What Header will be used to retrieve the Clients IP Address")
	flag.Parse()
	log.Printf("Starting Server and listening on %v, retrieving IP's from Header %v", bind, header)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Printf("Encountered error %v", err)
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
			log.Printf("Encountered error %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		break
	case "xml":
		remoteAddr := strings.Split(r.Header.Get(header), ",")[0]
		ip := IP{IP: remoteAddr}
		resp, err = xml.Marshal(ip)
		if err != nil {
			log.Printf("Encountered error %v", err)
		}
		w.Header().Set("Content-Type", "application/xml")
		break
	default:
		remoteAddr := strings.Split(r.Header.Get(header), ",")[0]
		err := r.Body.Close()
		if err != nil {
			log.Printf("Encountered error %v", err)
		}
		resp = []byte(remoteAddr)
		w.Header().Set("Content-Type", "text/plain")
		break
	}
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Encountered error %v", err)
	}
}

//IP represents structure for marshalling the IP into xml and json
type IP struct {
	IP string `json:"ip"`
}
