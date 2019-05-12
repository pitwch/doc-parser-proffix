package golang

type Waehrung struct {
	WaehrungNr  string  `json:WaehrungNr`
	Bezeichnung string  `json:Bezeichnung`
	Kurs        float64 `json:Kurs`
	Verhaeltnis int     `json:Verhaeltnis`
	Rundung     int     `json:Rundung`
}
