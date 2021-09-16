package golang

type Artikelpreis struct {
	ArtikelpreisNr     int                `json:"ArtikelpreisNr"`
	Artikel            Artikel            `json:"Artikel"`
	Artikelklasse      Artikelklasse      `json:"Artikelklasse"`
	Artikelgruppe      Artikelgruppe      `json:"Artikelgruppe"`
	Artikeluntergruppe Artikeluntergruppe `json:"Artikeluntergruppe"`
	Preisliste         Preisliste         `json:"Preisliste"`
	Verkaufspreis1     float64            `json:"Verkaufspreis1"`
	Verkaufspreis2     float64            `json:"Verkaufspreis2"`
	Verkaufspreis3     float64            `json:"Verkaufspreis3"`
	Verkaufspreis4     float64            `json:"Verkaufspreis4"`
	Verkaufspreis5     float64            `json:"Verkaufspreis5"`
}
