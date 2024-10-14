package db

type Spirits struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Icon string `json:"icon"`
	Star int    `json:"star"`
}

func (s *Spirits) GetName() string {
	if s == nil {
		return ""
	}
	return GetText(s.Name)
}
