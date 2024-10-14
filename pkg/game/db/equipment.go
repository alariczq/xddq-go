package db

type Equipment struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Desc               string `json:"desc"`
	Icon               string `json:"icon"`
	Quality            int    `json:"quality"`
	Level              string `json:"level"`
	Type               int    `json:"type"`
	PrimAttack         string `json:"primAttack"`
	PrimHP             string `json:"primHP"`
	PrimDenfence       string `json:"primDenfence"`
	PrimSpeed          string `json:"primSpeed"`
	SecAttributeRan    string `json:"secAttributeRan"`
	Stun               string `json:"stun"`
	CriticalHit        string `json:"criticalHit"`
	DoubleAttack       string `json:"doubleAttack"`
	Dodge              string `json:"dodge"`
	AttackBack         string `json:"attackBack"`
	LifeSteal          string `json:"lifeSteal"`
	SecDefAttributeRan string `json:"secDefAttributeRan"`
	ReStun             string `json:"reStun"`
	ReCriticalHit      string `json:"reCriticalHit"`
	ReDoubleAttack     string `json:"reDoubleAttack"`
	ReDodge            string `json:"reDodge"`
	ReAttackBack       string `json:"reAttackBack"`
	ReLifeSteal        string `json:"reLifeSteal"`
	Weight             int    `json:"weight"`
}

func (e *Equipment) GetName() string {
	if e == nil {
		return ""
	}
	return GetText(e.Name)
}

func (e *Equipment) GetQualityName() string {
	if e == nil {
		return ""
	}
	return GetEquipmentQuality(e.Quality)
}
