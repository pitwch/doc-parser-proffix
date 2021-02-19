package golang

type Kostenart struct {
	KostenartNr      string `json:KostenartNr`
	Bezeichnung      string `json:Bezeichnung`
	BudgetCode       int    `json:BudgetCode`
	Sperre           string `json:Sperre`
	Verantwortlicher string `json:Verantwortlicher`
}
