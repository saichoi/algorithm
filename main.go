package main

import "fmt"

// 버블 정렬 (오름차순)
func AscBubbleSort(array []int) []int {
	// 배열을 비교한 후에 다시 처음부터 비교하기 위한 반복문
	for i := 0; i < len(array)-1; i++ { // 정렬이 다 완료된 시점에서 반복문을 종료하기 위해 -1을 한다.
		// 배열의 값을 비교하기 위한 반복문
		for j := 0; j < len(array)-1; j++ { // 첫번째 비교하는 배열의 인덱스가 전체 배열의 길이의 -1이다.
			// 왼쪽에 있는 숫자가 오른쪽에 있는 숫자보다 크다면 서로 자리를 교환한다.
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
	return array
}

// 버블 정렬 (내림차순)
func DescBubbleSort(array []int) []int {
	// 배열을 비교한 후에 다시 처음부터 비교하기 위한 반복문
	for i := 0; i < len(array)-1; i++ {
		// 배열의 값을 비교하기 위한 반복문
		for j := 0; j < len(array)-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
	return array
}

// 버블 정렬 (오름차순) 2
func AscBubbleSort2(array []int) []int {
	swapped := true
	// 배열의 값 비교가 끝날때까지 도는 반복문
	for swapped {
		swapped = false
		// 배열의 값을 비교하기 위한 반복문
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
		}
	}
	fmt.Println(array)
	return array
}

// 선택 정렬 (오름차순)
func selectedSort(array []int) []int {
	count := 0
	for i := 0; i < len(array)-1; i++ {
		minIndex := i
		// 배열의 최소값을 찾기 위한 반복문
		for j := i; j < len(array); j++ {
			// 배열의 값을 비교할 때 값이 작은 쪽과 교환한다.
			if array[j] < array[minIndex] {
				minIndex = j
				count++
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
	fmt.Println(array)
	fmt.Println(count)
	return array
}

func main() {
	array1 := []int{5, 4, 3, 2, 1, 0}
	array2 := []int{0, 1, 2, 3, 4, 5}
	array3 := []int{5, 4, 3, 2, 1, 0}
	array4 := []int{5, 4, 3, 2, 1, 0}
	AscBubbleSort(array1)
	DescBubbleSort(array2)
	AscBubbleSort2(array3)
	selectedSort(array4)
}
