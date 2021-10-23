package entity

import (
	"fmt"
	"strings"
)

type Subscription []interface{}

func NewSubscription(id int64) Subscription {
	subscription := make(Subscription, 3)
	mining := make([][]string, 2)
	mining[0] = []string{
		"mining.set_difficulty",
		fmt.Sprintf("luxorminer_%d", id),
	}
	mining[1] = []string{
		"mining.notify",
		fmt.Sprintf("miner_%d", id),
	}
	subscription[0] = mining
	subscription[1] = strings.ToUpper(fmt.Sprintf("%07x", id))
	subscription[2] = 4
	return subscription
}
