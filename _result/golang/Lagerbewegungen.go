package golang

type Lagerbewegungen struct {
	Datum                  string   `json:Datum`
	Preis                  float64  `json:Preis`
	PreisSchreiben         bool     `json:PreisSchreiben`
	Waehrung               Waehrung `json:Waehrung`
	Kurs                   float64  `json:Kurs`
	Beleg                  string   `json:Beleg`
	Adresse                Adresse  `json:Adresse`
	Auftrag                Auftrag  `json:Auftrag`
	Bemerkungen            string   `json:Bemerkungen`
	LagerOrtPlatzAbArtikel bool     `json:LagerOrtPlatzAbArtikel`
}
