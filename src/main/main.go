package main

import (
	"fmt"
	"intobjects"
)

//Go program that uses functions closure style inside maps

//makes a map and calls a func from it

func main() {
  tester := intobjects.CreateIntList()
  tester.Push(5)
  tester.Push(8)
  tester.Pushleft(1)
  fmt.Println(tester)
  tester.Set(1, 88)
  fmt.Println(tester)
}
//prints 7

