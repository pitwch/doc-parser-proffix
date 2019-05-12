package golang

type Region struct {
	RegionNr    string `json:RegionNr`
	Bezeichnung string `json:Bezeichnung`
	Land        Land   `json:Land`
}
