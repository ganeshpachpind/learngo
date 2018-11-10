package main

import "fmt"

func main() {

	x := make([]float64, 5) // slice with size 5
	x[0] = 10
	x[1] = 20
	x[2] = 45
	x[3] = 55
	x[4] = 65

	for i := 0; i < 5; i++ {
		fmt.Println("val", x[i])
	}

	y := x[1:3] // [low:high] includes 1,2 index (high-1)
	for j := 0; j < 2; j++ {
		fmt.Println("slice ", y[j])
	}

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
