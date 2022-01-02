package main

import (
	"fmt"

	"github.com/prasannaraavi/assignments/variable"
)

func main() {

	variable.VarDeclaration()
	
	variable.ScopeDeclaration()
	
	variable.NextFunc()      //Functional scope

	variable.Scope1Func()    //package level scope

	fmt.Println("pkgvariable3:", variable.Pkgvariable3)  //global scope
	

}
