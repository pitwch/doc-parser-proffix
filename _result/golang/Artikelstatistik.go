package golang

type Artikelstatistik struct {
	ArtikelstatistikNr int           `json:ArtikelstatistikNr`
	Adresse            Adresse       `json:Adresse`
	Artikel            Artikel       `json:Artikel`
	Auftrag            Auftrag       `json:Auftrag`
	Beleg              string        `json:Beleg`
	Bemerkung          string        `json:Bemerkung`
	Bestellung         Bestellung    `json:Bestellung`
	Charge             bool          `json:Charge`
	Datum              string        `json:Datum`
	Dim1               float64       `json:Dim1`
	Dim2               float64       `json:Dim2`
	Dim3               float64       `json:Dim3`
	Dokument           Dokument      `json:Dokument`
	Einheit            Einheit       `json:Einheit`
	EinheitRechnung    Einheit       `json:EinheitRechnung`
	Lagerort           Lagerort      `json:Lagerort`
	Lagerplatz         Lagerplatz    `json:Lagerplatz`
	MengeAus           float64       `json:MengeAus`
	MengeEin           float64       `json:MengeEin`
	MengeVerr          float64       `json:MengeVerr`
	Installation       Installation  `json:Installation`
	Preis              float64       `json:Preis`
	Rahmenvertrag      Rahmenvertrag `json:Rahmenvertrag`
	SerieNr            string        `json:SerieNr`
	Typ                string        `json:Typ`
	Waehrung           Waehrung      `json:Waehrung`
}
