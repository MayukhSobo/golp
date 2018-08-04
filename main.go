package main

import (
	"golp/vectorize"
)

func main() {
	r1 := "This pasta is very tasty and affordable"
	r2 := "This pasta is not tasty and is affordable"
	r3 := "This pasta is delicious and cheap"
	r4 := "Pasta is tasty and pasta tastes good"

	// corpus := [][]string{
	// 	[]string{r1},
	// 	[]string{r2},
	// 	[]string{r3},
	// 	[]string{r4},
	// }
	// fmt.Println(corpus)
	vectorize.FindUniquesInCorpus(r1, r2, r3, r4)
}
