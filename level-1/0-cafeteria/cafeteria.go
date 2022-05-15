package main

import "fmt"

func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {
	seats := make([]int, N)
	var tally int64

	for j := range S {
		if j == 0 {
			for i := int(S[j]) - 1; i >= 0; i -= int(M) {
				seats[i] = 1
				tally++
			}
		}
		if j == len(S)-1 {
			for i := int(S[len(S)-1]) - 1; i < len(seats); i += int(M) {
				seats[i] = 1
				tally++
			}
		}
		fmt.Println(j)
		//for i := int(S[j]); i < int(S[j+1]); i += int(M) {
		//	seats[i] = 1
		//	tally++
		//}

	}
	fmt.Println(seats)
	return tally - int64(len(S))
}

func main() {
	var n, k int64
	var m int32
	s := []int64{2, 6}

	n, k = 10, 1
	m = 2

	x := getMaxAdditionalDinersCount(n, k, m, s)
	fmt.Println(x)
}
