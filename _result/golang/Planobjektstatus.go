package golang

type Planobjektstatus struct {
	PlanobjektstatusNr int    `json:PlanobjektstatusNr`
	Bezeichnung        string `json:Bezeichnung`
	Farbe              int    `json:Farbe`
	Fixiert            bool   `json:Fixiert`
}
