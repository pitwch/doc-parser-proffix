package golang

type Notiz struct {
	NotizNr             int         `json:NotizNr`
	Artikel             Artikel     `json:Artikel`
	Mitarbeiter         Mitarbeiter `json:Mitarbeiter`
	Notizart            Notizart    `json:Notizart`
	Datum               string      `json:Datum`
	Termin              string      `json:Termin`
	Notiz               string      `json:Notiz`
	NotizRTF            string      `json:NotizRTF`
	AlarmDatumVon       string      `json:AlarmDatumVon`
	AlarmDatumBis       string      `json:AlarmDatumBis`
	Kundeninstallation  bool        `json:Kundeninstallation`
	Auftragsbearbeitung bool        `json:Auftragsbearbeitung`
	Einkauf             bool        `json:Einkauf`
	Intern              bool        `json:Intern`
	Webshop             bool        `json:Webshop`
	Sprache             Sprache     `json:Sprache`
}
