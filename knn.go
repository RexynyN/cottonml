package main

import "cottonml/utils"

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

func rawMajorityVote(labels []any) any {
	counter := utils.CounterFromSlice(labels)

	result := counter.MostCommonOne()
	return result
}

// Weighted majority vote (Assumes that the labels are sorted by nearest to farthest)
func majorityVote(labels []any) any {
	counter := utils.CounterFromSlice(labels)

	items := counter.MostCommonItems()
	if len(items) == 1 {
		return items[0].Item
	}

	// Retry without the farthest item
	return majorityVote(labels[:len(labels)-1])

}

func (k *KNN) bruteForceKNNFit(x *Matrix, y *Vector) {
}

func (k *KNN) Fit(x *Matrix, y *Vector) {

}

func (k *KNN) Predict(x *Matrix) {

}
