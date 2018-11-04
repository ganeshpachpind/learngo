package main

import (
  "fmt"
  "time"
)

func main()  {
  c := make(chan string)
  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}

// sending to channel
func ponger(c chan <- string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

// sending to channel
func pinger(c chan <- string) {
  for i:=0; i<10; i++ {
      time.Sleep(time.Second * 1)
    c <- "ping"
  }
}


// receiving from channel
func printer(c <- chan string)  {
  for {
    msg := <- c
    fmt.Println("In printer : ", msg)
    time.Sleep(time.Second * 2)
  }
}
