// Package interfaces for common interfaces
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
	"fmt"
	"net/http"
	"reflect"
)

// Slide internal members
type Slide struct {
	// Base component
	*API
	// internal members
	Name string
}

// SlideInterface Test all package methods
type ISlide interface {
	APIInterface
	HandlerStaticPOST() func(w http.ResponseWriter, r *http.Request)
}

// PostConstruct this API
func (api *Slide) PostConstruct(name string) error {
	// define all methods
	api.methods = []APIMethod{{path: "/api/slides", handler: "HandlerStatic", method: "GET", addr: reflect.ValueOf(api).MethodByName("HandlerStatic")}, {path: "/api/slides", handler: "HandlerStaticPOST", method: "POST", addr: reflect.ValueOf(api).MethodByName("HandlerStaticPOST")}}
	// Call base class
	api.API.Init()
	return nil
}

// Validate this API
func (api *Slide) Validate(name string) error {
	return nil
}

// HandlerStaticPOST is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func (api *Slide) HandlerStaticPOST() func(w http.ResponseWriter, r *http.Request) {
	anonymous := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(202)
		fmt.Fprintf(w, "{}")
	}
	return anonymous
}
