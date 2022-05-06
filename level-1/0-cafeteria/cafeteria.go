package main

import "fmt"

func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {
	seats := make([]int, N)

	for i := range S {
		seats[S[i]-1] = 1
	}
	fmt.Println(seats)

	for i := range seats {
		if (int64(i) < int64(len(seats))-K ) && (seats[i+int(K)] != 1)  {
			fmt.Println(seats[i], "yes")
		} else {
			fmt.Println(seats[i], "no")
		}
	}
	return 0
}

func main() {
	var n, k int64
	var m int32
	s := []int64{2, 6}

	n, k = 10, 1
	m = 2

	getMaxAdditionalDinersCount(n, k, m, s)
}
