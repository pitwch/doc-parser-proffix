package golang

type Mitarbeiter struct {
	MitarbeiterNr                  int             `json:MitarbeiterNr`
	Name                           string          `json:Name`
	Kurzzeichen                    string          `json:Kurzzeichen`
	Adresse                        Adresse         `json:Adresse`
	Positionsart                   Positionsart    `json:Positionsart`
	Stundenart                     Stundenart      `json:Stundenart`
	Feriensaldo                    float64         `json:Feriensaldo`
	BadgeID                        string          `json:BadgeID`
	Geloescht                      bool            `json:Geloescht`
	Benutzergruppe                 Benutzergruppe  `json:Benutzergruppe`
	Bemerkungen                    string          `json:Bemerkungen`
	BemerkungenLeistungsverwaltung string          `json:BemerkungenLeistungsverwaltung`
	BemerkungenZeitverwaltung      string          `json:BemerkungenZeitverwaltung`
	Austrittsdatum                 string          `json:Austrittsdatum`
	Eintrittsdatum                 string          `json:Eintrittsdatum`
	Funktion                       Funktion        `json:Funktion`
	Sollstundenplan                Sollstundenplan `json:Sollstundenplan`
	ADUser                         string          `json:ADUser`
}
