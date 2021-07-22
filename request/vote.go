package request

type Vote struct {
	Score   int  `json:"score"`
	Comment string  `json:"comment"`
}
