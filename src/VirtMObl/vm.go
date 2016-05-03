
package VirtMObl
//file to implement the virtual machine


type VirtualMachine struct {
	/*
	This map implements callable operator functions in O(1) time
	Contains both math and comparison
	*/
	Ints map[string]func(a, b int) int

	//map for bool comparisons of ints
	Intbools map[string]func(a, b int) bool
	//used for storing vars for ints
	Intvars map[string]int
	
}
//method that sets a value into the intvars map.
func (vm VirtualMachine) Setint(name string, val int) {
	vm.Intvars[name] = val
}

func CreateVM() VirtualMachine {
	machine := VirtualMachine{Ints:make(map[string]func(a, b int)int), Intvars:make(map[string]int), Intbools:make(map[string]func(a, b int)bool)}
	machine.Ints["+"] = func(a, b int)int {return a+b}
	machine.Ints["-"] = func(a, b int)int {return a-b}
	machine.Ints["*"] = func(a, b int)int {return a*b}
	machine.Ints["/"] = func(a, b int)int {return a/b}
	machine.Ints["mod"] = func(a, b int)int {return a+b}
	machine.Intbools["=="] = func(a, b int)bool {return a==b}
	machine.Intbools["!="] = func(a, b int)bool {return a!=b}
	machine.Intbools["<"] = func(a, b int)bool {return a<b}
	machine.Intbools[">"] = func(a, b int)bool {return a>b}
	machine.Intbools[">="] = func(a, b int)bool {return a>=b}
	machine.Intbools["<="] = func(a, b int)bool {return a<=b}
	return machine
}

