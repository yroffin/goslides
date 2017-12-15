/**
 * this is the name of our package
 */
package main

import (
	// fmt has methods for formatted IO

	// the "net/http" library has methods for HTTP

	"net/http"
	// Gorilla router
	"github.com/gorilla/mux"
	// Apis
	"github.com/yroffin/goslides/apis"
	"github.com/yroffin/goslides/bean"
	"github.com/yroffin/goslides/manager"
)

// Rest()
func main() {
	// define all routes
	var r = mux.NewRouter()

	var m = manager.Manager{}
	m.Init()
	m.Register("router", &apis.Router{Bean: &bean.Bean{}})
	m.Register("slide", &apis.Slide{API: &apis.API{Router: r}})
	m.Boot()

	// After defining our server, we finally "listen and serve" on port 8080
	http.ListenAndServe(":8080", nil)
}
