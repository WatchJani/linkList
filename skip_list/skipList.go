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
	stack.Stack[*Node]
}

type Node struct {
	LeftLink  *Node
	RightLink *Node
	UpNode    *Node
	DownNode  *Node
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

		newNode.DownNode = prevuesNode //next node is node up
		prevuesNode.UpNode = newNode
		//update left right link
		if node := s.Stack.Pop(); node != nil {
			if key > node.Key {
				rightNode := node.RightLink
				newNode.LeftLink = node
				newNode.RightLink = rightNode

				node.RightLink = newNode
				if rightNode != nil {
					rightNode.LeftLink = newNode
				}
			} else {
				newNode.RightLink = node
				node.LeftLink = newNode
			}
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

func (n *Node) Update(key, value int) {
	n.Key = key
	n.Value = value
}

// left <-> right | search
func (s *SkipList) SearchInsert(key int) *Node {
	current := s.Root

	for i := s.CurrentLevel; i > 0; i-- {
		for current.RightLink != nil && current.RightLink.Key < key { //stop on left side
			current = current.RightLink
		}

		for current.LeftLink != nil && current.Key > key {
			current = current.LeftLink
		}

		if current.DownNode != nil {
			s.Stack.Push(current)
			current = current.DownNode
		}
	}

	return current
}

func (s *SkipList) Add(key, value int) {
	zeroLevelNode := NewNode(key, value)
	current := s.SearchInsert(key)

	if current != nil {
		if current.Key == key {
			for current != nil {
				current.Value = value
				current = current.UpNode
			}
			return
		}

		if current.RightLink != nil && current.RightLink.Key == key {
			for current.RightLink != nil {
				current.RightLink.Value = value
				current.RightLink = current.RightLink.UpNode
			}
			return
		}

		if key > current.Key {
			rightNode := current.RightLink      //save right
			zeroLevelNode.LeftLink = current    //link left
			zeroLevelNode.RightLink = rightNode //link

			current.RightLink = zeroLevelNode //link left node with new node
			if rightNode != nil {
				rightNode.LeftLink = zeroLevelNode
			}
		} else {
			//just left side
			zeroLevelNode.RightLink = current
			current.LeftLink = zeroLevelNode
		}
	}

	s.BuildTower(zeroLevelNode)
	s.Flush()
}

func (s *SkipList) Search(key int) *Node {
	currentNode := s.SearchInsert(key)

	if currentNode != nil && currentNode.Key == key {
		return currentNode
	}

	if currentNode.RightLink != nil && currentNode.RightLink.Key == key {
		return currentNode.RightLink
	}

	return nil
}

func (s *SkipList) Delete(key int) {
	node := s.Search(key)

	if node != nil && s.Root.Key == node.Key {
		for s.Root != nil && s.Root.LeftLink == nil && s.Root.RightLink == nil {
			s.CurrentLevel--
			s.Root = s.Root.DownNode
		}

		if s.Root != nil {
			if s.Root.LeftLink != nil {
				s.Root = s.Root.LeftLink
			} else if s.Root.RightLink != nil {
				s.Root = s.Root.RightLink
			}
		}
	}

	for node != nil {
		if node.LeftLink != nil {
			node.LeftLink.RightLink = node.RightLink
		}

		if node.RightLink != nil {
			node.RightLink.LeftLink = node.LeftLink
		}

		node.LeftLink = nil
		node.RightLink = nil

		node = node.UpNode
	}
}

func (s *SkipList) PrintLeftRight() {
	current := s.Root

	if current == nil {
		return
	}

	for i := s.CurrentLevel; i > 0; i-- {
		for current.LeftLink != nil {
			current = current.LeftLink
		}

		if current.DownNode != nil {
			current = current.DownNode
		}
	}

	for current != nil {
		fmt.Println(current.Key)
		current = current.RightLink
	}
}

func (s *SkipList) PrintRightLeft() {
	current := s.Root

	if current == nil {
		return
	}

	for i := s.CurrentLevel; i > 0; i-- {
		for current.RightLink != nil {
			current = current.RightLink
		}

		if current.DownNode != nil {
			current = current.DownNode
		}
	}

	for current != nil {
		fmt.Println(current.Key)
		current = current.LeftLink
	}
}
