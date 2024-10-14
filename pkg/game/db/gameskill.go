package db

import (
	"strconv"
	"strings"
)

type GameSkill struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Icon          string `json:"icon"`
	Star          int    `json:"star"`
	Type          string `json:"type"`
	Params        string `json:"params"`
	UpgradeParams string `json:"upgradeParams"`
	MaxLevel      int    `json:"maxLevel"`
	UpgradeItem   string `json:"upgradeItem"`
	UpgradeType   string `json:"upgradeType"`
	Sort          int    `json:"sort"`
	TextStyle     int    `json:"textStyle"`
}

func (g *GameSkill) GetName() string {
	if g == nil {
		return ""
	}
	return GetText(g.Name)
}

func (g *GameSkill) GetDesc() string {
	if g == nil {
		return ""
	}
	return GetText(g.Desc)
}

func (g *GameSkill) GetParams() []int64 {
	if g == nil {
		return nil
	}

	params := make([]int64, 4)
	for _, v := range strings.Split(g.Params, "|") {
		n, _ := strconv.ParseInt(v, 10, 64)
		params = append(params, n)
	}

	return params
}

func (g *GameSkill) GetUpgradeParams() []int64 {
	if g == nil {
		return nil
	}

	params := make([]int64, 4)
	for _, v := range strings.Split(g.UpgradeParams, "|") {
		n, _ := strconv.ParseInt(v, 10, 64)
		params = append(params, n)
	}

	return params
}

func GetGameSkill(id int) *GameSkill {
	return gameSkills[id]
}

func ListGameSkills() []*GameSkill {
	var _skills []*GameSkill
	for _, skill := range gameSkills {
		_skills = append(_skills, skill)
	}
	return _skills
}
