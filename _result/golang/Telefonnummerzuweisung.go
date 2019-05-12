package golang

type Telefonnummerzuweisung struct {
	TelefonNr      string  `json:TelefonNr`
	Adresse        Adresse `json:Adresse`
	Auftrag        Auftrag `json:Auftrag`
	Kontakt        Kontakt `json:Kontakt`
	Name           string  `json:Name`
	TelefonNrSuche string  `json:TelefonNrSuche`
}
