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

/*unsigned integers:An unsigned integer is a 32-bit datum(it depends on the pc if it is 32 bit or 64bit) that encodes a nonnegative integer in the range [0 to 4294967295]*/

/*Operation	Output Range	             Bytes per Element
uint8	     0 to 255	                        1
uint16	     0 to 65,535	                    2
uint32	     0 to 4,294,967,295	                4
uint64	     0 to 18,446,744,073,709,551,615	8 */

var y5 uint
y5 = 184675876798709
fmt.Println("value of y5:", y5)

var y6 uint8
y6 = 255   /* if y6 value is given as 256,it shows  "use 256 (untyped int constant) as uint8 value in assignment " */
fmt.Println("value of y6:", y6)

var y7 uint16
y7 = 56765
fmt.Println("value of y7:", y7)

var y8 uint32
y8 = 657869708  
fmt.Println("value of y8:", y8)

//floating point :float32 is a 32 bit number - float64 uses 64 bits

var num float32 //16bit for decimal and 16bit for integral part
num = 77.56    //maximum value : 2^(16-1)-1.2^(16-1)
fmt.Println(num)

num = -544.786
fmt.Println(num)

var num1 float64 // maximum values : 2^(32-1)-1.2^(32-1)
num1 = 33333.555555
fmt.Println(num1)


//Arithmetic operations
         
//signed integers

sum := int16(y1) + y2
fmt.Println("sum of signed integers:",sum)

sum1 := int64(y3) + y4
fmt.Println("sum of signed integers:",sum1)

//unsigned integers

sum2 := uint8(y5) + y6
fmt.Println("sum of unsigned integers:",sum2)

sum3 := uint16(y6) + y7
fmt.Println("sum of unsigned integers:",sum3)

//floating point

sum4 := num + num
fmt.Println("sum4:",sum4)

sum5 := float64(num) + num1
fmt.Println("sum5:",sum5)

}