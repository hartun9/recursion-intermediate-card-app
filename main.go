package main

import (
	"fmt"

	"github.com/hartun9/recursion-intermediate-card-app/internal/helper"
)

func main() {
	arr1 := []int{1, 9, 19, 3, 4, 6}
	fmt.Println(helper.MaxInArrayIndex(arr1))

	arr2 := []int{5, 2, 1, 3, 5, 5}
	fmt.Println(helper.MaxInArrayIndex(arr2))
}
