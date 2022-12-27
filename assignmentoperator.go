package main
import "fmt"

func main()  {

// eg for assignment (=) operator
	// var x int = 10
	// var y int
	// y = x 
	//	fmt.Println(y)

// eg for (add and assign(+=)) | (substract and assign (-=)) | [multiply and assign (*=)] | [divide and assign quotient (/=)] | [divide and assign modulus (%=)] operator

	var x, y int = 10 , 20 

//	x += y  // it means x = x+y
//	x -= y  // it means x = x-y
//	x *= y  // it means x = x * y
//	x /= y  // it means x = x / y
	x %= y  // it means x = x % y
fmt.Println(x)
}

/* {
	var x, y int = 100,9
	x /= y
	fmt.Println(x)
	x %= y
	fmt.Println(x)
/*}