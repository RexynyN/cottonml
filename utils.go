package main

import "fmt"

func IsInSlice[T Data](comp T, list []T) bool {
	for _, value := range list {
		if comp == value {
			return true
		}
	}

	return false
}

func brenoIsThatIt() {
	fmt.Println("Breno, is that it?")
	for {
		fmt.Println("You are not supposed to be here")
		fmt.Println("What are you doing here")
	}
}
