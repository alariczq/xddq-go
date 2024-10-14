package assist

import (
	"errors"
	"time"
)

type AssistantConfig struct {
	Username          string        `json:"username"`
	Password          string        `json:"password"`
	Disabled          bool          `json:"disabled"`
	ReconnectDuration time.Duration `json:"reconnectDuration"`

	ExchangeKey []string `json:"exchangeKey"`

	Challenge   ChallengeConfig `json:"challenge"`   // 是否开启挑战
	DailyTask   bool            `json:"dailyTask"`   // 是否开启日常任务
	Homeland    HomelandConfig  `json:"homeland"`    // 是否开启家园
	StarTrial   bool            `json:"startrial"`   // 是否开启星辰试炼
	Yard        bool            `json:"yard"`        // 是否开启仙居
	Union       UnionConfig     `json:"union"`       // 是否开启联盟
	Palace      bool            `json:"palace"`      // 是否开启仙宫
	Destiny     bool            `json:"destiny"`     // 是否开启仙友
	Universe    bool            `json:"universe"`    // 是否开启小世界
	SecretTower bool            `json:"secretTower"` // 是否开启秘境
	Yuebao      bool            `json:"yuebao"`      // 是否开启仙玉宝府
	Gather      GatherConfig    `json:"gather"`      // 是否开启采集
	Reward      bool            `json:"reward"`      // 是否开启领奖
	Role        RoleConfig      `json:"role"`        // 是否开启角色
}

type UnionConfig struct {
	CutPrice  bool `json:"cutPrice"`  // 是否开启砍价
	UnionTask bool `json:"unionTask"` // 是否开启妖盟任务
	Bounty    bool `json:"bounty"`    // 是否开启妖盟对战
}

type ChallengeConfig struct {
	Enabled        bool  `json:"enabled"`        // 是否开启挑战
	Tower          bool  `json:"tower"`          // 是否开启镇妖塔
	TowerIdx       int32 `json:"towerIdx"`       // 镇妖塔分身
	SecretTower    bool  `json:"secretTower"`    // 是否开启秘境
	SecretTowerIdx int32 `json:"secretTowerIdx"` // 秘境分身
	StageChallenge bool  `json:"stage"`          // 是否开启章节挑战
	StageIdx       int32 `json:"stageIdx"`       // 章节分身
}

type GatherConfig struct {
	Enabled   bool  `json:"enabled"`   // 是否开启采集
	MinIncome int64 `json:"minIncome"` // 目标收益
}

type RoleConfig struct {
	SeparationIdx *int32 `json:"separationIdx"` // 选择分身
}

type HomelandConfig struct {
	Enabled bool   `json:"enabled"` // 是否开启家园
	Rules   []Rule `json:"rules"`   // 采集规则
}

type Rule struct {
	Enabled bool  `json:"enabled"`
	ItemId  int32 `json:"itemId"`
	MinLv   int32 `json:"minLv"`
}

func (a *AssistantConfig) Validate() error {
	challenge := a.Challenge
	if challenge.StageIdx != 0 && challenge.StageIdx != 1 && challenge.StageIdx != 2 {
		return errors.New("challenge.separationIdx must be 0, 1 or 2")
	}
	if challenge.TowerIdx != 0 && challenge.TowerIdx != 1 && challenge.TowerIdx != 2 {
		return errors.New("challenge.towerIdx must be 0, 1 or 2")
	}
	if challenge.SecretTowerIdx != 0 && challenge.SecretTowerIdx != 1 && challenge.SecretTowerIdx != 2 {
		return errors.New("challenge.secretTowerIdx must be 0, 1 or 2")
	}

	if idx := a.Role.SeparationIdx; idx != nil && *idx != 0 && *idx != 1 && *idx != 2 {
		return errors.New("role.separationIdx must be 0, 1 or 2")
	}

	return nil
}
