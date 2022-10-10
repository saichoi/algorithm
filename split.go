package main

import "fmt"

// Arr : 분할에 사용할 변수를 한 곳에 담아 둘 하나의 구조체(타입).
type Arr struct {
	data  []int
	left  int
	right int
	pivot int
}

// NewArr : 구조체(타입)에 맞게 넣을 데이터를 만들어주는 함수.
func NewArr(array []int) Arr {
	return Arr{
		data:  array,
		left:  0, // 가장 첫번째에 있는 값을 왼쪽 포인터로 지정한다.
		right: len(array) - 2, // 피벗의 앞에 있는 값을 오른쪽 포인터로 지정한다.
		pivot: len(array) - 1, // 가장 마지막에 있는 값을 피벗으로 지정한다.
	}
}

func main() {
	// NewArr : 배열을 가지고 구조체에 맞는 데이터를 만들어준다.
	arr := NewArr([]int{0, 5, 2, 1, 6, 3})
	// data = [0, 5, 2, 1, 6, 3]
	// left = 0 왼쪽 포인터의 인덱스 값
	// right = 4 오른쪽 포인터의 인덱스 값
	// pivot = 5 피벗의 인덱스 값

	// 분할 함수를 호출한다.
	arr.Split()
	fmt.Println(arr.data)
}

// Split : *Arr 구조체가 가지고 있는 메서드. go 언어에서는 class가 없어 메서드를 정의 할 때 이와 같은 리시버를 사용하여 정의한다.
// * 를 사용하면 scope 영역 밖에서도 복사된 데이터가 아닌 주소가 같은 동일 데이터를 사용할 수 있다.
func (a *Arr) Split() {
	for {
		// 1. 왼쪽 포인터를 한 칸 씩 오른쪽으로 이동하고 피벗의 값보다 크거나 같을 때 멈춘다.
		a.leftPointerMove()

		// 3. 왼쪽 포인터가 오른쪽 포인터와 같은 값을 가르킬 때 혹은 왼쪽 포인터가 오른쪽 포인터를 넘어서면 분할 반복문을 종료한다.
		if a.left >= a.right {
			break
		}

		// 2. 오른쪽 포인터를 한 칸 씩 왼쪽으로 이동하고 피벗의 값보다 작거나 같을 때 멈춘다.
		a.rightPointerMove()

		// 4. 왼쪽 포인터 이동과 오른쪽 포인터 이동이 멈추면 두 값을 교환한다.
		a.swap(a.left, a.right)
	}
	// 3 or 5. 반복문을 종료하고 왼쪽 포인터가 가르키고 있는 값과 피벗을 교환한다.
	if a.data[a.left] > a.data[a.pivot] {
		a.swap(a.left, a.pivot)
	}
}

// leftPointerMove : 왼쪽 포인터 이동 함수. 계속 오른쪽으로 포인터를 이동(a.left++) 하다가 피벗의 값보다 작은 값을 만날 경우 멈춘다.
func (a *Arr) leftPointerMove() {
	for a.data[a.left] < a.data[a.pivot] {
		a.left++
	}
}

// rightPointerMove : 오른쪽 포인터 이동 함수. 계속 왼쪽으로 포인터를 이동(a.right--) 하다가 피벗의 값보다 큰 값을 만날 경우 멈춘다.
func (a *Arr) rightPointerMove() {
	for a.data[a.right] > a.data[a.pivot] {
		a.right--
	}
}

// swap : 왼쪽 값과 오른쪽 값을 교환하는 함수.
func (a *Arr) swap(left, right int) {
	a.data[left], a.data[right] = a.data[right], a.data[left]
}
