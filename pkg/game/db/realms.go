package db

type Realms struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Desc                string `json:"desc"`
	Icon                string `json:"icon"`
	BigType             int    `json:"bigType"`
	LittleType          int    `json:"littleType"`
	DemonicMax          string `json:"demonicMax"`
	DemonicEachDream    string `json:"demonicEachDream"`
	DemonicEachPractice string `json:"demonicEachPractice"`
	AttackBase          string `json:"attackBase"`
	HpBase              string `json:"hpBase"`
	AttackDef           string `json:"attackDef"`
	SpeedBase           string `json:"speedBase"`
	EquipmentLevel      string `json:"equipmentLevel"`
	EquipmentType       string `json:"equipmentType"`
	Appearance          int    `json:"appearance"`
	DestinyEnergy       int    `json:"destinyEnergy"`
}

func (r *Realms) GetName() string {
	if r == nil {
		return ""
	}
	return GetText(r.Name)
}
