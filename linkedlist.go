package main

import "fmt"

type LinkedList struct {
	First *Node
	Last *Node
	Length int
}

type Node struct {
	Prev *Node
	Data int
	Next *Node
}

func NewList() LinkedList {
	return LinkedList{nil, nil, 0}
}

func NewNode(prev *Node, data int, next *Node) Node {
	return Node{prev, data, next}
}

func main() {
	linkedList := NewList()

	data := 1
	Push(linkedList, data)

	data2 := 2
	Push(linkedList, data2)
}


func Push(linkedList LinkedList, data int) {
	// 1. Node를 생성한다.
	node := NewNode(nil, data, nil)
	// 2. First가 nil이면, 즉 리스트에 값이 비어있다면
	fmt.Println("first", linkedList.First)
	if linkedList.First == nil {
		// 2-1. First와 Last에 Node의 주소값을 넣고 length를 +1한다.
		linkedList.First, linkedList.Last = &node, &node
		linkedList.Length++
		fmt.Println("2", linkedList)
		
	// 3. First가 nil이 아니라면, 즉 리스트에 값이 비어있지 않다면
	} else {
		// 3-1. 새로운 Node의 Next에 First값을 넣는다.
		node.Next = linkedList.First

		// 3-2. First 노드의 Prev에 새로운 Node의 주소값을 넣는다.
		linkedList.First.Prev = &node

		// 3-3. First에 새로운 Node의 주소값을 넣는다.
		linkedList.First = &node

		// 3-4. Length를 +1한다.
		linkedList.Length++

		fmt.Println("3", linkedList)
	}
	fmt.Println("결과", linkedList)
	fmt.Println("결과", &linkedList)
}
