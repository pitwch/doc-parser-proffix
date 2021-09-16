package golang

type KundBez struct {
	KundBezNr      int     `json:"KundBezNr"`
	AdressNr       int     `json:"AdressNr"`
	ArikelNrKunde  string  `json:"ArikelNrKunde"`
	ArtikelNrLager string  `json:"ArtikelNrLager"`
	Barcode        string  `json:"Barcode"`
	Bezeichnung1   string  `json:"Bezeichnung1"`
	Bezeichnung2   string  `json:"Bezeichnung2"`
	Bezeichnung3   string  `json:"Bezeichnung3"`
	Bezeichnung4   string  `json:"Bezeichnung4"`
	Bezeichnung5   string  `json:"Bezeichnung5"`
	EAN            string  `json:"EAN"`
	TrustBox       bool    `json:"TrustBox"`
	Verkaufspreis  float64 `json:"Verkaufspreis"`
}
