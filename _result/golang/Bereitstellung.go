package golang

type Bereitstellung struct {
	BereitstellungNr              int          `json:"BereitstellungNr"`
	BereitstellungAufInstallation Installation `json:"BereitstellungAufInstallation"`
	BereitstellungAufArtikel      Artikel      `json:"BereitstellungAufArtikel"`
	BereitstPosArt                Positionsart `json:"BereitstPosArt"`
	Artikel                       Artikel      `json:"Artikel"`
	Anzahl                        float64      `json:"Anzahl"`
	Buchen                        bool         `json:"Buchen"`
	Geloescht                     bool         `json:"Geloescht"`
	BuchenPer                     int          `json:"BuchenPer"`
}
