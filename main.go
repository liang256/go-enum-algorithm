package main

import "fmt"

func main() {
	dis := [][]int{
		{0, 5, 3, 2},
		{5, 0, 2, 1},
		{3, 2, 0, 7},
		{2, 1, 7, 0},
	}
	fmt.Println(dis)

	start_id := 0
	route := []int{}
	// set start point is con(0)
	route = append(route, start_id)

	// all available contries
	cons := []int{0, 1, 2, 3}
	// remove start-contry-id in contries
	for i := 0; i < len(cons); i++ {
		if cons[i] == start_id {
			cons[i] = -1
		}
	}
	// fmt.Println(cons)
	enum(route, cons, dis)
}

func enum(route []int, cons []int, dis [][]int) {
	if len(route) >= 4 {
		fmt.Println(route, sum_total(route, dis))
	}

	for i := range cons {
		if cons[i] == -1 {
			continue
		}
		route = append(route, cons[i])
		cons[i] = -1

		enum(route, cons, dis)

		cons[i] = route[len(route)-1]
		route = route[:len(route)-1]
	}
}

func sum_total(route []int, dis [][]int) int {
	sum := 0
	for i := 0; i+1 < len(route); i++ {
		sum += dis[route[i]][route[i+1]]
	}
	return sum
}
