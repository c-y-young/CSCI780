package main

import "fmt"

var data = []int{2, 7, 11, 15}
var sum = 9

func main() {
	// The time complexity if O(n) because there are two for loops and in each for loop takes linear time

	fmt.Println(data, sum)

	var m = make(map[int]int)

	for i := 0; i < len(data); i++ {
		m[data[i]] = i
	}

	for i := 0; i < len(data); i++ {
		_, ok := m[sum-data[i]]
		if ok {
			fmt.Printf("The indexes of data that sum up to %d is %d and %d\n", sum, i, m[sum-data[i]])
			return
		}
	}

	fmt.Printf("No indexes of data that sum up to %d are found.\n", sum)

	return

}
