package golang

type Lohnabrechnung struct {
	LohnabrechnungNr        int                     `json:"LohnabrechnungNr"`
	Periode                 int                     `json:"Periode"`
	SubNr                   int                     `json:"SubNr"`
	Buchungsdatum           string                  `json:"Buchungsdatum"`
	Titel                   string                  `json:"Titel"`
	Positionen              string                  `json:"Positionen"`
	Datum                   string                  `json:"Datum"`
	LohnbewegungPositionNr  float64                 `json:"LohnbewegungPositionNr"`
	Mitarbeiter             Mitarbeiter             `json:"Mitarbeiter"`
	Lohnart                 Lohnart                 `json:"Lohnart"`
	Menge                   float64                 `json:"Menge"`
	Ansatz                  float64                 `json:"Ansatz"`
	Betrag                  float64                 `json:"Betrag"`
	Anrede                  string                  `json:"Anrede"`
	Name                    string                  `json:"Name"`
	Vorname                 string                  `json:"Vorname"`
	Strasse                 string                  `json:"Strasse"`
	Region                  Region                  `json:"Region"`
	Adresszeile1            string                  `json:"Adresszeile1"`
	Adresszeile2            string                  `json:"Adresszeile2"`
	Adresszeile3            string                  `json:"Adresszeile3"`
	Adresszeile4            string                  `json:"Adresszeile4"`
	Adresszeile5            string                  `json:"Adresszeile5"`
	PLZ                     string                  `json:"PLZ"`
	Ort                     string                  `json:"Ort"`
	PLZPostfach             string                  `json:"PLZPostfach"`
	Postfach                string                  `json:"Postfach"`
	Land                    Land                    `json:"Land"`
	Arbeitskanton           string                  `json:"Arbeitskanton"`
	Auftrag                 Auftrag                 `json:"Auftrag"`
	Beschaeftigungsgrad     float64                 `json:"Beschaeftigungsgrad"`
	Bezeichnung             string                  `json:"Bezeichnung"`
	HabenKonto              Konto                   `json:"HabenKonto"`
	SollKonto               Konto                   `json:"SollKonto"`
	Kostenart               Kostenart               `json:"Kostenart"`
	Kostenstelle            Kostenstelle            `json:"Kostenstelle"`
	Steuercode              Steuercode              `json:"Steuercode"`
	Unfallversicherungscode Unfallversicherungscode `json:"Unfallversicherungscode"`
}
