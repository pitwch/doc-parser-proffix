package golang

type Spesenart struct {
	SpesenartNr  int          `json:"SpesenartNr"`
	Bezeichnung  string       `json:"Bezeichnung"`
	Betrag       float64      `json:"Betrag"`
	TextZwang    bool         `json:"TextZwang"`
	Steuercode   Steuercode   `json:"Steuercode"`
	Ertragskonto Konto        `json:"Ertragskonto"`
	Kostenart    Kostenart    `json:"Kostenart"`
	Kostenstelle Kostenstelle `json:"Kostenstelle"`
}
