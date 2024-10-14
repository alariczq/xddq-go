package db

type SystemInfo struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Icon          int    `json:"icon"`
	Sort          int    `json:"sort"`
	ConditionType string `json:"conditionType"`
	Value         string `json:"value"`
	Reward        string `json:"reward"`
	Notice        int    `json:"notice"`
	Collect       int    `json:"collect"`
	Condition     int    `json:"condition"`
	IsActivity    int    `json:"isActivity"`
}

func (s *SystemInfo) GetName() string {
	if s == nil {
		return ""
	}
	return GetText(s.Name)
}

func (s *SystemInfo) GetDesc() string {
	if s == nil {
		return ""
	}
	return GetText(s.Desc)
}
