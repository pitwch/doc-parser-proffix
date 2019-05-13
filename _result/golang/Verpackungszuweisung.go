package golang

type Verpackungszuweisung struct {
	VerpackungszuweisungNr int              `json:VerpackungszuweisungNr`
	Bezeichnung            string           `json:Bezeichnung`
	Artikel                Artikel          `json:Artikel`
	Geloescht              bool             `json:Geloescht`
	Gtin                   string           `json:Gtin`
	GtinStufe              string           `json:GtinStufe`
	Haupt                  bool             `json:Haupt`
	Menge                  float64          `json:Menge`
	Verpackung             Verpackung       `json:Verpackung`
	Verpackungsstufe       Verpackungsstufe `json:Verpackungsstufe`
}
