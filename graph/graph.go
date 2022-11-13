package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	Visited bool
	Name    string
	Friends []*Graph
}

func main() {
	var g1, g2, g3, g4, g5, g6, g7, g8 Graph

	g1 = Graph{
		Name:    "상인",
		Friends: []*Graph{&g2, &g3, &g4},
	}
	g2 = Graph{
		Name:    "은지",
		Friends: []*Graph{&g1},
	}
	g3 = Graph{
		Name:    "범기",
		Friends: []*Graph{&g1, &g5},
	}
	g4 = Graph{
		Name:    "혜진",
		Friends: []*Graph{&g1, &g6},
	}
	g5 = Graph{
		Name:    "영욱",
		Friends: []*Graph{&g2, &g7},
	}
	g6 = Graph{
		Name:    "윤희",
		Friends: []*Graph{&g4},
	}
	g7 = Graph{
		Name:    "상미",
		Friends: []*Graph{&g5, &g8},
	}
	g8 = Graph{
		Name:    "수현",
		Friends: []*Graph{&g7},
	}
	//fmt.Println(g8.Friends[0].Name)
	g1.Bfs()
}

func (g *Graph) Bfs() {
	// 방문할때 이름 출력
	// 1. 정점과 인접한 각 정점을 방문하고 방문한 정점이면 표시하고 큐에 쌓는다.
	queue := list.New()
	g.Visited = true
	for _, friend := range g.Friends {
		friend.Visited = true
		queue.PushBack(friend)
	}
	fmt.Println(g.Name)

	// 2. 인접 정점을 모두 방문한 경우, 큐로 부터 값을 꺼내 1번을 반복한다.
	for queue.Len() != 0 {
		breadth := queue.Remove(queue.Front()).(*Graph)
		//fmt.Println("****** 현재 ****** ", breadth.Name)
		breadth.Visited = true
		for _, friend := range breadth.Friends {
			if friend.Visited {
				//fmt.Println("이미 방문함 tachi ===== ", friend.Name)
				continue
			}
			//fmt.Println("방문 안한 친구 tachi ===== ", friend.Name)
			friend.Visited = true
			queue.PushBack(friend)
		}
		fmt.Println(breadth.Name)
	}

	// 3. 모든 정점에 방문 표시를 했고 큐가 비었다면 종료한다.

}
