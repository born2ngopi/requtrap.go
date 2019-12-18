package http

import (
	"fmt"
	"github.com/libidev/requtrap.go/cli/config"
	"log"
	"net/http"
	"strconv"
)

// Serve - run HTTP Server
func Serve(conf *config.Yaml) {
	var host = conf.Host
	var port = strconv.Itoa(conf.Port)
	var uri = host + ":" + port

	fmt.Printf("%s running on http://%s\n", conf.Name, uri)

	handler := &Handler{}
	for _, service := range conf.Services {
		handler.AddRoute(service)
	}
	handler.Cors = conf.Cors

	err := http.ListenAndServe(uri, handler)
	if err != nil {
		log.Fatal(err)
	}
}
