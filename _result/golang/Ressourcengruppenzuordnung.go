package golang

type Ressourcengruppenzuordnung struct {
	RessourcengruppenzuordnungNr int              `json:RessourcengruppenzuordnungNr`
	Ressourcengruppe             Ressourcengruppe `json:Ressourcengruppe`
	Ressource                    Ressource        `json:Ressource`
}
