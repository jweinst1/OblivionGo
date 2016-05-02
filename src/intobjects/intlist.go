package intobjects
//file to implement expanding/shrinking intlist with a map
//These work entirely by inplace operations
import (
	"strconv"
)
type IntList struct {
	list map[int]int
	//for keeping track of placement indices and insertion spots
	record map[string]int
}

//creates an IntList and returns it
func CreateIntList() IntList {
	lst := IntList{list:make(map[int]int), record:make(map[string]int)}
	lst.record["insert"] = 0
	lst.record["last"] = -1
	return lst
}
//appends a new element to the end of the list, and updates the insertion and last element positions
func (il IntList) Push(elem int) {
	il.list[il.record["insert"]] = elem
	il.record["insert"] += 1
	il.record["last"] += 1
}
//removes the last element from the intlist and returns it
func (il IntList) Pop() int {
	popped, has := il.list[il.record["last"]]
	if has {
		delete(il.list, il.record["last"])
		il.record["last"] -= 1
		il.record["insert"] -= 1
		return popped
	}
	return 0
}
//gets an element in the inlist by the given int key
func (il IntList) Get(key int) int {
	got, has := il.list[key]
	if has {
		return got
	}
	return 0
}

//allows the setting of a new number to an existing key in the IntList. Cannot be used to push
func (il IntList) Set(key int, val int) {
	if key >= 0 && key < il.record["insert"] {
		il.list[key] = val
	}
}

//appends a new integer to the left side of the list, by moving all current elements up by one
func (il IntList) Pushleft(elem int) {
	for i:=il.record["last"] ;i>=0;i-- {
		il.list[i+1] = il.list[i]
	}
	il.list[0] = elem
	il.record["last"] += 1
	il.record["insert"] += 1
}

//gets a pretty string form of the IntList
func (il IntList) Tostring() string {
	lststr := "[ "
	for key := range il.list {
		lststr += strconv.Itoa(il.list[key]) + " "
	}
	lststr += "]"
	return lststr
}


