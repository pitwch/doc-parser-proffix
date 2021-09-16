package golang

type BVGVersicherer struct {
	BVGVersichererNr int     `json:"BVGVersichererNr"`
	Bezeichnung      string  `json:"Bezeichnung"`
	Adresse          Adresse `json:"Adresse"`
	BVGCodes         string  `json:"BVGCodes"`
	GueltigAb        string  `json:"GueltigAb"`
	KundenNr         string  `json:"KundenNr"`
	VersichererNr    string  `json:"VersichererNr"`
	VertragsNr       string  `json:"VertragsNr"`
}
