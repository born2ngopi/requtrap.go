package http

import (
	"bytes"
	"fmt"
	"github.com/libidev/requtrap.go/cli/config"
	"io/ioutil"
	"net/http"
	"time"
)

// Handler - main HTTP handler
type Handler struct {
	Routes []config.Service
	Cors   config.Cors
}

// GetRequestMethod - to get request method, eg. GET, POST, PUT, etc
func (h Handler) GetRequestMethod(r *http.Request) string {
	return r.Method
}

// AddRoute - to add service routes into route list
func (h *Handler) AddRoute(service config.Service) {
	h.Routes = append(h.Routes, service)
}

// ServeHTTP - entry point for HTTP handler
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Cors.Enable {
		EnableCors(&w, h.Cors)
	}

	if r.Method == http.MethodOptions {
		return
	}

	if r.URL.Path != "favicon.ico" {
		for _, service := range h.Routes {
			path := r.URL.Path
			if path == service.Path {
				h.Gateway(w, r)(service)
			}
		}
	}
}

// Gateway - frowarding client request to all services
func (h Handler) Gateway(w http.ResponseWriter, r *http.Request) func(config.Service) {
	var reqbody []byte
	var err error

	if h.GetRequestMethod(r) == http.MethodPost {
		reqbody, err = ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return func(service config.Service) {
		url := service.Upstream + service.Path
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest(h.GetRequestMethod(r), url, bytes.NewBuffer(reqbody))
		req.Header.Set("Content-type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("")
		fmt.Println("Status :", resp.Status)
		fmt.Println("Header :")
		for k, v := range resp.Header {
			fmt.Println("  ", k+":", v[0])
		}
		fmt.Println("Body   :", string(body))

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
