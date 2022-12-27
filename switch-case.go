/* 

switch expression {
case value_1:
	// execeutes when the expression equals to value_1

case value_1:
	// execeutes when the expression equals to value_1

case value_1:
	// execeutes when the expression equals to value_1

default:
	//executes when no matche is found

}

*/

/*
--------------
switch case eg
---------------

package main

import "fmt"

func main()  {

	var i int = 10
	switch i  {
	case 10:
		fmt.Println("i is 10")
		
	case 100 ,200:
		fmt.Println("i is 100 or 200")
		
	default:
		fmt.Println("i is neither 0 , 100 , 200")	
	}
}

*/

/*  fallthrough in switch-case

the fall through is used in switch-case to force the execution flow to fall through the successive case blocks

basically prinits all the cases until it finfs the default or until it goes through the case with no fall through 

eg for fall through
-----------------------------
package main

import "fmt"

func main()  {

	var i int = 10
	switch i  {
	case 10:
		fmt.Println("i is 10")
		fallthrough
	case 100 ,200:
		fmt.Println("i is 100 or 200")
		fallthrough
	default:
		fmt.Println("i is neither 0 , 100 , 200")	
	}
}

*/

/*
// switch case with condition  we do not need to write an expression for the switch
	
switch  {
case condition_1:
	 // executre when condition_1 is true
case condition_2:
	// executes when condition_2 is true
default:
	//executes when no conditon is true	
}
 */

 package main

 import "fmt"

 func main()  {

	var a, b int = 10, 30

	switch  {
	
	case a+b == 30:
		fmt.Println("equal to 30")

	case a+b <= 30:
		fmt.Println("less than or equal to 30")
	
	default:
		fmt.Println("greater than 30")
		
	}
	
 }





