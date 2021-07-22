package request

type Film struct {
	File           string   `json:"file"`
	Name           string   `json:"name"`
	Producers      []string `json:"producers"`
	ProductionYear int      `json:"production_year"`
	Explanation    string   `json:"explanation"`
	View           int      `json:"view"`
	Price          int      `json:"price"`
	Tags           []string `json:"tags"`
}


//{
//"file": "gobol",
//"name": "raha",
//"producers": ["a", "b"],
//"production_year": 1373,
//"explanation": "koochooloo",
//"price": 20,
//"tags": ["a", "b"]
//}