package main

import (
	skipList "root/skip_list"
)

func main() {
	skipList := skipList.NewSkipList(32, 0.4)

	//insert
	for i := 0; i < 100; i++ {
		skipList.Add(i, 3)
	}

	// current := skipList.Root

	// for i := skipList.CurrentLevel; i > 0; i-- {
	// 	for current.LeftLink != nil {
	// 		current = current.LeftLink
	// 	}

	// 	if current.NextNode != nil {
	// 		current = current.NextNode
	// 	}
	// }

	// fmt.Println(current.RightLink)

	// skipList.PrintRightLeft()
	// fmt.Println("========")

	skipList.PrintRightLeft()
	// //delete
	for i := 0; i < 100; i++ {
		skipList.Delete(i)
	}

	skipList.Add(1, 3)
}
