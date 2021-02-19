package golang

type Artikeldokument struct {
	ArtikeldokumentNr int     `json:ArtikeldokumentNr`
	Artikel           Artikel `json:Artikel`
	Bemerkungen       string  `json:Bemerkungen`
	Bezeichnung       string  `json:Bezeichnung`
	DateiNr           string  `json:DateiNr`
	Datum             string  `json:Datum`
	Dokumentgruppe    string  `json:Dokumentgruppe`
	Drucken           bool    `json:Drucken`
	Modul             string  `json:Modul`
}
