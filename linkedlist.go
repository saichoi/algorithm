package main

import (
	"fmt"
	"errors"
)

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
	linkedList.Push(data)

	fmt.Println(linkedList.First.Data)
	data2 := 2
	linkedList.Push(data2)

	fmt.Println(linkedList.First.Data)

	fmt.Println(linkedList.Get(0))

	fmt.Println(linkedList.Find(1))
	fmt.Println(linkedList.Find(2))
	fmt.Println(linkedList.Find(3))

	fmt.Println(linkedList.Delete())
	fmt.Println(linkedList)
}


func (linkedList *LinkedList) Push(data int) {
	// 1. Node를 생성한다.
	node := NewNode(nil, data, nil)
	// 2. First가 nil이면, 즉 리스트에 값이 비어있다면
	if linkedList.First == nil {
		// 2-1. First와 Last에 Node의 주소값을 넣고 length를 +1한다.
		linkedList.First, linkedList.Last = &node, &node
		linkedList.Length++
		
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

	}
}

func (linkedList *LinkedList) Get(index int) (int, error) {
	// 1. index가 0보다 작거나 Length보다 크면 에러를 반환한다.
	if index < 0 || index > linkedList.Length {
		return 0, errors.New("찾을 수 없습니다.")
	}
	// 2. First부터 index에 도달할 때 까지 Next 를 한다.
	now := linkedList.First
	for i := 0; i < index ; i++ {
		now = now.Next
	}
	// 3. index의 Data를 반환한다.
	return now.Data, nil
}

func (linkedList *LinkedList) Find(target int) bool {
	// 1. Length가 0이면 false를 반환한다.
	if linkedList.Length == 0 {
		return false
	}
	// 2. First부터 Last에 도달할 때 까지 Next 를 반복한다.
	now := linkedList.First
	for i := 0; i < linkedList.Length; i++ {
		// 3. Data와 target값을 비교하여 찾으면 true를 반환한다.
		if now.Data == target {
			return true
		}
		now = now.Next
	}

	// 4. Last 노드까지 도착해도 값이 없다면 false를 반환한다.
	return false
}

func (linkedList *LinkedList) Delete() bool {
	// Length == 0 인 경우
	if linkedList.Length == 0 {
		return true
	}
	// Length == 1 인 경우
	if linkedList.Length == 1 {
		linkedList.First = nil
		linkedList.Last = nil
		linkedList.Length = 0
		return true
	}
	// Length 1 보다 큰 경우
	if linkedList.Length > 1 {
		linkedList.First.Next.Prev = nil
		linkedList.First = linkedList.First.Next
		linkedList.Length--
		return true
	}

	return true
}
