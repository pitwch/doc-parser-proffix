package main

import (
	"encoding/json"
	"github.com/pitw/doc-parser-proffix/doc"
	"github.com/pitw/doc-parser-proffix/generator"
	"github.com/pitw/doc-parser-proffix/model"
	"github.com/pitw/doc-parser-proffix/table"
	"io/ioutil"
	"log"
	"strconv"
)

var basePath = "_result/json_base/"

func main() {

	//Hmm...
	log.Println("Do ut des")

	//Build base model / parse the tricky tables from PROFFIX in the hope they do a lot of copy & paste...
	createBaseJson()

	//Build our models from base model
	createModels()

}

func createModels() {

	//Create Models for Golang
	err := generator.CreateGoStruct(basePath)

	if err != nil {
		log.Printf("Error on create Go struct: %v", err)
	}
	//Create CSV
	err = generator.CreateCSV(basePath)
	if err != nil {
		log.Printf("Error on create CSV: %v", err)
	}

}

func createBaseJson() {

	// Object with all entities
	allDocs := doc.CreateDocsAll()

	//Collect links
	doclinks := doc.GetDocLinks()
	for _, y := range doclinks {

		//Check if starts with int -> Release -> we don't need em
		if _, err := strconv.Atoi(y.Name[:1]); err != nil {

			mod := table.TableToStrings(y.Link)
			//doc := table.TableToStrings("https://www.proffix.net/Portals/0/content/REST%20API/export/lohnbewegung.html")

			docsJson, err := json.MarshalIndent(mod, "", " ")

			if err != nil {
				log.Fatal("Cannot encode to JSON ", err)
			}

			//Check if it makes sense to parse...
			if len(mod.Fields) > 0 {

				err = ioutil.WriteFile(basePath+y.Name+".json", docsJson, 0644)
				if err != nil {
					log.Fatal("Cannot write to File ", err)
				}

				// Add entity to object
				allDocs.Entities = append(allDocs.Entities, model.DocLink{
					Name: y.Name,
					Link: y.Link,
					Doc:  mod,
				})
			}
		}
	}

	// Serialize object with all entities
	allDocsJson, err := json.MarshalIndent(allDocs, "", " ")

	// Write JSON aof all entities to file "_AllEntities.json"
	err = ioutil.WriteFile(basePath+"_AllEntities.json", allDocsJson, 0644)
	if err != nil {
		log.Fatal("Cannot write to File ", err)
	}
}
