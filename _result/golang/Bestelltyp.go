package golang

type Bestelltyp struct {
	BestelltypNr            string   `json:BestelltypNr`
	ArtikelDokumenteDrucken bool     `json:ArtikelDokumenteDrucken`
	Aufgabe                 Aufgabe  `json:Aufgabe`
	Belegart                Belegart `json:Belegart`
	Bestellofferte          bool     `json:Bestellofferte`
	Bezeichnung             string   `json:Bezeichnung`
	HauptBestellung         bool     `json:HauptBestellung`
	Bestellformular1        string   `json:Bestellformular1`
	Bestellformular2        string   `json:Bestellformular2`
	Bestellformular3        string   `json:Bestellformular3`
	Reihenfolge             float64  `json:Reihenfolge`
	StatusLagereingang      bool     `json:StatusLagereingang`
	StatusOffen             bool     `json:StatusOffen`
	StatusPreiskorrektur    bool     `json:StatusPreiskorrektur`
	Titel                   string   `json:Titel`
}
