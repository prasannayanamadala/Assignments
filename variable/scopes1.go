package variable

import "fmt"

func Scope1Func() {

	/* variable defined outside a function can be accesible by all the functions with the same package name */

	fmt.Println("pkgvariable:" , pkgvariable)
	
	
}