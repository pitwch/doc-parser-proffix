package golang

type Bestellvorschlag struct {
	BestellvorschlagNr  int                `json:BestellvorschlagNr`
	Artikel             Artikel            `json:Artikel`
	Bestellpunkt        float64            `json:Bestellpunkt`
	Bestellrhythmus     int                `json:Bestellrhythmus`
	Bestellvorschlag    int                `json:Bestellvorschlag`
	Artikelgruppe       Artikelgruppe      `json:Artikelgruppe`
	Artikelklasse       Artikelklasse      `json:Artikelklasse`
	Lagerort            Lagerort           `json:Lagerort`
	Lieferadresse       Adresse            `json:Lieferadresse`
	Lieferantadresse    Adresse            `json:Lieferantadresse`
	Minimalbestand      float64            `json:Minimalbestand`
	Maximalbestand      float64            `json:Maximalbestand`
	Mindestbestellmenge float64            `json:Mindestbestellmenge`
	Optimalbestellmenge float64            `json:Optimalbestellmenge`
	Sicherheitstage     int                `json:Sicherheitstage`
	Artikeluntergruppe  Artikeluntergruppe `json:Artikeluntergruppe`
	Wiederbeschaffung   int                `json:Wiederbeschaffung`
}
