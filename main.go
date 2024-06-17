package main

import (
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList(32, 0.6)

	for i := range 100 {
		skipList.Add(i, 0)
	}

	skipList.PrintLeftRight()

	for i := range 100 {
		skipList.Delete(i)
	}

	skipList.PrintLeftRight()

}
