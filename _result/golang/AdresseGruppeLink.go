package golang

type AdresseGruppeLink struct {
	AdresseGruppeLinkNr int                `json:"AdresseGruppeLinkNr"`
	Adresse             Adresse            `json:"Adresse"`
	Artikelklasse       Artikelklasse      `json:"Artikelklasse"`
	Artikelgruppe       Artikelgruppe      `json:"Artikelgruppe"`
	Artikeluntergruppe  Artikeluntergruppe `json:"Artikeluntergruppe"`
	Preisgruppe         Preisgruppe        `json:"Preisgruppe"`
	Preisklasse         Preisklasse        `json:"Preisklasse"`
}
