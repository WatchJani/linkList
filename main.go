package main

import (
	"fmt"
	"math/rand"
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList(32, 0.2)

	for range 200 {
		skipList.Add(rand.Intn(5000000), 3)
	}

	fmt.Println("current level", skipList.CurrentLevel)

	skipList.Print()
}
