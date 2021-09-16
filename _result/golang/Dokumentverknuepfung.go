package golang

type Dokumentverknuepfung struct {
	DokumentverknuepfungNr             int         `json:"DokumentverknuepfungNr"`
	DokumentAktuell                    Dokument    `json:"DokumentAktuell"`
	DokumenttypAktuell                 Dokumenttyp `json:"DokumenttypAktuell"`
	Dokument                           Dokument    `json:"Dokument"`
	Dokumenttyp                        Dokumenttyp `json:"Dokumenttyp"`
	DokumentSchlussrechnung            Dokument    `json:"DokumentSchlussrechnung"`
	DokumenttypSchlussrechnung         Dokumenttyp `json:"DokumenttypSchlussrechnung"`
	DokumentLieferscheinSammelrechnung Dokument    `json:"DokumentLieferscheinSammelrechnung"`
	DokumentRechnungSammelrechnung     Dokument    `json:"DokumentRechnungSammelrechnung"`
	Parameter                          string      `json:"Parameter"`
}
