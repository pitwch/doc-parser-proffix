package golang

type Lagerort struct {
	LagerortNr     int        `json:"LagerortNr"`
	Artikel        Artikel    `json:"Artikel"`
	Lagerort       Lagerort   `json:"Lagerort"`
	Lagerplatz     Lagerplatz `json:"Lagerplatz"`
	Einkauf        float64    `json:"Einkauf"`
	Faktor         float64    `json:"Faktor"`
	FaktorFix      bool       `json:"FaktorFix"`
	Verkaufspreis1 float64    `json:"Verkaufspreis1"`
	Verkaufspreis2 float64    `json:"Verkaufspreis2"`
	Verkaufspreis3 float64    `json:"Verkaufspreis3"`
	Verkaufspreis4 float64    `json:"Verkaufspreis4"`
	Verkaufspreis5 float64    `json:"Verkaufspreis5"`
	Waehrung       Waehrung   `json:"Waehrung"`
}
