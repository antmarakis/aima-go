package main

import (
	"strings"
)

type Item struct {
	features []float64
	class    string
}

/*-------------------------------------------------------*/
//Plurality Learner
func MostPopular(items []Item) string {
	/*Find class that appears most often in items*/
	m := make(map[string]int)

	for _, i := range items {
		class := i.class
		if KeyInMap(class, m) {
			m[class] += 1
		} else {
			m[class] = 1
		}
	}

	return FindMaxInMap(m)
}

func PluralityLearner(dataset []Item) func([]float64) string {
	/*Always predict the most popular class in the dataset*/
	mostPopular := MostPopular(dataset)
	return func(_ []float64) string {
		return mostPopular
	}
}

/*-------------------------------------------------------*/
// k-Nearest Neighbors
func NearestNeighborLearner(dataset []Item, k int) func([]float64) string {
	return func(example []float64) string {
		itemHeap := InitMinHeap()
		for _, item := range dataset {
			dis := euclidean(example, item.features)
			itemHeap.HeapPush(dis, item.class)
		}

		nearest := itemHeap.SmallestN(k)
		class := FindMaxInMap(Occurences(nearest))
		return strings.TrimSpace(class)
	}
}
