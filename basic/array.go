package main

import "fmt"

func main()  {
  arr:=[4]int{1,4,6,7}

  for i:=0 ;i< len(arr); i++ {
    fmt.Println("val", arr[i])
  }
}
