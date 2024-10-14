package db

import (
	"fmt"
	"strconv"
	"strings"
)

type Reward struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Num  int32  `json:"num"`
}

func (r Reward) String() string {
	return fmt.Sprintf("%s(%d)=%d", r.Name, r.Id, r.Num)
}

type Rewards []*Reward

func (r Rewards) String() string {
	parts := make([]string, 0)
	for _, reward := range r {
		parts = append(parts, reward.String())
	}
	return strings.Join(parts, "|")
}

func ParseRewards(reward string) Rewards {
	if reward == "" {
		return nil
	}

	rewards := make(Rewards, 0)
	for _, reward := range strings.Split(reward, "|") {
		parts := strings.Split(reward, "=")
		id, _ := strconv.Atoi(parts[0])
		num, _ := strconv.Atoi(parts[1])
		rewards = append(rewards, &Reward{
			Id:   int32(id),
			Name: GetItemName(int32(id)),
			Num:  int32(num),
		})
	}

	return rewards
}
