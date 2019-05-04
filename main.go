package main

import (
	"fmt"
	"github.com/gocolly/colly/debug"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type DocLink struct {
	Name string
	Link string
}
type Doc struct {
	Fields    []DocFields
	Methods   []DocMethods
	Parameter []DocParameter
}
type DocFields struct {
	Feld        string
	Datentyp    string
	NamePROFFIX string
	Besonderes  string
	seitVersion string
}

type DocMethods struct {
	Endpunkt         string
	Rueckgabewert    string
	Beschreibung     string
	BenoetigteLizenz string
	seitVersion      string
}

type DocParameter struct {
	Parameter              string
	Format                 string
	Beschreibung           string
	VerhaltenOhneParameter string
	seitVersion            string
}

func GetDocLinks() (doclinks []DocLink) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.Debugger(&debug.LogDebugger{}),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fullLink := "https://www.proffix.net/Portals/0/content/REST%20API/export/" + link

		// Print link
		if (strings.HasSuffix(link, "html")) {
			doclinks = append(doclinks, DocLink{
				Name: e.Text,
				Link: fullLink,
			})

		}
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.proffix.net/Portals/0/content/REST%20API/export/proffix_rest_api_entwicklerhandbuch_content.html")

	return doclinks
}

func TableToJson(link string) (parsed [][]string) {
	d := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	)
	// On every a element which has href attribute call callback
	d.OnHTML("tr", func(e *colly.HTMLElement) {

		tempMap := []string{}
		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			//log.Print(el.Text)
			tempMap = append(tempMap,el.Text)
		})
		parsed = append(parsed,tempMap)
	})

	d.Visit(link)
	return parsed
}

func ParsedTable(input [][]string) (doc Doc){
	for _,val :=range input{
		log.Print(val[0][1])
			if (val[1] == "Feld") {

				log.Print("TEST")
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
func main() {
	//doclinks := GetDocLinks()
	//
	//for _, y := range doclinks {
	//	fmt.Println(y.Name)
	//	fmt.Println(y.Link)

	doc := TableToJson("https://www.proffix.net/Portals/0/content/REST%20API/export/lohnbewegung.html")

	test := ParsedTable(doc)
	fmt.Print(test)
}
