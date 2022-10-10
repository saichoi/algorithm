package main

import "fmt"

func main() {
	arr := []int{0, 5, 2, 1, 6, 3}
	quicksort(arr)
	fmt.Println(arr)
}

func split(arr []int) int {
	pivot := len(arr) - 1
	left, right := 0, len(arr) - 2

	for {
		// 1. 왼쪽 포인터를 피벗의 값보다 작을 때까지 오른쪽으로 이동한다. 
		for arr[left] < arr[pivot] { 
			left++ 
		}
		// 3. 왼쪽 포인터와 오른쪽 포인터가 같은 값을 가르키거나 넘어설 경우
		if left >= right { 
			break
		}
		// 2. 오른쪽 포인터를 피벗의 값보다 클 때까지 왼족으로 이동한다.
		for arr[right] > arr[pivot] { 
			right-- 
		}
		// 4. 왼쪽 포인터와 오른쪽 포인터의 이동이 멈추면 두 값을 교환한다.
		arr[left], arr[right] = arr[right], arr[left]
	}
	// 3. 왼쪽 값과 피벗의 값을 교환한다.
	if arr[left] > arr[pivot] {
		arr[left], arr[pivot] = arr[pivot], arr[left]
	}
	return left
}

func quicksort(arr []int) {
	// 기저 조건 : 분할된 배열의 원소 개수가 0개 또는 1개일 경우
	if len(arr) > 1 {
		// 대상 배열을 분할을 한 후에 피벗의 인덱스를 가져온다.
		pivot := split(arr)
		// 피벗의 왼쪽에 있는 배열을 분할 재귀
		quicksort(arr[0:pivot]) // [ 0 1 2 ]
		// 피벗의 오른쪽에 있는 배열을 분할 재귀
		quicksort(arr[pivot+1:]) // [ 6 5 ]
	}
}
