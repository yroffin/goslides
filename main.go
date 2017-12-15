/**
 * this is the name of our package
 */
package main

import (
	// fmt has methods for formatted IO

	// the "net/http" library has methods for HTTP
	"fmt"
	"net/http"
	// Gorilla router
	"github.com/gorilla/mux"
	// Apis
	"github.com/yroffin/goslides/apis"
	"github.com/yroffin/goslides/interfaces"
)

// Rest()
func main() {
	var m = interfaces.Manager{}
	m.Init()
	m.Register("router", apis.Router{})
	m.Register("slide", apis.Slide{})
	m.Boot()

	// define all routes
	var r = mux.NewRouter()

	// declare API slides
	slide := apis.Slide{StructAPI: &apis.StructAPI{Router: r}}
	slide.PostConstruct()
	fmt.Println("slide", slide)
	slide.Init()

	// handle now all requests
	http.Handle("/", r)
	// After defining our server, we finally "listen and serve" on port 8080
	http.ListenAndServe(":8080", nil)
}
