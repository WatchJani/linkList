package skipList

import (
	"fmt"
	"math/rand"
	"root/stack"
)

type SkipList struct {
	Root         *Node
	CurrentLevel int
	MaxLevel     int
	Luck         float64
	*stack.Stack[*Node]
}

type Node struct {
	LeftLink  *Node
	RightLink *Node
	NextNode  *Node
	Key       int
	Value     int
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

func NewSkipList(maxLevel int, luck float64) *SkipList {
	return &SkipList{
		MaxLevel: maxLevel,
		Luck:     luck,
		Stack:    stack.NewStack[*Node](maxLevel),
	}
}

func FlipTheCoin(luck float64) bool {
	return luck > rand.Float64()
}

func (s *SkipList) BuildTower(prevuesNode *Node) {
	_, _, tower := prevuesNode.Key, prevuesNode.Value, 1 //this is the problem!!!

	for FlipTheCoin(s.Luck) && tower < s.MaxLevel {
		//update next link by horizontally
		// newNode := NewNode(key, value)
		// newNode.NextNode = prevuesNode //next node is node up

		// //update left right link
		// if node := s.Stack.Pop(); node != nil {
		// 	rightNode := node.RightLink
		// 	newNode.LeftLink = node
		// 	newNode.RightLink = rightNode

		// 	node.RightLink = newNode
		// 	if rightNode != nil {
		// 		rightNode.LeftLink = newNode
		// 	}
		// }

		// prevuesNode = newNode
		// tower++
	}

	//update root node
	if s.CurrentLevel < tower {
		s.Root = prevuesNode
		s.CurrentLevel = tower
	}
}

// left <-> right | search
func (s *SkipList) Search(key int) *Node {
	current := s.Root

	for i := s.CurrentLevel; i > 0; i-- {
		for current.RightLink != nil && current.RightLink.Key < key { //stop on left side
			current = current.RightLink
		}

		//what if we have one element(root) and new element is lower then root
		for current.LeftLink != nil && current.Key > key {
			current = current.LeftLink
		}

		if current.NextNode != nil {
			current = current.NextNode
			s.Stack.Push(current)
		}
	}

	return current
}

func (s *SkipList) Add(key, value int) {
	zeroLevelNode := NewNode(key, value)
	current := s.Search(key)

	if current != nil {
		if key > current.Key {
			rightNode := current.RightLink      //save right
			zeroLevelNode.LeftLink = current    //link left
			zeroLevelNode.RightLink = rightNode //link

			current.RightLink = zeroLevelNode //link left node with new node
			if rightNode != nil {
				rightNode.LeftLink = zeroLevelNode
			}
		} else {
			zeroLevelNode.RightLink = current
			current.LeftLink = zeroLevelNode
		}
	}

	s.BuildTower(zeroLevelNode)
	s.Flush()
}

func (s *SkipList) Print() {
	current := s.Root

	for i := s.CurrentLevel; i > 0; i-- {
		for current.LeftLink != nil {
			current = current.LeftLink
		}

		if current.NextNode != nil {
			current = current.NextNode
		}
	}

	var counter int
	for current != nil {
		counter++
		fmt.Println(current.Key)
		current = current.RightLink
	}

	fmt.Println(counter)
}

//broj gresaka je jednak broju popa unutar builda
