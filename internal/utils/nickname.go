package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var adjectives = []string{"Brave", "Silent", "Wild", "Clever", "Happy"}
var nouns = []string{"Fox", "Eagle", "Wolf", "Tiger", "Panda"}

func GenerateNickname() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%s%d", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))], rand.Intn(1000))
}
