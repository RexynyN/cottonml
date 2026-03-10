package main

import (
	"cottonml/utils"
	"fmt"
	"math"
)

type KNN[D Data, L Label] struct {
	KNeighbors   int
	DistanceFunc string
	Approach     string
}

type knnPoint[T Data, L Label] struct {
	Point *Vector[T]
	Label []L
}

func (k *KNN[D, L]) kdTreeKNNFit(x *Matrix[D], y *Vector[L]) {

}

func (k *KNN[D, L]) rawMajorityVote(labels Vector[D]) any {
	counter := utils.CounterFromSlice(labels)

	result := counter.MostCommonOne()
	return result
}

// Weighted majority vote (Assumes that the labels are sorted by nearest to farthest)
func (k *KNN[D, L]) majorityVote(labels []L) any {
	counter := utils.CounterFromSlice(labels)

	items := counter.MostCommonItems()
	if len(items) == 1 {
		return items[0].Item
	}

	// Retry without the farthest item
	return k.majorityVote(labels[:len(labels)-1])

}

func (k *KNN[D, L]) bruteForceKNNFit(x *Matrix[D], y *Vector[L]) {
	high, highIdx := math.MaxFloat64, 0
	for _, row := range x {
		fmt.Println(high, highIdx)
		fmt.Println(row)
	}
}

func (k *KNN[D, L]) Fit(x *Matrix[D], y *Vector[L]) {
	k.bruteForceKNNFit(x, y)
}

func (k *KNN[D, L]) Predict(x *Matrix[D]) {

}
