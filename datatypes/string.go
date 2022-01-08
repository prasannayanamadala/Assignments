package datatypes

import "fmt"

func StringImmutable() {

	/* strings are immutable in golang i.e once assigned cannot be changed. It can be changed  only through reassignment */
	
	var first string
	first = "this is first"
	fmt.Println(first)
	first = "this is"
	fmt.Println(first)


	//concatination of strings

	s1 := "can we go to the library"
	fmt.Println(s1)
	s2:= "get some books"
	fmt.Println(s2)
	s3 := s1 +" "+  "and" +s2 +"."
	fmt.Println(s3)
	fmt.Println(len(s1)) //length of string

	





}
