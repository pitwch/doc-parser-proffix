package golang

type Stuecklistenposition struct {
	StuecklistenpositionNr int     `json:"StuecklistenpositionNr"`
	Artikel                Artikel `json:"Artikel"`
	Bemerkungen            string  `json:"Bemerkungen"`
	BemerkungenRTF         string  `json:"BemerkungenRTF"`
	Lagerabtrag            bool    `json:"Lagerabtrag"`
	NichtAnzeigen          bool    `json:"NichtAnzeigen"`
	NichtBestellen         bool    `json:"NichtBestellen"`
	PositionNr             int     `json:"PositionNr"`
	PreisUebernehmen       bool    `json:"PreisUebernehmen"`
	PreisSW                float64 `json:"PreisSW"`
	Sprache                Sprache `json:"Sprache"`
	StuecklisteArtikel     Artikel `json:"StuecklisteArtikel"`
	Anzahl                 float64 `json:"Anzahl"`
}
