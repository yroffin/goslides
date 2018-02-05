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
	"reflect"
	"text/template"

	rice "github.com/GeertJohan/go.rice"

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	folder_models "github.com/yroffin/goslides/models"
)

// Folder internal members
type Folder struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// Resource
	templateBox *rice.Box
	// Slide with injection mecanism
	SetSlide func(interface{}) `bean:"slide"`
	Slide    *Slide
	// mounts
	crud         string `path:"/api/folders"`
	presentation string `path:"/api/presentation/{id:[0-9a-zA-Z-_]*}" handler:"RenderPresentation" method:"GET" mime-type:""`
}

// IFolder implements IBean
type IFolder interface {
	core_bean.IBean
}

// Init this API
func (p *Folder) Init() error {
	// inject store
	p.SetSlide = func(value interface{}) {
		if assertion, ok := value.(*Slide); ok {
			p.Slide = assertion
		} else {
			log.Fatalf("Unable to validate injection with %v type is %v", value, reflect.TypeOf(value))
		}
	}
	// Crud
	p.HandlerGetAll = func() (string, error) {
		return p.GenericGetAll(&folder_models.FolderBean{}, core_models.IPersistents(&folder_models.FolderBeans{Collection: make([]core_models.IPersistent, 0)}))
	}
	p.HandlerGetByID = func(id string) (string, error) {
		return p.GenericGetByID(id, &folder_models.FolderBean{})
	}
	p.HandlerPost = func(body string) (string, error) {
		return p.GenericPost(body, &folder_models.FolderBean{})
	}
	p.HandlerPutByID = func(id string, body string) (string, error) {
		return p.GenericPutByID(id, body, &folder_models.FolderBean{})
	}
	p.HandlerDeleteByID = func(id string) (string, error) {
		return p.GenericDeleteByID(id, &folder_models.FolderBean{})
	}
	p.HandlerPatchByID = func(id string, body string) (string, error) {
		return p.GenericPatchByID(id, body, &folder_models.FolderBean{})
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
func (p *Folder) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p)
	return nil
}

// Validate this API
func (p *Folder) Validate(name string) error {
	return nil
}

// RenderContext for rendering the html
type RenderContext struct {
	Namee  string
	Folder *Folder
	ID     string
}

// Wget load in local the resource
func (p *RenderContext) Wget(typ string, resource string) string {
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

// Folders render all slides
func (p *RenderContext) Folders(id string) string {
	log.Printf("Iterate on folder %v", id)
	// retrieve all slides
	var folder = folder_models.FolderBean{}
	folder.SetID(id)
	p.Folder.CrudBusiness.Get(&folder)
	var stringBuffer string
	for index := 0; index < len(folder.Children); index++ {
		log.Printf("Iterate on folder %v", folder.Children[index])
		var slide = folder_models.SlideBean{}
		slide.SetID(folder.Children[index].Reference)
		p.Folder.Slide.CrudBusiness.Get(&slide)
		stringBuffer += fmt.Sprintf("<section>\n%s\n</section>\n", slide.Body)
	}
	return stringBuffer
}

// RenderPresentation render presentation
func (p *Folder) RenderPresentation() func(string) (string, error) {
	anonymous := func(id string) (string, error) {
		log.Printf("Render presentation for %v", id)
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
		// ender the template
		var tpl bytes.Buffer
		context := new(RenderContext)
		context.ID = id
		context.Folder = p
		tmplMessage.Execute(&tpl, context)
		return tpl.String(), nil
	}
	return anonymous
}
