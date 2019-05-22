package generator

import (
	"encoding/csv"
	"encoding/json"
	"github.com/pitw/doc-parser-proffix/model"
	"log"
	"os"
	"strings"
)

var pathCSV = "_result/csv/"

var methodsCSV = []string{}

func CreateCSV(basepath string) (err error) {

	file, err := os.Open(basepath)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()
	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		err = BuildCSV(basepath+name, strings.TrimSuffix(name, ".json"))
		err = MethodsCSV(basepath+name, strings.TrimSuffix(name, ".json"))

	}

	err = writeMethodsCSV()
	if err != nil {
		log.Fatalf("Error on write CSV Methods: %v", err)

	}
	return err
}
func MethodsCSV(path string, name string) (err error) {
	byteValue := ReadJson(path)

	//We can use this again - nice...
	var doc model.Doc
	json.Unmarshal([]byte(byteValue), &doc)

	for _, d := range doc.Methods {
		methodsCSV = append(methodsCSV, name)
		methodsCSV = append(methodsCSV, d.Endpunkt)
		methodsCSV = append(methodsCSV, d.Beschreibung)
		methodsCSV = append(methodsCSV, d.Rueckgabewert)
		methodsCSV = append(methodsCSV, d.BenoetigteLizenz)
		methodsCSV = append(methodsCSV, "\r\n")

	}
	return err
}

func writeMethodsCSV() (err error) {
	file, err := os.Create(pathCSV + "Methoden" + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %v", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(methodsCSV)
	if err != nil {
		log.Fatalf("failed creating file: %v", err)
	}
	return err
}
func BuildCSV(path string, name string) (err error) {

	byteValue := ReadJson(path)

	//We can use this again - nice...
	var doc model.Doc
	json.Unmarshal([]byte(byteValue), &doc)

	//Clean name
	//cleanName := strings.Replace(name, " ", "", -1)
	//cleanName = strings.Replace(name, "-", "", -1)

	//Write csv

	file, err := os.Create(pathCSV + name + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %v", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	line := []string{}
	for _, d := range doc.Fields {
		line = append(line, d.Feld)
		line = append(line, d.Datentyp)
		line = append(line, d.NamePROFFIX)
		line = append(line, d.Besonderes)
		line = append(line, "\r\n")

		err = writer.Write(line)
		if err != nil {
			log.Fatalf("failed creating file: %v", err)
		}

	}

	return err
}
