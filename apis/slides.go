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
	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	slide_models "github.com/yroffin/goslides/models"
)

// Slide internal members
type Slide struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	crud string `path:"/api/slides"`
}

// ISlide implements IBean
type ISlide interface {
	core_bean.IBean
}

// PostConstruct this API
func (p *Slide) Init() error {
	// Crud
	p.HandlerGetAll = func() (string, error) {
		return p.GenericGetAll(&slide_models.SlideBean{}, core_models.IPersistents(&slide_models.SlideBeans{Collection: make([]core_models.IPersistent, 0)}))
	}
	p.HandlerGetByID = func(id string) (string, error) {
		return p.GenericGetByID(id, &slide_models.SlideBean{})
	}
	p.HandlerPost = func(body string) (string, error) {
		return p.GenericPost(body, &slide_models.SlideBean{})
	}
	p.HandlerPutByID = func(id string, body string) (string, error) {
		return p.GenericPutByID(id, body, &slide_models.SlideBean{})
	}
	p.HandlerDeleteByID = func(id string) (string, error) {
		return p.GenericDeleteByID(id, &slide_models.SlideBean{})
	}
	p.HandlerPatchByID = func(id string, body string) (string, error) {
		return p.GenericPatchByID(id, body, &slide_models.SlideBean{})
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
