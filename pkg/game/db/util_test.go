package db

import (
	"fmt"
	"testing"
)

func TestParseReward(t *testing.T) {
	rewards := ParseRewards("100003=890")
	fmt.Println(rewards)
}
