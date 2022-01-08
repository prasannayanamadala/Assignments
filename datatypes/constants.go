package datatypes

import "fmt"

const GlobalConst int =12  /*global constants(can be used anywhere with the 
							package name datatypes.values cannot be changed because they are constants)*/

const pkgConst int = 15    /*package constants(can be used anywhere in the package.values cannot be 
							changed once assigned as they are constants)*/

func ConstantDefine() {     /*functional constants(can be used inside the function.values 
								cannot be changed as they are sonstants)*/
	const cvar = 101
	fmt.Println(cvar)
	/*cvar := 55  const cvar int = 101
cannot assign to cvarcompiler */

	const str = "look"
	fmt.Println(str)

	const bool = true
	fmt.Println(bool)

	x := 101 + cvar  //constant variables values cannot be changed but can be used
	fmt.Println(x)
	 
}
func PkgConst() {
	c := pkgConst + GlobalConst
	fmt.Println(c)
}