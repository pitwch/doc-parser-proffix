package golang

type Preishistory struct {
	ArtikelpreisNr int        `json:ArtikelpreisNr`
	Artikel        Artikel    `json:Artikel`
	Datum          string     `json:Datum`
	AlteWaehrung   Waehrung   `json:AlteWaehrung`
	AlterPreis     float64    `json:AlterPreis`
	NeueWaehrung   Waehrung   `json:NeueWaehrung`
	NeuerPreis     float64    `json:NeuerPreis`
	Preisliste     Preisliste `json:Preisliste`
	Preistype      int        `json:Preistype`
}
