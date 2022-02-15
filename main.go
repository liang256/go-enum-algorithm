package main

import "fmt"

func main() {
	// distance of each contries
	dis := [][]int{
		{0, 14, 15, 17, 16},
		{14, 0, 24, 8, 36},
		{15, 24, 0, 34, 4},
		{17, 8, 34, 0, 30},
		{16, 36, 4, 30, 0},
	}
	fmt.Println(dis)

	// choose a start point
	start_id := 0
	// init route slice
	route := []int{}
	// set start point is con(0)
	route = append(route, start_id)

	// all available contrie-ids
	cons := []int{0, 1, 2, 3, 4}
	// remove start-contry-id in contries since it
	// is already uesed as the start point
	for i := 0; i < len(cons); i++ {
		if cons[i] == start_id {
			cons[i] = -1
		}
	}
	// run the algorithm
	enum(route, cons, dis)
}

func enum(route []int, cons []int, dis [][]int) {
	if len(route) >= 5 {
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
