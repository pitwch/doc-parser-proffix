package golang

type Artikelbild struct {
	ArtikelbildNr       int     `json:ArtikelbildNr`
	Bezeichnung         string  `json:Bezeichnung`
	Anzeigen            bool    `json:Anzeigen`
	Artikel             Artikel `json:Artikel`
	DateiNr             string  `json:DateiNr`
	Hauptbild           bool    `json:Hauptbild`
	Webshop             bool    `json:Webshop`
	WebshopBildErstellt bool    `json:WebshopBildErstellt`
	Dateipfad           string  `json:Dateipfad`
}
