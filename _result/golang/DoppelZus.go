package golang

type DoppelZus struct {
	DoppelZusNr int     `json:"DoppelZusNr"`
	AdresseNeu  Adresse `json:"AdresseNeu"`
	AdresseAlt  Adresse `json:"AdresseAlt"`
}
