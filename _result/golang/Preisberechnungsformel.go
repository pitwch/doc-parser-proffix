package golang

type Preisberechnungsformel struct {
	PreisberechnungsformelNr int    `json:"PreisberechnungsformelNr"`
	Bezeichnung              string `json:"Bezeichnung"`
	Runden                   int    `json:"Runden"`
	RundenTyp                int    `json:"RundenTyp"`
}
