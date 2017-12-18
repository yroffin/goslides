// Package business for business interface
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
package business

import (
	"github.com/yroffin/goslides/bean"
	"github.com/yroffin/goslides/models"
)

// SlideBusiness internal members
type SlideBusiness struct {
	// Base component
	*bean.Bean
}

// ISlideBusiness interface
type ISlideBusiness interface {
	bean.IBean
	Get(id string) (models.SlideBean, error)
	Create(id string, toCreate models.SlideBean) (models.SlideBean, error)
	Update(id string, toUpdate models.SlideBean) (models.SlideBean, error)
	Delete(id string) (models.SlideBean, error)
	Patch(id string, toPatch models.SlideBean) (models.SlideBean, error)
}

// Init Init this API
func (p *SlideBusiness) Init() error {
	return nil
}

// PostConstruct Init this API
func (p *SlideBusiness) PostConstruct(name string) error {
	return nil
}

// Validate Init this API
func (p *SlideBusiness) Validate(name string) error {
	return nil
}

// Get retrieve this slide by its id
func (p *SlideBusiness) Get(id string) (models.SlideBean, error) {
	var bean = models.SlideBean{ID: "1", Name: "Slide"}
	return bean, nil
}

// Create create a new slide
func (p *SlideBusiness) Create(toCreate models.SlideBean) (models.SlideBean, error) {
	var bean = models.SlideBean{ID: "2", Name: "New Slide"}
	return bean, nil
}

// Update a new slide
func (p *SlideBusiness) Update(id string, toUpdate models.SlideBean) (models.SlideBean, error) {
	var bean = models.SlideBean{ID: "3", Name: "Update Slide"}
	return bean, nil
}

// Delete a slide
func (p *SlideBusiness) Delete(id string) (models.SlideBean, error) {
	var bean = models.SlideBean{ID: "4", Name: "Deleted Slide"}
	return bean, nil
}

// Patch a slide
func (p *SlideBusiness) Patch(id string, toPatch models.SlideBean) (models.SlideBean, error) {
	var bean = models.SlideBean{ID: "5", Name: "Patched Slide"}
	return bean, nil
}
