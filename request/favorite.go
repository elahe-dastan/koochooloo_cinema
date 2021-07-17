package request

type Favorite struct {
	Username string `json:"username"`
	Film     []int    `json:"film"`
	Album    string `json:"album"`
}
