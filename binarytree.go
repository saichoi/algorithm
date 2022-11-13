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
	fmt.Println(tree.Search(33))
	//fmt.Println(tree.Insert(77))
	//fmt.Println(tree.Search(77))
	//fmt.Println(tree.Insert(33))
	//
	//tree2 := BianryTree{}
	//fmt.Println(tree2.Insert(45))
	//fmt.Println(tree2.Root.Data)
}

func (b *BianryTree) Search(target int) (*Node, *Node) {

	n := b.Root
	return search(nil, n, target)
}

func search(parentNode *Node, node *Node, target int) (*Node, *Node) { // 부모랑 타켓값을 반환하는 검색

	// 1. 루트 노드의 값을 확인하여 타켓과 값이 같다면 true를 반환한다.
	if node.Data == target {
		return parentNode, node
	}

	// 2. 타켓이 루트보다 작으면 left를 확인, 크면 right를 확인한다.
	if node.Data < target {
		if node.Right == nil {
			return nil, nil
		}
		return search(node, node.Right, target)
	} else {
		if node.Left == nil {
			return nil, nil
		}
		return search(node, node.Left, target)
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

func (b *BianryTree) Delete(target int) error {
	parent, child := b.Search(target) // 삭제하고자 하는 노드와 부모 노드

	// search 했는데 타켓이 없는 경우
	if child == nil {
		return errors.New("삭제하려는 값이 없습니다.")
	}

	// case 1. 타겟의 자식이 0개인 경우
	if child.Left == nil && child.Right == nil {
		//if parent.Data < child.Data { // 부모보다 자식이 크다면 오른쪽 자식
		//	parent.Right = nil
		//} else { // 부모보다 자식이 크다면 왼쪽 자식
		//	parent.Left = nil
		//}

		// target이 root에 있다면
		if parent == nil {
			// 루트를 지워준다.
			b.Root = nil

		  // 자식의 값이 부모의 left에 있으면 부모의 left를 nil로 변경해주고
		} else if parent.Left == child {
			parent.Left = nil

		  // 자식의 값이 부모의 right에 있으면 부모의 rihgt를 nil로 변경해준다.
		} else if parent.Right == child {
			parent.Right = nil
		}
	// case 2. 자식이 둘 다 있는 경우
	} else if child.Left != nil && child.Right !=nil {
		// 1. 후속 노드를 찾는다.
		subNodeParent, subNode := child, child.Right
		for {
			if subNode.Left == nil {
				break
			}
			subNodeParent, subNode = subNode, subNode.Left
		}
		// 2. 후속 노드를 타겟 위치에 넣는다.
		// 2-1. target이 부모보다 작으면 부모의 왼쪽, 크면 부모의 오른쪽에 후속 값을 넣는다.

		// root에 target이 있다면
		if parent == nil {
			// root에 후속 노드를 넣어준다.
			b.Root = subNode
		} else if parent.Data > child.Data {
			parent.Left = subNode
		} else if parent.Data < child.Data {
			parent.Right = subNode
		}
		subNodeRight := subNode.Right // 없을때 nil, 있을때 주소
		// 2-2. 후속값의 left에 target의 left값을 넣는다.
		subNode.Left = child.Left
		// 2-3. 후속값의 right에 target의 right값을 넣는다.
		subNode.Right = child.Right
		// 2-4. 후속값을 삭제한다.

		subNodeParent.Left = subNodeRight
		//if subNodeRight != nil {
		//	// 후속 노드에 자식이 있다면 후속 노드 위치에 후속 노드의 자식을 넣는다.
		//	subNodeParent.Left = subNodeRight
		//} else {
		//	// subNodeRight == nil
		//	subNodeParent.Left = nil
		//}

	// case 3. 자식이 하나밖에 없는 경우(오른쪽 혹은 왼쪽)
	} else {
		// root가 target이라면
		if parent == nil {
			// left에 값이 있다면 root의 left 값을 root에 넣어주고
			if child.Left != nil {
				b.Root = child.Left

			// right에 값이 있다면 root의 right 값을 root에 넣어준다.
			} else {
				b.Root = child.Right
			}
		  // 자식의 값이 부모의 left에 있으면
		} else if parent.Left == child {
			// 자식의 왼쪽(자식의 자식)에 값이 있다면 부모의 왼쪽 값(자식이 있던 곳)에 자식의 왼쪽(자식의 자식) 값을 넣어준다.
			if child.Left != nil {
				parent.Left = child.Left
			// 자식의 오른쪽(자식의 자식)에 값이 있다면 부모의 왼쪽 값(자식이 있던 곳)에 자식의 오른쪽(자식의 자식) 값을 넣어준다.
			} else {
				parent.Left = child.Right
			}
		// 자식의 값이 부모의 right에 있으면 
		} else if parent.Right == child {
			// 자식의 왼쪽(자식의 자식)에 값이 있으면 부모의 오른쪽 값(자식이 있던 곳)에 자식의 왼쪽 값(자식의 자식) 값을 넣어준다.
			if child.Left != nil {
				parent.Right = child.Left
			// 자식의 오른쪽(자식의 자식)에 값이 있으면 부모의 오른쪽 값(자식이 있던 곳)에 자식의 오른쪽 값(자식의 자식) 값을 넣어준다.
			} else {
				parent.Right = child.Right
			}
		}
		
	}
	return nil
}

