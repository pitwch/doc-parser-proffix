package golang

type Garantiezusatz struct {
	GarantiezusatzNr int    `json:"GarantiezusatzNr"`
	Bezeichnung      string `json:"Bezeichnung"`
	Tage1            int    `json:"Tage1"`
	Tage2            int    `json:"Tage2"`
	Tage3            int    `json:"Tage3"`
	Zusatztext1      string `json:"Zusatztext1"`
	Zusatztext2      string `json:"Zusatztext2"`
	Zusatztext3      string `json:"Zusatztext3"`
}
