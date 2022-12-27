
/* if else

if (condition) {
	
	executes when condition is true 
	
} else {     // the else should start from the same line where the if ends or else we'll get a error 

	executes when condition is false

}


--------------------------------
eg for else if 
----------------------

package main
import "fmt"
func main() {

	 var a string = "apple"
	 
	 if a == "orange" {
		fmt.Println(a)
	 } else {
		fmt.Println("does not matches ")
	 }
}

---------------------------------------------------

*/

/* nested if  

if condition_1 {
	
	executes when condition_1 is true

} else if condition_2 {

	executes when condition_1 is false, and condition_2 is true

	} else if condition_3{

	executes when condition_1 and 2 is false, and condition_3 is true
} else {

	executes when none of the above conditions are true
}
*/

----------------------------
nested if 
---------------------------

package main
import "fmt"
func main() {

	 var a string = "none"
	 
	 if a == "orange" {
		fmt.Println("i love",a)
	 } else if  a == "orange" {
		fmt.Println("i love ", a)
	 } else if a == "grape"{
		fmt.Println("i love", a)
	 } else {
		fmt.Println("none of the above")
	 }
}