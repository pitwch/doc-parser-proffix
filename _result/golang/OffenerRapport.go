package golang

type OffenerRapport struct {
	OffenerRapportNr          float64        `json:OffenerRapportNr`
	Mitarbeiter               Mitarbeiter    `json:Mitarbeiter`
	Auftrag                   Auftrag        `json:Auftrag`
	Bezeichnung               string         `json:Bezeichnung`
	Positionsart              Positionsart   `json:Positionsart`
	Bemerkungen               string         `json:Bemerkungen`
	BemerkungenRTF            string         `json:BemerkungenRTF`
	Serviceauftrag            Serviceauftrag `json:Serviceauftrag`
	ServiceauftragspositionNr float64        `json:ServiceauftragspositionNr`
	DauerTotal                float64        `json:DauerTotal`
	Kontaktname               string         `json:Kontaktname`
	Kontakt                   Kontakt        `json:Kontakt`
	Rapport                   Rapport        `json:Rapport`
	Pausiert                  bool           `json:Pausiert`
	Abgeschlossen             bool           `json:Abgeschlossen`
	Zeiten                    int            `json:Zeiten`
	Adresse                   Adresse        `json:Adresse`
	StartDatum                string         `json:StartDatum`
	StartZeit                 string         `json:StartZeit`
	EndDatum                  string         `json:EndDatum`
	EndZeit                   string         `json:EndZeit`
	Dauer                     float64        `json:Dauer`
}
