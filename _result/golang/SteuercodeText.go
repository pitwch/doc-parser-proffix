package golang

type SteuercodeText struct {
	SteuercodeTextNr int        `json:"SteuercodeTextNr"`
	Bezeichung       string     `json:"Bezeichung"`
	Sprache          Sprache    `json:"Sprache"`
	Steuercode       Steuercode `json:"Steuercode"`
}
