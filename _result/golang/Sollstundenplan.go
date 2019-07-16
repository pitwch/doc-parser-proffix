package golang

type Sollstundenplan struct {
	SollstundenplanNr int     `json:SollstundenplanNr`
	Bezeichnung       string  `json:Bezeichnung`
	StundenMontag     float64 `json:StundenMontag`
	StundenDienstag   float64 `json:StundenDienstag`
	StundenMittwoch   float64 `json:StundenMittwoch`
	StundenDonnerstag float64 `json:StundenDonnerstag`
	StundenFreitag    float64 `json:StundenFreitag`
	StundenSamstag    float64 `json:StundenSamstag`
	StundenSonntag    float64 `json:StundenSonntag`
	StundenFerientag  float64 `json:StundenFerientag`
}
