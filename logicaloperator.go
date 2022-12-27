package main

import "fmt"
func main()  {
	
	var a int = 10 
	var x, y int = 15, 25

	

	// and(&&) logical operator
	fmt.Println((a<100) && (a<200))
	fmt.Println((a<20) && (a>200)) 

	// or (||) logical operartor
	fmt.Println((a>5) || (a<5))
	fmt.Println((a<5) || (a<5))

	// not (!) logical operator(unary operator - means it will work only on operand or one expression)
	fmt.Println(!(x<y))
	fmt.Println(!(x>y))
	fmt.Println(!(true))
	fmt.Println(!(false))
	
}