package db

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed *.json
var fs embed.FS

var (
	attributes         = make(map[int]string)
	languageWords      = make(map[string]*LanguageWord)
	systemInfos        = make(map[int]*SystemInfo)
	equipments         = make(map[int]*Equipment)
	equipmentQualities = make(map[int]string)
	spirits            = make(map[int]*Spirits)
	realms             = make(map[int]*Realms)
	gameSkills         = make(map[int]*GameSkill)
	cmds               = make(map[int32]*Cmd)
	recvCmds           = make(map[int32]*Cmd)
)

func init() {
	load("AttributeDB.json", &attributes)
	load("LanguageWordDB.json", &languageWords)
	load("SystemInfoDB.json", &systemInfos)
	load("EquipmentDB.json", &equipments)
	load("EquipmentQualityDB.json", &equipmentQualities)
	load("SpiritsDB.json", &spirits)
	load("RealmsDB.json", &realms)
	load("GameSkillDB.json", &gameSkills)

	initCmds()
}

func load(name string, v any) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(bytes, v); err != nil {
		panic(err)
	}
}

func GetText(id string) string {
	if word, ok := languageWords[id]; ok {
		return word.Text
	}
	return ""
}

func GameError(ret interface{ GetRet() int32 }) error {
	if ret == nil || ret.GetRet() == 0 {
		return nil
	}

	return GetGameError(ret.GetRet())
}

func GetGameError(id int32) error {
	if id == 0 {
		return nil
	}
	return fmt.Errorf("%d: %s", id, GetText(fmt.Sprintf("Game_Error-%d", id)))
}

func GetItemName(id int32) string {
	return GetText(fmt.Sprintf("Items-%d", id))
}

func GetSystemInfo(id int) *SystemInfo {
	return systemInfos[id]
}

func ListSystemInfos() []*SystemInfo {
	var _systemInfos []*SystemInfo
	for _, systemInfo := range systemInfos {
		_systemInfos = append(_systemInfos, systemInfo)
	}
	return _systemInfos
}

func GetEquipment(id int) *Equipment {
	return equipments[id]
}

func GetEquipmentQuality(id int) string {
	return equipmentQualities[id]
}

func GetSpirits(id int) *Spirits {
	return spirits[id]
}

func GetRealms(id int) *Realms {
	return realms[id]
}

func GetAttribute(id int) string {
	return attributes[id]
}
