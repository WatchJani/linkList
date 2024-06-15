package main

import (
	"fmt"
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList(32, 0.2)

	skipList.Add(123, 3)
	skipList.Add(124, 3)
	skipList.Add(125, 3)
	skipList.Add(126, 3)
	skipList.Add(127, 3)
	skipList.Add(128, 3)
	skipList.Add(129, 3)
	skipList.Add(130, 3)
	skipList.Add(131, 3)
	skipList.Add(132, 3)
	skipList.Add(133, 3)
	skipList.Add(134, 3)
	skipList.Add(135, 3)
	skipList.Add(136, 3)
	skipList.Add(137, 3)
	skipList.Add(138, 3)
	skipList.Add(139, 3)
	skipList.Add(140, 3)
	skipList.Add(140, 3)
	skipList.Add(141, 3)

	fmt.Println(skipList.CurrentLevel)

	skipList.Print()
}
