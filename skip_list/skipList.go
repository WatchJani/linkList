package skipList

import (
	"fmt"
	"math/rand"
	"root/stack"
)

const (
	UpNode    = 0
	RightNode = 1
	DownNode  = 2
	LeftNode  = 3
)

type SkipList struct {
	Root         *Node
	CurrentLevel int
	MaxLevel     int
	Luck         float64
	stack.Stack[*Node]
}

type Node struct {
	Direction []*Node
	Key       int
	Value     int
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:       key,
		Value:     value,
		Direction: make([]*Node, 4),
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

		newNode.Direction[DownNode] = prevuesNode //next node is node up
		prevuesNode.Direction[UpNode] = newNode
		//update left right link
		if node := s.Stack.Pop(); node != nil {
			if key > node.Key {
				rightNode := node.Direction[RightNode]
				newNode.Direction[LeftNode] = node
				newNode.Direction[RightNode] = rightNode

				node.Direction[RightNode] = newNode
				if rightNode != nil {
					rightNode.Direction[LeftNode] = newNode
				}
			} else {
				newNode.Direction[RightNode] = node
				node.Direction[LeftNode] = newNode
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
		for current.Direction[RightNode] != nil && current.Direction[RightNode].Key < key { //stop on left side
			current = current.Direction[RightNode]
		}

		for current.Direction[LeftNode] != nil && current.Key > key {
			current = current.Direction[LeftNode]
		}

		if current.Direction[DownNode] != nil {
			s.Stack.Push(current)
			current = current.Direction[DownNode]
		}
	}

	return current
}

func (s *SkipList) Add(key, value int) {
	zeroLevelNode := NewNode(key, value)
	current := s.SearchInsert(key)

	if current != nil {
		if key > current.Key {
			rightNode := current.Direction[RightNode]      //save right
			zeroLevelNode.Direction[LeftNode] = current    //link left
			zeroLevelNode.Direction[RightNode] = rightNode //link

			current.Direction[RightNode] = zeroLevelNode //link left node with new node
			if rightNode != nil {
				rightNode.Direction[LeftNode] = zeroLevelNode
			}
		} else {
			//just left side
			zeroLevelNode.Direction[RightNode] = current
			current.Direction[LeftNode] = zeroLevelNode
		}
	}

	s.BuildTower(zeroLevelNode)
	s.Flush()
}

// for delete
func (s *SkipList) PrintLeftRight() {
	current := s.Root

	if current == nil {
		return
	}

	for i := s.CurrentLevel; i > 0; i-- {
		for current.Direction[LeftNode] != nil {
			current = current.Direction[LeftNode]
		}

		if current.Direction[DownNode] != nil {
			current = current.Direction[DownNode]
		}
	}

	for current != nil {
		fmt.Println(current.Key)
		current = current.Direction[RightNode]
	}
}

// for delete
func (s *SkipList) PrintRightLeft() {
	current := s.Root

	if current == nil {
		return
	}

	for i := s.CurrentLevel; i > 0; i-- {
		for current.Direction[RightNode] != nil {
			current = current.Direction[RightNode]
		}

		if current.Direction[DownNode] != nil {
			current = current.Direction[DownNode]
		}
	}

	for current != nil {
		fmt.Println(current.Key)
		current = current.Direction[LeftNode]
	}
}

func (s *SkipList) Search(key int) *Node {

	currentNode := s.SearchInsert(key)

	if currentNode.Key == key {
		return currentNode
	}

	if currentNode.Direction[RightNode] != nil && currentNode.Direction[RightNode].Key == key {
		return currentNode.Direction[RightNode]
	}

	return nil
}

func (s *SkipList) Delete(key int) {
	node := s.Search(key)

	if node != nil && s.Root.Key == node.Key {
		for s.Root != nil && s.Root.Direction[LeftNode] == nil && s.Root.Direction[RightNode] == nil {
			s.CurrentLevel--
			s.Root = s.Root.Direction[DownNode]
		}

		if s.Root != nil {
			if s.Root.Direction[LeftNode] != nil {
				s.Root = s.Root.Direction[LeftNode]
			} else if s.Root.Direction[RightNode] != nil {
				s.Root = s.Root.Direction[RightNode]
			}
		}
	}

	for node != nil {
		if node.Direction[LeftNode] != nil {
			node.Direction[LeftNode].Direction[RightNode] = node.Direction[RightNode]
		}

		if node.Direction[RightNode] != nil {
			node.Direction[RightNode].Direction[LeftNode] = node.Direction[LeftNode]
		}

		node.Direction[LeftNode] = nil
		node.Direction[RightNode] = nil

		node = node.Direction[UpNode]
	}
}
