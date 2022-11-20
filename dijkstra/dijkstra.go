package main

import (
	"container/list"
	"fmt"
)

type City struct {
	Name   string
	Routes map[*City]int
}

func NewCity(name string) City {
	return City{
		Name:   name,
		Routes: make(map[*City]int),
	}
}

func (c *City) AddRoute(city *City, price int) {
	c.Routes[city] = price
}

func main() {
	var tokyo, sapporo, osaka, fukuoka, okinawa City

	tokyo = NewCity("도쿄")
	tokyo.AddRoute(&osaka, 40)
	tokyo.AddRoute(&sapporo, 100)

	sapporo = NewCity("삿뽀로")
	sapporo.AddRoute(&okinawa, 180)

	osaka = NewCity("오사카")
	osaka.AddRoute(&fukuoka, 80)
	osaka.AddRoute(&sapporo, 80)

	fukuoka = NewCity("후쿠오카")
	fukuoka.AddRoute(&okinawa, 40)
	fukuoka.AddRoute(&sapporo, 120)

	okinawa = NewCity("오키나와")
	okinawa.AddRoute(&osaka, 180)
	okinawa.AddRoute(&tokyo, 160)

	fmt.Println(lcc(&tokyo, &okinawa))
}

// 출발 도시에서 도착 도시까지의 최저가를 구하는 메서드
func lcc(src *City, desc *City) int {
	// 도시와 해당도시까지의 최저 비용을 담을 변수
	cost := map[*City]int{}
	// 방문한 정점을 체크하기 위한 변수
	visited := map[*City]bool{}
	// 아직 방문하지 않은 도시를 담은 변수(스택)
	stack := list.New()

	// 출발지의 연결된 도시의 이름과 가격을 최저가 담아두는 변수에 넣는다.
	//for city, price := range src.Routes {
	//	cost[city] = price
	//	stack.PushBack(city) // 앞으로 방문할 도시를 스택에 쌓는다.
	//}
	stack.PushBack(src)

	// 방문할 도시가 사라진 때까지 반복
	for stack.Len() > 0 {
		current := stack.Remove(stack.Back()).(*City)

		// 이미 방문한 경우
		if visited[current] {
			continue
		}

		// 최저가인 경우 비용을 cost에 기록한다.
		for key, val := range current.Routes {
			newPrice := cost[current] + val // 정점 도시까지의 최저가와 근접 도시의 가격을 더한다.
			// 이미 기록된 가격과 근접도시 가격
			if cost[key] > newPrice || cost[key] == 0 {
				// 이미 기록된 최저가와 새로운 가격을 비교해서 새로운 가격이 적다면 대체하고 값이 0이라면 그냥 넣어준다.(초기값이 0)
				cost[key] = newPrice
			}
			// 근접도시를 stack에 쌓는다.
			stack.PushBack(key)
		}
		// visited를 true로 변경한다.
		visited[current] = true
	}

	fmt.Println(visited)
	for key, val := range cost {
		fmt.Println(key, val)
	}

	// 최저가를 담아둔 cost에서 도착지(desc)의 value값을 리턴한다.
	return cost[desc]

}
