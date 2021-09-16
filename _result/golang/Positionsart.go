package golang

type Positionsart struct {
	PositionsartNr               int          `json:"PositionsartNr"`
	Bezeichnung                  string       `json:"Bezeichnung"`
	Vorgabezeit                  float64      `json:"Vorgabezeit"`
	LohnperiodePlus              bool         `json:"LohnperiodePlus"`
	Waehrung                     Waehrung     `json:"Waehrung"`
	Steuercode                   Steuercode   `json:"Steuercode"`
	Ertragskonto                 Konto        `json:"Ertragskonto"`
	Kostenstelle                 Kostenstelle `json:"Kostenstelle"`
	Kostenart                    Kostenart    `json:"Kostenart"`
	KeinRabatt                   bool         `json:"KeinRabatt"`
	Ferien                       bool         `json:"Ferien"`
	StdPreis                     float64      `json:"StdPreis"`
	GueltigVon                   string       `json:"GueltigVon"`
	GueltigBis                   string       `json:"GueltigBis"`
	Minimalzeit                  float64      `json:"Minimalzeit"`
	Runden                       float64      `json:"Runden"`
	Lohnart                      Lohnart      `json:"Lohnart"`
	LohnartÜberzeit              Lohnart      `json:"LohnartÜberzeit"`
	LohnartKm                    Lohnart      `json:"LohnartKm"`
	LohnartSpesen                Lohnart      `json:"LohnartSpesen"`
	Stundenabrechnung            bool         `json:"Stundenabrechnung"`
	Verrechnen                   bool         `json:"Verrechnen"`
	NichtZuVerrIstStundenRechnen bool         `json:"NichtZuVerrIstStundenRechnen"`
	NichtZusammenziehen          bool         `json:"NichtZusammenziehen"`
}
