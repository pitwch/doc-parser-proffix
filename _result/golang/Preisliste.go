package golang

type Preisliste struct {
	PreislisteNr string   `json:PreislisteNr`
	Bezeichnung  string   `json:Bezeichnung`
	GueltigVon   string   `json:GueltigVon`
	GueltigBis   string   `json:GueltigBis`
	Netto        bool     `json:Netto`
	Waehrung     Waehrung `json:Waehrung`
}
