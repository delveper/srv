package core

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/delveper/srv/cfg"
)

// ApiHandler meant to be handler
// and will keep all necessary configuration
type ApiHandler struct{ *cfg.Option }

// ServeHTTP will handle main logic
// getting random number from given range
func (hdl ApiHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// TODO: Is it a good place to use context?
	max := req.Context().Value("max").(int)
	min := req.Context().Value("min").(int)

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max-min) + min

	if _, err := fmt.Fprintf(rw, "<h1>The magic number is: %v</h1>", num); err != nil {
		http.Error(rw, "error writing response", http.StatusInternalServerError)
	}
}

// DefaultHandler will return 404 error if user
// will try to reach unpredictable path.
// In case of home page Welcome message will render
func DefaultHandler(hdl http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(rw, req)
			return
		}

		// it is not absolutely necessary,
		// just drilling code
		hdl = http.RedirectHandler("/index", http.StatusTemporaryRedirect)
		hdl.ServeHTTP(rw, req)
	})
}

// IndexHandler serve static file that contains
// basic html code to build the query for main logic
func IndexHandler(hdl http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, "./web/static/index.html")
	})
}
