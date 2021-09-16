package golang

type Serviceauftragsdokument struct {
	ServiceauftragsdokumentNr int            `json:"ServiceauftragsdokumentNr"`
	Serviceauftrag            Serviceauftrag `json:"Serviceauftrag"`
	Bemerkungen               string         `json:"Bemerkungen"`
	Bezeichnung               string         `json:"Bezeichnung"`
	DateiNr                   string         `json:"DateiNr"`
	Datum                     string         `json:"Datum"`
	Dokumentgruppe            string         `json:"Dokumentgruppe"`
	Drucken                   bool           `json:"Drucken"`
	Modul                     string         `json:"Modul"`
	ServiceauftragspositionNr int            `json:"ServiceauftragspositionNr"`
}
