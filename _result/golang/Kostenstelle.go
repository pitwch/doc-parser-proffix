package golang

type Kostenstelle struct {
	KostenstelleNr           string                   `json:KostenstelleNr`
	Bezeichnung              string                   `json:Bezeichnung`
	BudgetCode               int                      `json:BudgetCode`
	Kostenstellengruppe      Kostenstellengruppe      `json:Kostenstellengruppe`
	Kostenstellenklasse      Kostenstellenklasse      `json:Kostenstellenklasse`
	Kostenstellenuntergruppe Kostenstellenuntergruppe `json:Kostenstellenuntergruppe`
	Kumulation               bool                     `json:Kumulation`
	Sperre                   string                   `json:Sperre`
	Verantwortlicher         string                   `json:Verantwortlicher`
}
