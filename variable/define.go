package variable

import "fmt"

func VarDeclaration() {

	//Defining and Declaring two integer variables in different ways

	//variable Definition
	//variable Declaration

	/* using Shorthand declaration */

	age := 1
	fmt.Println(age)

	/* A variable can be defined only once,but can be declared any number of times */

	age = 2
	fmt.Println(age)

	//variableDeclaration
	//variable Definition

	age1 := 3
	fmt.Println(age1)

	//variable defition
	/* variable,variable name and variable type */

	var year int

	//variable declaration

	year = 2022
	fmt.Println(year)

	//variable definition
	//variable declaration

	var year1 int = 2000
	fmt.Println(year1)

	Add(1, 10)
}

/*performing simple arithmatic operation taking two integers*/

func Add(a int, b int) {
	var c int
	c = a + b
	fmt.Println(c)

	//Defining and Declaring two string variables

	/* using shorthand declaration */

	str := "Busy Bee"

	fmt.Println(str)

	str = "Bee"

	fmt.Println(str)

	//variable Definition
	//variable Declaration

	str1 := "snooze"

	fmt.Println(str1)

	var month string

	//variable declaration

	month = "December"
	fmt.Println(month)

	//variable definition
	//variable declaration

	var month1 string = "january"
	fmt.Println(month1)

}
