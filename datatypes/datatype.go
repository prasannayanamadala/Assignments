package datatypes

import (
	"fmt"
)
func DataType() {
	
//string -text type
	
var today string 
today = "monday"
fmt.Println("today is:",today)

//Boolean -true or false
	
var x bool

x = true
fmt.Println("x :", x)

x = false
fmt.Println("x :", x)

//Integer -number type

var y int
y =10
fmt.Printf("%v, %T\n", y, y) //%v - value  %T - Type 

/*Int Ranges 
Int8 — [-128 : 127]
Int16 — [-32768 : 32767]
Int32 — [-2147483648 : 2147483647]
Int64 — [-9223372036854775808 : 9223372036854775807]
Int128 — [-170141183460469231731687303715884105728 : 170141183460469231731687303715884105727]
Int256 — [-57896044618658097711785492504343953926634992332820282019728792003956564819968 : 57896044618658097711785492504343953926634992332820282019728792003956564819967]*/

var y1 int8
y1 = 0
fmt.Println("value of y1:", y1) //Int8 — [-128 : 127]

var y2 int16
y2 = 32766
fmt.Println("value of y2:", y2) //Int16 — [-32768 : 32767]

var y3 int32
y3 = 2147483647
fmt.Println("value of y3:", y3) //nt32 — [-2147483648 : 2147483647]

var y4 int64
y4 = -9223372036854775808
fmt.Println("value of y4:", y4) //Int64 — [-9223372036854775808 : 9223372036854775807]

}