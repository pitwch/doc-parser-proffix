package table

import (
	"github.com/gocolly/colly"
	"github.com/pitw/doc-parser-proffix/model"
	"log"
	"strings"
)

//TableToStrings converts PROFFIX Doc Tables to Strings
func TableToStrings(link string) (parsed model.Doc) {
	d := colly.NewCollector()
	doc := model.Doc{}
	// On every a element which has href attribute do callback
	d.OnHTML("#innerdiv", func(e *colly.HTMLElement) {

		e.ForEach("p", func(_ int, et *colly.HTMLElement) {
			var tmpTableName = ""

			tmpTableName = et.Text

			if strings.HasPrefix(et.Text, "ADR_") ||
				strings.HasPrefix(et.Text, "AUF_") ||
				strings.HasPrefix(et.Text, "BAS_") ||
				strings.HasPrefix(et.Text, "CRM_") ||
				strings.HasPrefix(et.Text, "DEB_") ||
				strings.HasPrefix(et.Text, "EDO_") ||
				strings.HasPrefix(et.Text, "EIN_") ||
				strings.HasPrefix(et.Text, "FIB_") ||
				strings.HasPrefix(et.Text, "KRE_") ||
				strings.HasPrefix(et.Text, "LAG_") ||
				strings.HasPrefix(et.Text, "LOH_") ||
				strings.HasPrefix(et.Text, "PRE_") ||
				strings.HasPrefix(et.Text, "PRO_") ||
				strings.HasPrefix(et.Text, "RES_") ||
				strings.HasPrefix(et.Text, "SRV_") ||
				strings.HasPrefix(et.Text, "STU_") ||
				strings.HasPrefix(et.Text, "ZEI_") {

				log.Printf("Table %v", tmpTableName)

				doc.TableName = tmpTableName
			}
		})

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

			tmpString := []string{}
			//log.Print(el.Text)
			el.ForEach("td", func(_ int, ef *colly.HTMLElement) {
				//fmt.Printf("A: %v, B: %v\n", el.Index, ef.Index)

				//Remove line break \n
				cleanLineBreak := strings.Replace(ef.Text, "\n", "", -1)

				//If cell contain this *** symbol -> messed up table -> fix it
				if strings.Contains(cleanLineBreak, "•") {

					//log.Println(cleanLineBreak)
				}

				tmpString = append(tmpString, cleanLineBreak)

			})

			if len(tmpString) == 5 {
				firstField := tmpString[0]

				//If first cell begins with Method -> docMethod
				if strings.HasPrefix(firstField, "PUT") ||
					strings.HasPrefix(firstField, "DELETE") ||
					strings.HasPrefix(firstField, "PATCH") ||
					strings.HasPrefix(firstField, "POST") ||
					strings.HasPrefix(firstField, "GET") {

					//Dont include Header Field
					if tmpString[0] != "Feld" &&
						tmpString[0] != "Endpunkt" {
						doc.Methods = append(doc.Methods, model.DocMethods{
							Endpunkt:         tmpString[0],
							Rueckgabewert:    tmpString[1],
							Beschreibung:     tmpString[2],
							BenoetigteLizenz: tmpString[3],
						})
					}
				} else if firstUppercase(tmpString[0]) {

					//Dont include Header Field
					if tmpString[0] != "Endpunkt" && tmpString[0] != "Feld" {
						doc.Fields = append(doc.Fields, model.DocFields{
							Feld:        tmpString[0],
							Datentyp:    tmpString[1],
							NamePROFFIX: tmpString[2],
							Besonderes:  tmpString[3],
						})

						if len(doc.PrimaryKey) == 0 {
							doc.PrimaryKey = tmpString[0]
						}
					}
				} else {
					doc.Parameter = append(doc.Parameter, model.DocParameter{
						Parameter:              tmpString[0],
						Format:                 tmpString[1],
						Beschreibung:           tmpString[2],
						VerhaltenOhneParameter: tmpString[3],
					})
				}
			}

		})

	})

	d.Visit(link)
	return doc
}

func firstUppercase(str string) bool {
	t := str[:1]
	if strings.ToUpper(t) == t {
		return true
	} else {
		return false
	}
}

func ParsedTable(input [][]string) (doc model.Doc) {
	for _, val := range input {
		if len(val) > 1 {
			if !strings.Contains(val[0], "Referenz") {

			}
			//for k, _ := range input[i] {
			//
			//	docfield := DocFields{
			//		Feld:        input[i][0],
			//		Datentyp:    input[i][1],
			//		NamePROFFIX: input[i][2],
			//		Besonderes:  input[i][3],
			//		seitVersion: input[i][4],
			//	}
			//	log.Print(k)
			//	log.Print(docfield)
			//}
		}
	}

	return doc

}
