package generator

import (
	"encoding/json"
	"github.com/pitw/doc-parser-proffix/model"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var pathGolang = "_result/golang/"

func CreateGoStruct(basepath string) (err error) {

	file, err := os.Open(basepath)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()
	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		err = BuildGoStruct(basepath+name, strings.TrimSuffix(name, ".json"))
	}

	return err
}

func BuildGoStruct(path string, name string) (err error) {

	byteValue := ReadJson(path)

	//We can use this again - nice...
	var doc model.Doc
	json.Unmarshal([]byte(byteValue), &doc)

	//Clean name
	cleanName := strings.Replace(name, " ", "", -1)
	cleanName = strings.Replace(name, "-", "", -1)

	//Start building the string
	mod := "package golang\n"
	mod += "type " + cleanName + " struct {\n"

	for _, d := range doc.Fields {
		mod += "\t" + d.Feld + "\t" + fixDatatype(d.Datentyp) + "\t" + "`json:" + d.Feld + "`" + "\n"

	}
	mod += "}"

	//Write to file
	err = ioutil.WriteFile(pathGolang+name+".go", []byte(mod), 0644)
	if err != nil {
		log.Fatal("Cannot write to File ", err)
	}
	return err
}

func fixDatatype(dt string) (fixed string) {
	if strings.Contains(dt, "String") {
		return "string"
	}

	if strings.Contains(dt, "int") {
		return "int"
	}

	if strings.Contains(dt, "Boolean") {
		return "bool"
	}

	if strings.Contains(dt, "Number") && !(strings.Contains(dt, "int")) {
		return "float64"
	}

	//Some hacky stuff...
	if strings.Contains(dt, "Object") {
		o := strings.Replace(dt, "Object", "", -1)
		o = strings.Replace(o, ">", "", -1)
		o = strings.Replace(o, "<", "", -1)
		return o
	}

	//Fallback if not found
	return "string"

}
