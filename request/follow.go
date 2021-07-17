package request

type Follow struct {
	Username  string `json:"username"`
	Following string `json:"following"`
}
