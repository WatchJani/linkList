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
		tower++
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
		for current.RightLink != nil && current.Key < key {
			current = current.RightLink
		}
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

// search
// insert new node(search *node)()
// need i update
// update ()(new *node)
func (s *SkipList) Add(key, value int) {
	// Search for the position to insert the new node.
	current := s.Search(key)

	//Create new Node
	zeroLevelNode := NewNode(key, value)

	//update if exist
	if current != nil {
		rightNode := current.RightLink
		zeroLevelNode.LeftLink = current
		zeroLevelNode.RightLink = rightNode
	}

	s.BuildTower(zeroLevelNode)
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

	for current != nil {
		fmt.Println(current.Key)
		current = current.RightLink
	}
}
