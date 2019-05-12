package golang

type Stundenart struct {
	StundenartNr string  `json:StundenartNr`
	Bezeichnung  string  `json:Bezeichnung`
	ZeitErfassen bool    `json:ZeitErfassen`
	Vorgabezeit  float64 `json:Vorgabezeit`
}
