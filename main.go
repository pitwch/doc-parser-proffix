package main

import (
	"encoding/json"
	"github.com/pitw/doc-parser-proffix/doc"
	"github.com/pitw/doc-parser-proffix/table"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {

	//Hmm...
	log.Println("Do ut des")

	//Collect links
	doclinks := doc.GetDocLinks()
	for _, y := range doclinks {

		//Check if starts with int -> Release -> we don't need em
		if _, err := strconv.Atoi(y.Name[:1]); err != nil {

			mod := table.TableToStrings(y.Link)
			//doc := table.TableToStrings("https://www.proffix.net/Portals/0/content/REST%20API/export/lohnbewegung.html")

			docsJson, err := json.Marshal(mod)
			if err != nil {
				log.Fatal("Cannot encode to JSON ", err)
			}
			err = ioutil.WriteFile("_result/json/"+y.Name+".json", docsJson, 0644)
			if err != nil {
				log.Fatal("Cannot write to File ", err)
			}
		}
	}

}
