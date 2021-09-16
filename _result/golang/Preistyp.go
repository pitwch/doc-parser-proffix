package golang

type Preistyp struct {
	PreistypNr  int    `json:"PreistypNr"`
	Artikeltyp  int    `json:"Artikeltyp"`
	Bezeichnung string `json:"Bezeichnung"`
	Geloescht   bool   `json:"Geloescht"`
	Kundentyp   int    `json:"Kundentyp"`
	Lager       bool   `json:"Lager"`
	Nullen      bool   `json:"Nullen"`
	Reihenfolge int    `json:"Reihenfolge"`
	ShopPreis   bool   `json:"ShopPreis"`
	ZPreise     bool   `json:"ZPreise"`
}
