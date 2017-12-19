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
	"github.com/yroffin/goslides/bean"
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
}

// ISlide implements IBean
type ISlide interface {
	bean.IBean
}

// PostConstruct this API
func (p *Slide) Init() error {
	// Crud
	p.HandlerGetByID = func(id string) (string, error) {
		return p.genericGetByID(id, &models.SlideBean{})
	}
	p.HandlerPost = func(body string) (string, error) {
		return p.genericPost(body, &models.SlideBean{})
	}
	p.HandlerPutByID = func(id string, body string) (string, error) {
		return p.genericPutByID(id, body, &models.SlideBean{})
	}
	p.HandlerDeleteByID = func(id string) (string, error) {
		return p.genericDeleteByID(id, &models.SlideBean{})
	}
	p.HandlerPatchByID = func(id string, body string) (string, error) {
		return p.genericPatchByID(id, body, &models.SlideBean{})
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Slide) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p)
	return nil
}

// Validate this API
func (p *Slide) Validate(name string) error {
	return nil
}
