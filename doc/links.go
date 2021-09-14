package doc

import (
	"github.com/gocolly/colly"
	"github.com/pitw/doc-parser-proffix/model"
	"log"
	"strings"
)

func GetDocLinks() (doclinks []model.DocLink) {
	// Instantiate default collector
	c := colly.NewCollector()
	// On every a element which has href attribute call callback
	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fullLink := "https://www.proffix.net/Portals/0/content/REST%20API/" + link
		// Print link
		if strings.HasSuffix(link, "html") {
			doclinks = append(doclinks, model.DocLink{
				Name: e.Text,
				Link: fullLink,
			})

		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting %v", r.URL.String())
	})

	err := c.Visit("https://www.proffix.net/Portals/0/content/REST%20API/hmcontent.html")

	if err != nil {
		log.Print(err)
	}
	log.Print(doclinks)
	return doclinks
}
