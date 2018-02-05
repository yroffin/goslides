// Package models for all models
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
package models

import (
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
)

// SlideBean simple slide model
type SectionBean struct {
	// Id
	ID string `json:"id"`
	// Timestamp
	Timestamp core_models.JSONTime `json:"timestamp"`
	// Name
	Name string `json:"name"`
	// Slide
	Slide string `json:"slide"`
	// Slides
	Slides []string `json:"slides"`
}

// SetName get set name
func (p *SectionBean) SetName() string {
	return "Slide"
}

// GetID retrieve ID
func (p *SectionBean) GetID() string {
	return p.ID
}

// SetID retrieve ID
func (p *SectionBean) SetID(ID string) {
	p.ID = ID
}

// SetTimestamp set timestamp
func (p *SectionBean) SetTimestamp(stamp core_models.JSONTime) {
	p.Timestamp = stamp
}

// GetTimestamp get timestamp
func (p *SectionBean) GetTimestamp() core_models.JSONTime {
	return p.Timestamp
}

// Copy retrieve ID
func (p *SectionBean) Copy() core_models.IPersistent {
	clone := *p
	return &clone
}

// SectionBeans simple slide model
type SectionBeans struct {
	// Collection
	Collection []core_models.IPersistent
}

// Add new bean
func (p *SectionBeans) Add(slide core_models.IPersistent) {
	p.Collection = append(p.Collection, slide)
}

// Get collection of bean
func (p *SectionBeans) Get() []core_models.IPersistent {
	return p.Collection
}
