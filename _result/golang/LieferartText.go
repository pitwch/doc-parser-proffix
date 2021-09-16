package golang

type LieferartText struct {
	LieferartTextNr int     `json:"LieferartTextNr"`
	Bezeichnung     string  `json:"Bezeichnung"`
	Lieferart       string  `json:"Lieferart"`
	Modul           Sprache `json:"Modul"`
}
