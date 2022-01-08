package arrays

import "fmt"

func ArrayDeclaration() {
 /*array : An array is a systematic arrangement of 
 similar objects, usually in rows and columns */
	
 var a [5] int     //same datatype and same objects or products

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	fmt.Println(a)
	fmt.Println(len(a)) //length of array 
}