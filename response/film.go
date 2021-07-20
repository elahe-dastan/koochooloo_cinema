package response

type Film struct {
	ID             int      `json:"id"`
	File           string   `json:"file"`
	Name           string   `json:"name"`
	Producers      []string `json:"producers"`
	ProductionYear int      `json:"production_year"`
	Explanation    string   `json:"explanation"`
	View           int      `json:"view"`
	Price          int      `json:"price"`
	Tags           []string `json:"tags"`
}
