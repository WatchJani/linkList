package skipList

import (
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
	key, value, tower := prevuesNode.Key, prevuesNode.Value, 1 //this is the problem!!!

	for FlipTheCoin(s.Luck) && tower < s.MaxLevel {
		//update next link by horizontally
		newNode := NewNode(key, value)
		newNode.NextNode = prevuesNode

		//update left right link
		if node := s.Stack.Pop(); node != nil {
			rightNode := node.RightLink
			newNode.LeftLink = node
			newNode.RightLink = rightNode
		}

		prevuesNode = newNode
	}

	//update root node
	if s.CurrentLevel < tower {
		s.Root = prevuesNode
		s.CurrentLevel = tower
	}
}

func (s *SkipList) Add(key, value int) error {
	//search -> return left node from search

	//create node
	//update left right side
	//flip coin
	//link to next node from another level

	//create new node
	newNode := NewNode(key, value)

	//
	s.BuildTower(newNode)

	return nil
}
