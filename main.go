package main

import (
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList(32, 0.2)

	skipList.Add(123, 3)
	skipList.Add(125, 3)

	skipList.Print()
}
