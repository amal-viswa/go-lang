//  array are a collection of similar data elements stored in contiguous memory location (1234)
// array in go are of fixed length once it's is defined we cannot change 
// elements should be of thw same data type

// sysntax var <name of arry > <size of array> <data type>

package main
import "fmt"

func main(){
// example array	

var grades [5] int
fmt.Println(grades)

 var fruits [5] string
 fmt.Println(fruits) 


 // initialising array
var value [4]int = [4]int{10,20,30,40}

//shorthand init
value1 := [3]int{5,10,15}

//array using elipses(not specifying any length or size of array)
val := [...]string{"hello","world", "hello", "go"}

fmt.Println(value)
fmt.Println(value1)
fmt.Println(val)


// "  len() " the length of an array refers to the number of elements stored in array (takes an array as input and returs the size of the array)
fmt.Println(len(val))

// print based on indexes
fmt.Println(val[2])

//changing init values of an array 
 fmt.Println("without changing the values of int" ,value)
value[1] = 15 // changing initialised value at the postion of array[1]
fmt.Println(" after changing the value", value) 

//changing init value of string
   
 fmt.Println("without changing the value of string" ,val) 
 val[0] = "happynewyear"
 fmt.Println("After changing the string value", val)

// looping through an array

for i := 0; i < len((value)); i++ {
	fmt.Println(value[i])  // looping through int
}

for i := 0; i < len((val)); i++ {
	fmt.Println(val[i])     // looping through string

}
// loop using range keyword
/*
loop using range keyword "=>" .
 the range keyword is used in for loops to itreate over all the elements of the array ,slies or maps
 
syntax :- 

for index, element := range <array-name> {  // range is the scope of itreation upto the length of the array

		fmt.Println(index, "=>", element)  // index and element are the varables that stores the return values for the rang expression
}
*/
// example for looping through range keyword

for index , element := range val{
	fmt.Println(index,"=>",element)
}


// multidiamensional arrays (array of array)

// example for twodimensional array
	arr := [3][2]int{{2,4},{4,8},{8,16}}  // "[3] refers to the size of the array or elements in the array "/ and [2] refers to the size of each elements itself 
	fmt.Println(arr[0][1]) 
// [0] means that go to the first element in the main array and from that element we want the first index ie "[1] "
	fmt.Println(arr[1][0])
	fmt.Println(arr[2][1])
}
