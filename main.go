package main

import "fmt"

func main() {
	cache := make(map[int]int)
	var nums [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&nums[i])
	}
	for i := 0; i < 10; i++ {
		if result, ok := cache[nums[i]]; ok {
			fmt.Print(result)
		} else {

		}
	}

}
