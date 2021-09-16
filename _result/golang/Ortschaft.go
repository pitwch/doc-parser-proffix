package golang

type Ortschaft struct {
	OrtschaftNr int    `json:"OrtschaftNr"`
	GemeindeNr  int    `json:"GemeindeNr"`
	Region      Region `json:"Region"`
	Land        Land   `json:"Land"`
	PLZ         string `json:"PLZ"`
	Ort         string `json:"Ort"`
}
