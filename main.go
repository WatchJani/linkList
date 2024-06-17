package main

import (
	"fmt"
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList(32, 0.6)

	fmt.Println(skipList.Root)

	//insert
	for i := 0; i < 100; i++ {
		skipList.Add(i, 3)
	}

	for i := 0; i < 100; i++ {
		skipList.Delete(i)
	}

	skipList.Add(2, 1)
	skipList.Add(15, 1)

	fmt.Println(skipList.Search(2))
}
