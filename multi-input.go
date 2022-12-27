package main
import "fmt"

func main()  {

		var name string
		var age int
		var place string
		
		fmt.Println("enter you name age and place : " )
		count, err := fmt.Scanf("%s %d %s", &name, &age, &place )
		fmt.Println("count :", count)
		fmt.Println("error:", err)
		fmt.Scanf("%s %d %s", &name, &age, &place)	
		fmt.Println(" Welcome ", name, age, place)
		

}
