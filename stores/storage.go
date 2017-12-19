// Package stores for all sgbd operation
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
package stores

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	// for import driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/yroffin/goslides/bean"
)

// Store internal members
type Store struct {
	// Base component
	*bean.Bean
	// Store SQL lite
	database *sql.DB
}

// IStore interface
type IStore interface {
	bean.IBean
}

// Init Init this bean
func (p *Store) Init() error {
	return nil
}

// PostConstruct this bean
func (p *Store) PostConstruct(name string) error {
	// Create database
	database, _ := sql.Open("sqlite3", "./database.db")
	p.database = database

	var tables []string
	tables = append(tables, "SlideBean")
	for i := 0; i < len(tables); i++ {
		// prepare statement
		statement, _ := p.database.Prepare("CREATE TABLE IF NOT EXISTS " + tables[i] + " (id TEXT NOT NULL PRIMARY KEY, json TEXT)")
		statement.Exec()
	}

	return nil
}

// Validate Init this bean
func (p *Store) Validate(name string) error {
	return nil
}

// uuid generates a random UUID according to RFC 4122
func (p *Store) uuid(entity interface{}, set func(id string)) (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	var text = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	set(text)
	return text, nil
}

// Create this persistent bean n store
func (p *Store) Create(entity interface{}, set func(id string)) error {
	// get entity name
	var entityName = reflect.TypeOf(entity).Elem().Name()
	// fix UUID
	uuid, _ := p.uuid(entity, set)
	// insert
	statement, _ := p.database.Prepare("INSERT INTO " + entityName + " (id, json) VALUES (?,?)")
	data, _ := json.Marshal(&entity)
	statement.Exec(uuid, string(data))
	return nil
}

// Update this persistent bean
func (p *Store) Update(id string, entity interface{}, set func(id string)) error {
	// get entity name
	var entityName = reflect.TypeOf(entity).Elem().Name()
	// Fix ID
	set(id)
	// prepare statement
	statement, _ := p.database.Prepare("UPDATE " + entityName + " SET json = ? WHERE id = ?")
	data, _ := json.Marshal(&entity)
	statement.Exec(id, string(data))
	return nil
}

// Update this persistent bean
func (p *Store) Delete(id string, entity interface{}, set func(id string)) error {
	// get entity name
	var entityName = reflect.TypeOf(entity).Elem().Name()
	// Fix ID
	set(id)
	// prepare statement
	statement, _ := p.database.Prepare("DELETE " + entityName + " SET json = ? WHERE id = ?")
	data, _ := json.Marshal(&entity)
	statement.Exec(id, string(data))
	return nil
}

// Get this persistent bean
func (p *Store) Get(id string, entity interface{}) error {
	// get entity name
	var entityName = reflect.TypeOf(entity).Elem().Name()
	// prepare statement
	rows, _ := p.database.Query("SELECT id, json FROM "+entityName+" WHERE id = ?", &id)
	var data string
	for rows.Next() {
		rows.Scan(&id, &data)
		var bin = []byte(data)
		json.Unmarshal(bin, &entity)
	}
	return nil
}