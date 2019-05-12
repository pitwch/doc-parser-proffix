package golang

type Wareneingangsliste struct {
	WareneingangslisteNr string     `json:WareneingangslisteNr`
	Adresse              Adresse    `json:Adresse`
	Artikel              Artikel    `json:Artikel`
	Bestellung           Bestellung `json:Bestellung`
	DatumBis             string     `json:DatumBis`
	DatumVon             string     `json:DatumVon`
	Eingangsdatum        string     `json:Eingangsdatum`
	Lagerort             Lagerort   `json:Lagerort`
	Lagerplatz           Lagerplatz `json:Lagerplatz`
	Total                float64    `json:Total`
	Verarbeitet          bool       `json:Verarbeitet`
}
