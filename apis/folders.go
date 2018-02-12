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
	json "encoding/json"
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
	app_models "github.com/yroffin/goslides/models"
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
		return p.GenericGetAll(&app_models.FolderBean{}, core_models.IPersistents(&app_models.FolderBeans{Collection: make([]core_models.IPersistent, 0)}))
	}
	p.HandlerGetByID = func(id string) (string, error) {
		return p.GenericGetByID(id, &app_models.FolderBean{})
	}
	p.HandlerPost = func(body string) (string, error) {
		return p.GenericPost(body, &app_models.FolderBean{})
	}
	p.HandlerTasks = func(name string, body string) (string, error) {
		if name == "export" {
			// export handler
			return p.Download()
		}
		if name == "import" {
			// export handler
			return p.Upload(body)
		}
		return "", nil
	}
	p.HandlerPutByID = func(id string, body string) (string, error) {
		return p.GenericPutByID(id, body, &app_models.FolderBean{})
	}
	p.HandlerDeleteByID = func(id string) (string, error) {
		return p.GenericDeleteByID(id, &app_models.FolderBean{})
	}
	p.HandlerPatchByID = func(id string, body string) (string, error) {
		return p.GenericPatchByID(id, body, &app_models.FolderBean{})
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
	var folder = app_models.FolderBean{}
	folder.SetID(id)
	p.Folder.CrudBusiness.Get(&folder)
	var stringBuffer string
	for index := 0; index < len(folder.Children); index++ {
		log.Printf("Iterate on folder %v", folder.Children[index])
		var slide = app_models.SlideBean{}
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

// Download all section and data
func (p *Folder) Download() (string, error) {
	var export = app_models.FoldersExportBean{}
	// find all folders
	var folder = app_models.FolderBean{}
	var folders = app_models.FolderBeans{}
	p.CrudBusiness.GetAll(&folder, &folders)
	log.Printf("Export %d folder(s)", len(folders.Collection))
	export.Folders = make([]app_models.FolderBean, len(folders.Collection))
	for index := 0; index < len(folders.Collection); index++ {
		export.Folders[index] = *folders.Index(index)
	}
	// find all slides
	var slide = app_models.SlideBean{}
	var slides = app_models.SlideBeans{}
	p.Slide.CrudBusiness.GetAll(&slide, &slides)
	log.Printf("Export %d slide(s)", len(slides.Collection))
	export.Slides = make([]app_models.SlideBean, len(slides.Collection))
	for index := 0; index < len(slides.Collection); index++ {
		export.Slides[index] = *slides.Index(index)
	}
	data, _ := json.MarshalIndent(export, "\t", "\t")
	return string(data), nil
}

// Upload all section and data
func (p *Folder) Upload(body string) (string, error) {
	var upload = app_models.FoldersExportBean{}
	json.Unmarshal([]byte(body), &upload)
	// iterate on folders
	log.Printf("Import %d folder(s)", len(upload.Folders))
	p.CrudBusiness.Truncate(&app_models.FolderBean{})
	for index := 0; index < len(upload.Folders); index++ {
		log.Printf("Import folder %v", upload.Folders[index].GetID())
		p.CrudBusiness.Create(&upload.Folders[index])
	}
	// iterate on slides
	log.Printf("Import %d slide(s)", len(upload.Slides))
	p.Slide.CrudBusiness.Truncate(&app_models.SlideBean{})
	for index := 0; index < len(upload.Slides); index++ {
		log.Printf("Import slide %v", upload.Slides[index].GetID())
		p.Slide.CrudBusiness.Create(&upload.Slides[index])
	}
	data, _ := json.MarshalIndent(upload, "\t", "\t")
	return string(data), nil
}
