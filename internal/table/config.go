package table

type BoxConfig struct {
	Points int    `json:"points"`
	Answer string `json:"answer"`
}

type Config struct {
	Boxes []BoxConfig `json:"answers"`
}
