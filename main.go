

package main
 
import "fmt"
 
var global string = "variable"
var a int = 15
var b bool = true

func main() {
	//var name string = "amal"
	age := 25
	//fmt.Printf("hi my name is,  %v",name)
	{
		country := "UK"
		fmt.Println(country)
		fmt.Println(age)
		fmt.Println(a)
	}
	
	
	
	fmt.Printf( "my age is %d " , age)
	fmt.Println(global)
	fmt.Println(a)
	fmt.Println(b)
}

