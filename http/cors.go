package http

import (
	"github.com/libidev/requtrap.go/cli/config"
	"net/http"
	"strings"
)

// EnableCors - Function to add CORS headers into HTTP response header
func EnableCors(w *http.ResponseWriter, conf config.Cors) {
	origins := strings.Join(conf.Origins, ",")
	methods := strings.Join(conf.Methods, ",")

	(*w).Header().Set("Access-Control-Allow-Origin", origins)
	(*w).Header().Set("Access-Control-Allow-Methods", methods)
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
