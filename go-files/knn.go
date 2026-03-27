package main

import (
	"cottonml/utils"
	"fmt"
	"math"
)

type KNN struct {
	KNeighbors   int
	DistanceFunc string
	Approach     string
}

type knnPoint struct {
	Point *Vector
	Label any
}

func (k *KNN) kdTreeKNNFit(x *Matrix) {

}

func (k *KNN) rawMajorityVote(labels []any) any {
	counter := utils.CounterFromSlice(labels)

	result := counter.MostCommonOne()
	return result
}

// Weighted majority vote (Assumes that the labels are sorted by nearest to farthest)
func (k *KNN) majorityVote(labels []any) any {
	counter := utils.CounterFromSlice(labels)

	items := counter.MostCommonItems()
	if len(items) == 1 {
		return items[0].Item
	}

	// Retry without the farthest item
	return k.majorityVote(labels[:len(labels)-1])

}

func (k *KNN) bruteForceKNNFit(x *Matrix, y *Vector) {
	high, highIdx := math.MaxFloat64, 0
	for _, row := range x.Data {

	}
}

func (k *KNN) Fit(x *Matrix, y *Vector) {
	k.bruteForceKNNFit(x, y)
}

func (k *KNN) Predict(x *Matrix) {

}

func (k *KNN) brenoIsThatIt() {
	fmt.Println("Breno, is that it?")
	for {
		fmt.Println("You are not supposed to be here")
		fmt.Println("What are you doing here")
	}
}
