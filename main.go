package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"github.com/rancher/rancher-public-dns/service"
	"flag"

)
var(
	debug           = flag.Bool("debug", false, "Debug")
)

func main() {
	SetEnv()

	log.Info("Starting Rancher Public DNS service")
	//fmt.Print("Starting Rancher Public DNS service")

	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":8089", router))
}

func SetEnv() {
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}
	
	textFormatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(textFormatter)
}