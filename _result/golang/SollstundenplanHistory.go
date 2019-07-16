package golang

type SollstundenplanHistory struct {
	SollstundenplanHistoryNr int             `json:SollstundenplanHistoryNr`
	DatumVon                 string          `json:DatumVon`
	DatumBis                 string          `json:DatumBis`
	Mitarbeiter              Mitarbeiter     `json:Mitarbeiter`
	Sollstundenplan          Sollstundenplan `json:Sollstundenplan`
}
