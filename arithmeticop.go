package main
import "fmt"

func main()  {
	
	var a, b int = 25, 2
	var ab, cd string = "hello", "web"

	// increment operator ( ++ )
	var i int = 2 
	 i++

	// decrement operator ( -- )
	var f int = 5
	f--
	fmt.Println(ab + cd) //using string in addition (it will add and show like concatation)
	fmt.Println(a + b) //using addition + operator
	fmt.Println(a - b) //using substraction - operator
	fmt.Println(a * b) //using multiplisction * operator
	fmt.Println(a % b) //using division % modules operator (it will return the reminder)
	fmt.Println(a / b) // using division / operator (it will print the quotient)
	fmt.Println(i) //using increment operator
	fmt.Println(f) //using decrement operator

}