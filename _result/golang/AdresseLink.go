package golang

type AdresseLink struct {
	AdresseLinkNr int         `json:"AdresseLinkNr"`
	Adresse       Adresse     `json:"Adresse"`
	Preisgruppe   Preisgruppe `json:"Preisgruppe"`
	Preisklasse   Preisklasse `json:"Preisklasse"`
	Preisliste    Preisliste  `json:"Preisliste"`
}
