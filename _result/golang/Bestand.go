package golang

type Bestand struct {
	Artikel           Artikel    `json:"Artikel"`
	Bestand           float64    `json:"Bestand"`
	BestandVerr       float64    `json:"BestandVerr"`
	Lagerort          Lagerort   `json:"Lagerort"`
	Lagerplatz        Lagerplatz `json:"Lagerplatz"`
	Charge            Charge     `json:"Charge"`
	BestandVerfuegbar float64    `json:"BestandVerfuegbar"`
	Reserviert        float64    `json:"Reserviert"`
	Rueckstand        float64    `json:"Rueckstand"`
	Bestellt          float64    `json:"Bestellt"`
}
