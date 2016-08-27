package main

import "fmt"
import "time"

import "github.com/manawasp/algo-go/sorting"
import "github.com/manawasp/algo-go/utils"

func main() {
	arr := utils.RandArray(30)
	fmt.Println("Initial array is:", arr)
	start := time.Now()
	arr = sorting.Merge(arr)
	elapsed := time.Since(start)
	fmt.Println("Sorted array is:", arr)
	fmt.Printf("Merge sort | values : %d | time : %s\n", len(arr), elapsed)
}
