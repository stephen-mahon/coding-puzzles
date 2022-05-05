package main

import (
	"fmt"
	"math/rand"
)

func getHitProbability(R int32, C int32, G [][]int32) float64 {

	var val int32
	for i := 0; i < int(R); i++ {
		for j := 0; j < int(C); j++ {
			val += G[i][j]
		}
	}
	return float64(val) / float64(R*C)
}

func main() {
	r, c := 2, 3

	g := make([][]int32, r)
	for i := 0; i < r; i++ {
		g[i] = make([]int32, c)
		for j := 0; j < c; j++ {
			g[i][j] = int32(rand.Intn(2))
		}
	}

	fmt.Println(getHitProbability(int32(r), int32(c), g))

}
