// Package apis for common interfaces
// MIT License
//
// Copyright (c) 2017 yroffin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package apis

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/yroffin/goslides/bean"
)

// API base class
type API struct {
	// members
	*bean.Bean
	// mux router
	Router *mux.Router
	// all mthods to declare
	methods []APIMethod
	// Router with injection mecanism
	SetRouterBean func(interface{}) `bean:"router"`
	RouterBean    *Router
}

// APIMethod single structure to modelise api declaration
type APIMethod struct {
	path    string
	handler string
	method  string
	addr    reflect.Value
}

// APIInterface all package methods
type APIInterface interface {
	bean.IBean
	Declare(APIMethod, interface{})
	GetMethods() []APIMethod
	HandlerStatic() func(w http.ResponseWriter, r *http.Request)
}

// Init initialize the API
func (api *API) Init() {
	api.SetRouterBean = func(value interface{}) {
		if assertion, ok := value.(*Router); ok {
			api.RouterBean = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// build arguments
	arr := [1]reflect.Value{reflect.ValueOf(api)}
	var arguments = arr[1:1]
	// build all static acess to low level function (private)
	var config = api.GetMethods()
	for i := 0; i < len(config); i++ {
		// compute rvalue
		var rvalue = config[i].addr.Call(arguments)[0]
		// declare this new method
		api.Declare(config[i], rvalue.Interface())
	}
}

// PostConstruct this API
func (api *API) PostConstruct(name string) error {
	return api.Bean.PostConstruct(name)
}

// GetMethods retrieve all method to declare in router
func (api *API) GetMethods() []APIMethod {
	return api.methods
}

// Declare a new interface
func (api *API) Declare(data APIMethod, intf interface{}) error {
	var result error
	// verify type
	if value, ok := intf.(func(http.ResponseWriter, *http.Request)); ok {
		log.Printf("Declare interface %s on %s with method %s (%s)", data.handler, data.path, data.method, (*api.RouterBean).GetName())
		// declare it to the router
		(*api.RouterBean).HandleFunc(data.path, value, data.method, "application/json")
		result = nil
	} else {
		// Error case
		result = errors.New("Unable to find any type for " + data.handler)
	}
	return result
}

// HandlerStatic is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func (api *API) HandlerStatic() func(w http.ResponseWriter, r *http.Request) {
	anonymous := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "{}")
	}
	return anonymous
}
