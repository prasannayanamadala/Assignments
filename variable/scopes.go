package variable

import (
	"fmt"
)
var pkgvariable int = 202    // package level scope : creating a variable at a package level
							/* variables defined outside a function can be accesible by all the functions with the same package name */
var pkgvariable1 bool = false

var pkgvariable2 string = "me"

var Pkgvariable3 int = 203   /* Global scope : variables that are defined using the beginning letter 
							   capitalized can be accesible anywhere i.e even in a different package */

func ScopeDeclaration() {

	var x int
	x=200
	fmt.Println("value of x :", x)

	var y string = "Newyear eve"
	fmt.Println("value of y :", y)

	z := true
	fmt.Println("value of z :", z)

	fmt.Println("pkgvariable2:" , pkgvariable2)
	
/* Functional scope : variables defined inside a function are local to that function and cannot be accesible outside of that function */
}
func NextFunc() {        

	var x int = 201
	fmt.Println("value of x in nextfunction:",x)

	fmt.Println("pkgvariable1:" , pkgvariable1)
}