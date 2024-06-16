package main

import (
	"fmt"
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList(32, 0.2)

	for i := range 20 {
		skipList.Add(i, 3)
	}

	fmt.Println("current level", skipList.CurrentLevel)

	skipList.Print()
}
