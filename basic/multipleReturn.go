package main

import "fmt"

func main()  {
  lat,lng := location("somecity")
  fmt.Println("location is ", lat , lng)


  result := multipleBy100(20)
  fmt.Println("Value is", result)

}

// multiple return values
func location(city string)(string ,string)  {
  return "lat1", "lng1"
}

// named return value
func multipleBy100(x int)(data int)  {
  data = x * 100
  return
}
