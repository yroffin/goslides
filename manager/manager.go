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
package manager

import (
	"log"

	"github.com/yroffin/goslides/bean"
)

// Manager interface
type Manager struct {
	Beans map[string]bean.BeanInterface
}

// ManagerInterface interface
type ManagerInterface interface {
	Init() error
	Register() error
	Boot() error
}

// Init a single bean
func (m *Manager) Init() {
	log.Printf("Manager::Init")
	m.Beans = map[string]bean.BeanInterface{}
}

// Register a single bean
func (m *Manager) Register(name string, bean bean.BeanInterface) error {
	m.Beans[name] = bean
	bean.SetName(name)
	log.Printf("Manager::Register %s", bean.GetName())
	return nil
}

// Boot Init this manager
func (m *Manager) Boot() error {
	log.Printf("Manager::Boot inject")
	for key, value := range m.Beans {
		value.Inject(key, m.Beans)
	}
	log.Printf("Manager::Boot post-construct")
	for key, value := range m.Beans {
		value.PostConstruct(key)
	}
	log.Printf("Manager::Boot validate")
	for key, value := range m.Beans {
		value.Validate(key)
	}
	return nil
}
