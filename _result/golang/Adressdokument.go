package golang

type Adressdokument struct {
	AdressdokumentNr int     `json:"AdressdokumentNr"`
	Adresse          Adresse `json:"Adresse"`
	Bemerkungen      string  `json:"Bemerkungen"`
	Bezeichnung      string  `json:"Bezeichnung"`
	DateiNr          string  `json:"DateiNr"`
	Datum            string  `json:"Datum"`
	Dokumentgruppe   string  `json:"Dokumentgruppe"`
	Kontakt          Kontakt `json:"Kontakt"`
	Modul            string  `json:"Modul"`
}
