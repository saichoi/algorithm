package main

import "fmt"
import "errors"

type Node struct {
	Left *Node
	Data int
	Right *Node
}

type BianryTree struct {
	Root  *Node
}

func main() {
	var n1, n2, n3, n4, n5, n6, n7 Node
	n1 = Node{ Left: &n2, Data: 50, Right: &n3 }
	n2 = Node{ Left: &n4, Data: 25, Right: &n5 }
	n3 = Node{ Left: &n6, Data: 75, Right: &n7 }
	n4 = Node{ Data: 10 }
	n5 = Node{ Data: 33 }
	n6 = Node{ Data: 56 }
	n7 = Node{ Data: 89 }

	tree := BianryTree{Root: &n1}
	fmt.Println(tree.Search(77))
	fmt.Println(tree.Insert(77))
	fmt.Println(tree.Search(77))
	fmt.Println(tree.Insert(33))

	tree2 := BianryTree{}
	fmt.Println(tree2.Insert(45))
	fmt.Println(tree2.Root.Data)
}

func (b *BianryTree) Search(target int) bool {

	n := b.Root
	return search(n, target)
}

func search(n *Node, target int) bool {

	// 1. 루트 노드의 값을 확인하여 타켓과 값이 같다면 true를 반환한다.
	if n.Data == target {
		return true
	}

	// 2. 타켓이 루트보다 작으면 left를 확인, 크면 right를 확인한다.
	if n.Data < target {
		if n.Right == nil {
			return false
		}
		return search(n.Right, target)
	} else {
		if n.Left == nil {
			return false
		}
		return search(n.Left, target)
	}
}

func (b *BianryTree) Insert(data int) error {
	// 1. 루트가 비어 있다면 새로운 노드를 삽입하고 종료
	newVal := Node{Data: data}
	if b.Root == nil{
		b.Root = &newVal
		return nil
	}

	n := b.Root
	// 2. 트리에서 자식이 없는 노드가지 찾아간다.
	for {
		// 2-1. 값이 이미 존재하는 경우 에러 처리를 한다.
		if n.Data == data {
			return errors.New("중복된 값입니다.")
		}

		// 2-2. 현재 노드와 값을 비교하여 크면 right 트리를 확인하고 값이 없다면 넣어준다.
		if n.Data < data {
			if n.Right == nil {
				n.Right = &newVal
				return nil
			}
			n = n.Right
		// 2-3. 현재 노드와 값을 비교하여 작으면 left 트리를 확인하고 값이 없다면 넣어준다.
		} else {
			if n.Left == nil {
				n.Left = &newVal
				return nil
			}
			n = n.Left
		}
	}
}
