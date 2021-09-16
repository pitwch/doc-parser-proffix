package golang

type Einheitsmatrix struct {
	EinheitLager    Einheit `json:"EinheitLager"`
	EinheitRechnung Einheit `json:"EinheitRechnung"`
	Bezeichnung     string  `json:"Bezeichnung"`
	EinheitDim1     string  `json:"EinheitDim1"`
	EinheitDim2     string  `json:"EinheitDim2"`
	EinheitDim3     string  `json:"EinheitDim3"`
	RechnenDim1     bool    `json:"RechnenDim1"`
	RechnenDim2     bool    `json:"RechnenDim2"`
	RechnenDim3     bool    `json:"RechnenDim3"`
	Divisor         float64 `json:"Divisor"`
	Multiplikator   float64 `json:"Multiplikator"`
	Rundung         int     `json:"Rundung"`
}
