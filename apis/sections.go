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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	rice "github.com/GeertJohan/go.rice"

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	slide_models "github.com/yroffin/goslides/models"
)

// Section internal members
type Section struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// Resource
	templateBox *rice.Box
	// mounts
	crud         string `path:"/api/sections"`
	presentation string `path:"/api/presentation/{id:[0-9a-zA-Z-_]*}" handler:"Render" method:"GET" mime-type:""`
}

// ISection implements IBean
type ISection interface {
	core_bean.IBean
}

// PostConstruct this API
func (p *Section) Init() error {
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
	// find a rice.Box
	resources, err := rice.FindBox("../resources")
	if err != nil {
		log.Fatal(err)
	}
	p.templateBox = resources
	return p.API.Init()
}

// PostConstruct this API
func (p *Section) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p)
	return nil
}

// Validate this API
func (p *Section) Validate(name string) error {
	return nil
}

// RenderSectionContext for rendering the html
type RenderSectionContext struct {
	Namee   string
	Section *Section
}

// Wget load in local the resource
func (p *RenderSectionContext) Wget(typ string, resource string) string {
	resp, _ := http.Get(resource)

	body, _ := ioutil.ReadAll(resp.Body)
	var raw = string(body)
	log.Printf("Successfully loaded %d bytes from %v", len(raw), resource)

	// analyze CSS to transform all url call
	if typ == "css" {
		/*
			re := regexp.MustCompile(`[u][r][l][(].*[)]`)
			matches := re.FindStringSubmatch(raw)
			for i := 0; i < len(matches); i++ {
				log.Printf("MATCHES %d '%v'", i, matches[i])
			}
		*/
	}
	return raw
}

// Sections render all slides
func (p *RenderSectionContext) Sections() string {
	// retrieve all slides
	var model = slide_models.SlideBean{}
	var models = slide_models.SlideBeans{Collection: make([]core_models.IPersistent, 0)}
	p.Section.CrudBusiness.GetAll(&model, core_models.IPersistents(&models))
	var stringBuffer string
	for index := 0; index < len(models.Collection); index++ {
		value, _ := models.Collection[index].(*slide_models.SlideBean)
		stringBuffer += fmt.Sprintf("<section>\n%s\n</section>\n", value.Body)
	}
	return stringBuffer
}

// Render render presentation
func (p *Section) Render() func() (string, error) {
	anonymous := func() (string, error) {
		// get file contents as string
		templateString, err := p.templateBox.String("reveal/reveal.html")
		if err != nil {
			log.Fatal(err)
		}
		// parse and execute the template
		tmplMessage, err := template.New("default").Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
		// render the template
		var tpl bytes.Buffer
		context := new(RenderSectionContext)
		context.Section = p
		tmplMessage.Execute(&tpl, context)
		return tpl.String(), nil
	}
	return anonymous
}
