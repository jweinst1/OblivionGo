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
  tester.Push(9)
  tester.Push(10)
  tester.Push(11)
  tester.Push(12)
  fmt.Println(tester.Getfirst())
}
//prints 7

