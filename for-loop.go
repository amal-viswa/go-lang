/* for loop 


for initialization statement; condition; post {
		
		// statements
}

*/

/* example for loop by initializing in loop 

 ----------------
  for loop
 -----------------

package main
 import "fmt"

func main()  {

	for i := 1; i <= 5; i++ {

		fmt.Println(i*i)
	}


	
}
 */

 /* example for for loop initializing outside the loop 

  -----------------
  for loop
  -----------------

 package main
 import "fmt"

 func main()  {
	
	i := 2
	for i <= 5 {
		fmt.Println(i*i)
	i += 2
	
	}

 }

 */

 /* infinte loop 

 examp-1

 package main
 import "fmt"

 func main()  {

	sum := 0
	for {
		sum++ // repeated forever
	}
	fmt.Println(sum) // never reached
	
 }

exapl 2

package main

import "fmt"

func main() {
        for {
                fmt.Println("Hello World!")
        }
}

 */

 /* 
  -----------------
  break in for loop
  -----------------

  
  package main
  import "fmt"

  func main()  {

	for i:=1; i <=5; i++{
		if i == 3 {
			break
		}
	
	fmt.Println(i)
	}

  }
  
  */

  /*  
------------------------
  Continue in for loop 
------------------------


package main
import "fmt"

func main() {

	for i:=1; i <=5; i++ {
		if i == 3 {
			continue
		}
	
	fmt.Println(i)
	}
}

*/

/* 
exaple with no output for for

package main

import "fmt"

func main() {
        i := 3
        for i > 10 {
                fmt.Println(i * 2)
                i += 1
        }
	}

/*

FROM alpine:latest

ENV MYVAR=$VAR
RUN apk update
CMD printenv | grep MYVAR