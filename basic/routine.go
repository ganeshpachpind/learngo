package main

import "fmt"

func print(j int)  {
  for  i:=0; i<10 ;i++ {
    fmt.Println("IN print", j)
  }
}


func main()  {
  fmt.Println("Executing main function")
  for  i:=0; i<10 ;i++ {
      go print(i)
  }
  var input string
  fmt.Scanln(&input)
}
