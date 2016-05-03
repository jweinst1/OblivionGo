package main

import (
	"fmt"
	"VirtMObl"
)

//Go program that uses functions closure style inside maps

//makes a map and calls a func from it

func main() {
  tester := VirtMObl.CreateVM()
  fmt.Println(tester.Ints["+"](1, 2))
  fmt.Println(tester.Intbools["=="](30, 4))
}
//prints 7

