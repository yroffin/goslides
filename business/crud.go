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
	"log"
	"reflect"

	"github.com/yroffin/goslides/bean"
	"github.com/yroffin/goslides/models"
	"github.com/yroffin/goslides/stores"
)

// CrudBusiness internal members
type CrudBusiness struct {
	// Base component
	*bean.Bean
	// Router with injection mecanism
	SetStore func(interface{}) `bean:"store-manager"`
	Store    *stores.Store
}

// ICrudBusiness interface
type ICrudBusiness interface {
	bean.IBean
	Get(models.IPersistent) (interface{}, error)
	Create(models.IPersistent) (interface{}, error)
	Update(models.IPersistent) (interface{}, error)
	Delete(models.IPersistent) (interface{}, error)
	Patch(models.IPersistent) (interface{}, error)
}

// Init Init this API
func (p *CrudBusiness) Init() error {
	// inject store
	p.SetStore = func(value interface{}) {
		if assertion, ok := value.(*stores.Store); ok {
			p.Store = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	return nil
}

// PostConstruct Init this API
func (p *CrudBusiness) PostConstruct(name string) error {
	return nil
}

// Validate Init this API
func (p *CrudBusiness) Validate(name string) error {
	return nil
}

// Get retrieve this slide by its id
func (p *CrudBusiness) Get(toGet models.IPersistent) error {
	p.Store.Get(toGet.GetID(), &toGet)
	return nil
}

// Create create a new slide
func (p *CrudBusiness) Create(toCreate models.IPersistent) (interface{}, error) {
	p.Store.Create(&toCreate, func(id string) { toCreate.SetID(id) })
	return toCreate, nil
}

// Update a new slide
func (p *CrudBusiness) Update(toUpdate models.IPersistent) (interface{}, error) {
	p.Store.Update(toUpdate.GetID(), &toUpdate, func(id string) { toUpdate.SetID(id) })
	return toUpdate, nil
}

// Delete a slide
func (p *CrudBusiness) Delete(toDelete models.IPersistent) (interface{}, error) {
	p.Store.Delete(toDelete.GetID(), &toDelete, func(id string) { toDelete.SetID(id) })
	return toDelete, nil
}

// Patch a slide
func (p *CrudBusiness) Patch(toPatch models.IPersistent) (interface{}, error) {
	p.Store.Update(toPatch.GetID(), &toPatch, func(id string) { toPatch.SetID(id) })
	return toPatch, nil
}