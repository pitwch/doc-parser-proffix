package golang

type Projekt struct {
	ProjektNr               string         `json:ProjektNr`
	Bezeichnung             string         `json:Bezeichnung`
	Adresse                 Adresse        `json:Adresse`
	Teilleistungen          string         `json:Teilleistungen`
	Abrechnungsart          Abrechnungsart `json:Abrechnungsart`
	Archivierungsdatum      string         `json:Archivierungsdatum`
	Archivierungsort        string         `json:Archivierungsort`
	Auftrag                 Auftrag        `json:Auftrag`
	Auftraggeber            Adresse        `json:Auftraggeber`
	Bemerkungen             string         `json:Bemerkungen`
	Gebaeudeart             Gebaeudeart    `json:Gebaeudeart`
	Intern                  bool           `json:Intern`
	KontaktAuftragsgeber    string         `json:KontaktAuftragsgeber`
	KontaktBauherr          string         `json:KontaktBauherr`
	KontaktRechnungsadresse string         `json:KontaktRechnungsadresse`
	KontaktVersandadresse   string         `json:KontaktVersandadresse`
	Kostenstelle            Kostenstelle   `json:Kostenstelle`
	Projektstart            string         `json:Projektstart`
	Rechnungsadresse        Adresse        `json:Rechnungsadresse`
	Statusart               Statusart      `json:Statusart`
	Versandadresse          Adresse        `json:Versandadresse`
	ZurVerrechnungFreigeben bool           `json:ZurVerrechnungFreigeben`
}
