package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, value := range vals {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(10, 20))

	slice := []int{1, 2, 3, 4}
	fmt.Println(sum(slice...))
}
