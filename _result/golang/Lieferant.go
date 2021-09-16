package golang

type Lieferant struct {
	LieferantNr       int       `json:"LieferantNr"`
	Adresse           Adresse   `json:"Adresse"`
	Artikel           Artikel   `json:"Artikel"`
	ArtikelLieferant  string    `json:"ArtikelLieferant"`
	Barcode           string    `json:"Barcode"`
	Dim1              float64   `json:"Dim1"`
	Dim2              float64   `json:"Dim2"`
	Dim3              float64   `json:"Dim3"`
	Einheit           Einheit   `json:"Einheit"`
	Einkaufspreis     float64   `json:"Einkaufspreis"`
	Nettopreis        float64   `json:"Nettopreis"`
	Name              string    `json:"Name"`
	Einstandspreis    float64   `json:"Einstandspreis"`
	GtinStufe         string    `json:"GtinStufe"`
	Hauptlieferant    bool      `json:"Hauptlieferant"`
	Lieferart         Lieferart `json:"Lieferart"`
	Rabatt1           float64   `json:"Rabatt1"`
	Rabatt2           float64   `json:"Rabatt2"`
	Waehrung          Waehrung  `json:"Waehrung"`
	Wiederbeschaffung int       `json:"Wiederbeschaffung"`
}
