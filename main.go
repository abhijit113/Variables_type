package main

import "fmt"

var y = 49 //scope aof this variable global

//Declare there is a variable with the identifier z
//and that the variable with the identifier z of type int
//assigns the zero value of type int to "z"
//false for booleans, o for integers, 0.0 for floats
//"" for strings and nil for pointers, functions, slices,
//channels and maps
var z int
var p=`It's lockdown and "learning go at the moment"`

//It's a static programming language and z can not be assigned other than int;
func main() {
	fmt.Println("Inside a function")
	x := 54 //only local scope
	fmt.Println(x)
	fmt.Println(z)

	foo()

}

func foo() {
	fmt.Println("again :", y)
	fmt.Printf("%T\n",y)    //T stands for typefmt.Println("again :", y)
	fmt.Printf("%b\n",y)
	fmt.Printf("%d\n",y)
	fmt.Printf("%x\n",y)
	fmt.Printf("%#x\n",y)
	fmt.Printf("%#x\t%b\t%d\n ",y,y,y)
	s:=fmt.Sprintf("%#x\t%b\t%d\n ",y,y,y)
	fmt.Println(s)
	z=99
	fmt.Println("Globally Declared :", z)
	fmt.Printf("%T\n",z)

	fmt.Println(p)

}
