package golang

type GS1 struct {
	Artikel      Artikel `json:"Artikel"`
	Feld         string  `json:"Feld"`
	Feldtyp      int     `json:"Feldtyp"`
	Gs1SatzLAG   string  `json:"Gs1SatzLAG"`
	Gtin         string  `json:"Gtin"`
	WertNummer   int     `json:"WertNummer"`
	WertString   string  `json:"WertString"`
	WertStringIT string  `json:"WertStringIT"`
	WertStringFR string  `json:"WertStringFR"`
	WertStringEN string  `json:"WertStringEN"`
}
