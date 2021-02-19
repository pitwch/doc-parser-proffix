package golang

type Preisstaffel struct {
	PreisstaffelNr         string                 `json:PreisstaffelNr`
	ArtikeltypNr           string                 `json:ArtikeltypNr`
	Artikeltyp             int                    `json:Artikeltyp`
	Assortierung           Assortierung           `json:Assortierung`
	Auftrag                Auftrag                `json:Auftrag`
	MengeAbDokument        bool                   `json:MengeAbDokument`
	MengeVerrAbDokument    bool                   `json:MengeVerrAbDokument`
	PreisAbDokument        bool                   `json:PreisAbDokument`
	TotalAbDokument        bool                   `json:TotalAbDokument`
	Preisberechnungsformel Preisberechnungsformel `json:Preisberechnungsformel`
	GueltigVon             string                 `json:GueltigVon`
	GueltigBis             string                 `json:GueltigBis`
	AdresstypNr            int                    `json:AdresstypNr`
	Adresstyp              int                    `json:Adresstyp`
	MengeVon               float64                `json:MengeVon`
	Pauschale              bool                   `json:Pauschale`
	Preistyp               Preistyp               `json:Preistyp`
	Prozent                bool                   `json:Prozent`
	Rabatt                 bool                   `json:Rabatt`
	Verkauf                bool                   `json:Verkauf`
	Waehrung               Waehrung               `json:Waehrung`
	Wert                   float64                `json:Wert`
	Zusatzartikel          Artikel                `json:Zusatzartikel`
}
