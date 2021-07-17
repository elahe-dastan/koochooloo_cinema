package request

type Introduction struct {
	Username   string  `json:"username"`
	Introducer string  `json:"introducer"`
}
