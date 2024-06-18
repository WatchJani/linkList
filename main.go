package main

import (
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList[int, int](32, 0.6)

	for range 100 {
		skipList.Add(2, 0)
	}

	skipList.Add(3, 1)

	skipList.PrintLeftRight()

}
