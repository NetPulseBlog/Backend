package entity

// NewsLineDefault
const (
	NWDPopular = iota + 1
	NWDFresh
)

// NewsLineSort
const (
	NWSByPopular = iota + 1
	NWSByDate
)

type UserSettings struct {
	NewsLineDefault string `json:"newsLineDefault"`
	NewsLineSort    string `json:"newsLineSort"`
}
