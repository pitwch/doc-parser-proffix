package golang

type Bank struct {
	BankNr          int    `json:BankNr`
	Name            string `json:Name`
	Strasse         string `json:Strasse`
	PLZ             string `json:PLZ`
	Ort             string `json:Ort`
	Land            Land   `json:Land`
	Telefon         string `json:Telefon`
	Fax             string `json:Fax`
	ClearingNummer  int    `json:ClearingNummer`
	PostcheckNummer string `json:PostcheckNummer`
	SwiftNummer     string `json:SwiftNummer`
	ServerTyp       int    `json:ServerTyp`
	Geloescht       bool   `json:Geloescht`
}
