package golang

type Depot struct {
	DepotNr        int     `json:"DepotNr"`
	Artikel        Artikel `json:"Artikel"`
	DepotArtikel   Artikel `json:"DepotArtikel"`
	Einheit        Einheit `json:"Einheit"`
	Haupt          bool    `json:"Haupt"`
	Inhalt         float64 `json:"Inhalt"`
	PaletteAnzahl  float64 `json:"PaletteAnzahl"`
	PaletteArtikel string  `json:"PaletteArtikel"`
}
