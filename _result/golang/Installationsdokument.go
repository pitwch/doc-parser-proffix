package golang

type Installationsdokument struct {
	InstallationsdokumentNr int          `json:InstallationsdokumentNr`
	Installation            Installation `json:Installation`
	Bemerkungen             string       `json:Bemerkungen`
	Bezeichnung             string       `json:Bezeichnung`
	DateiNr                 string       `json:DateiNr`
	Datum                   string       `json:Datum`
	Dokumentgruppe          string       `json:Dokumentgruppe`
	Drucken                 bool         `json:Drucken`
	Modul                   string       `json:Modul`
}
