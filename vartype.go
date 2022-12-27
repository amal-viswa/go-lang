
package main
import "fmt"
 
func main()  {

			// print the type of variable

			var number int = 42
			var name string = "hwwk" 
			var boolean bool = true
			var amount float32 = 550000.00
		
			fmt.Printf(" The var = %v is of type %T \n", number, number)
			fmt.Printf(" The var = '%v' is of type %T \n", name, name)
			fmt.Printf(" The var = '%v' is of type %T \n", boolean, boolean)
			fmt.Printf(" The var = %v is of type %T \n", amount, amount)

}