// -----------------------------------
// Solution 1 - Time Complexity O(1)
// -----------------------------------
package main

import (
	"fmt"
	"math"
)

func termCount(a int, b int) int{
  
  	// finding n of nth term(First valid term) of type  x*(x+1) in the range [A ... B]
	tempFirstPos := int(math.Floor(math.Sqrt(float64(a))))
	firstPos := tempFirstPos
	for tempFirstPos*(tempFirstPos+1) < a {
	  tempFirstPos++
	  firstPos = tempFirstPos
	}
	
	// finding n of nth term(Last valid term) of type  x*(x+1) in the range [A ... B]
	tempLastPos := int(math.Floor(math.Sqrt(float64(b))))
	lastPos := tempLastPos
	for tempLastPos*(tempLastPos+1) > b {
	  tempLastPos--
	  lastPos = tempLastPos
	}
	
	
  	//no of terms of type x*(x-1) becomes a simple substraction . 
  	return (lastPos - (firstPos - 1))
}

func main() {
	var a, b int = 0, 0
	
	// please enter 2 space or new line seperated values in STDIN
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	fmt.Println("Total terms = ", termCount(a,b))
}
