package entity

type UserSettings struct {
	// todo add enum for NewsLineDefault
	NewsLineDefault string `json:"newsLineDefault"`
	NewsLineSort    string `json:"newsLineSort"`
}
