package main

import "fmt"

func main() {
	arr := [4]int{1, 4, 6, 7}

	for i := 0; i < len(arr); i++ {
		fmt.Println("val", arr[i])
	}

	data := [5]float64{10, 20, 30, 40, 50.4}

	var res float64
	for i := 0; i < 5; i++ {
		res = res + data[i]
	}

	fmt.Println("Result is ", res)
}
