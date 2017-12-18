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
	"encoding/json"
	"log"
	"reflect"

	"github.com/yroffin/goslides/bean"
	"github.com/yroffin/goslides/business"
	"github.com/yroffin/goslides/models"
)

// Slide internal members
type Slide struct {
	// Base component
	*API
	// internal members
	Name string
	// mounts
	crud string `path:"/api/slides"`
	// Router with injection mecanism
	SetSlideBusiness func(interface{}) `bean:"slide-business"`
	SlideBusiness    *business.SlideBusiness
}

// ISlide implements IBean
type ISlide interface {
	bean.IBean
}

// PostConstruct this API
func (p *Slide) Init() error {
	// inject SlideBusiness
	p.SetSlideBusiness = func(value interface{}) {
		if assertion, ok := value.(*business.SlideBusiness); ok {
			p.SlideBusiness = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// Crud
	p.HandlerGetByID = func(id string) (string, error) {
		return p.GetByID(id)
	}
	p.HandlerPost = func(body string) (string, error) {
		return p.Post(body)
	}
	p.HandlerPutByID = func(id string, body string) (string, error) {
		return p.PutByID(id, body)
	}
	p.HandlerDeleteByID = func(id string) (string, error) {
		return p.DeleteByID(id)
	}
	p.HandlerPatchByID = func(id string, body string) (string, error) {
		return p.PatchByID(id, body)
	}
	return p.API.Init()
}

// PostConstruct this API
func (api *Slide) PostConstruct(name string) error {
	// Scan struct and init all handler
	api.ScanHandler(api)
	return nil
}

// Validate this API
func (api *Slide) Validate(name string) error {
	return nil
}

// GetByID default method
func (p *Slide) GetByID(id string) (string, error) {
	bean, _ := p.SlideBusiness.Get(id)
	data, _ := json.Marshal(&bean)
	return string(data), nil
}

// Post adefault method
func (p *Slide) Post(body string) (string, error) {
	var toCreate models.SlideBean
	var bin = []byte(body)
	json.Unmarshal(bin, &toCreate)
	bean, _ := p.SlideBusiness.Create(toCreate)
	data, _ := json.Marshal(&bean)
	return string(data), nil
}

// PutByID default method
func (p *Slide) PutByID(id string, body string) (string, error) {
	var toUpdate models.SlideBean
	var bin = []byte(body)
	json.Unmarshal(bin, &toUpdate)
	bean, _ := p.SlideBusiness.Update(id, toUpdate)
	data, _ := json.Marshal(&bean)
	return string(data), nil
}

// PatchByID default method
func (p *Slide) PatchByID(id string, body string) (string, error) {
	var toPatch models.SlideBean
	var bin = []byte(body)
	json.Unmarshal(bin, &toPatch)
	bean, _ := p.SlideBusiness.Patch(id, toPatch)
	data, _ := json.Marshal(&bean)
	return string(data), nil
}

// DeleteByID default method
func (p *Slide) DeleteByID(id string) (string, error) {
	old, _ := p.SlideBusiness.Delete(id)
	data, _ := json.Marshal(&old)
	return string(data), nil
}
