package main

import (
	"fmt"
	"math"
)

func termCount(a int, b int) int{
  
  // finding n of nth term(First valid term) of type  x*(x+1) in the range [A ... B]
	tempFirstPos := int(math.Floor(math.Sqrt(float64(a))))
	firstPos := tempFirstPos
	if tempFirstPos*(tempFirstPos+1) < a {
	  firstPos = tempFirstPos+1
	}
	
	// finding n of nth term(Last valid term) in the range [A ... B] of type x*(x+1)
	tempLastPos := int(math.Floor(math.Sqrt(float64(b))))
	lastPos := tempLastPos
	if tempLastPos*(tempLastPos+1) > b {
	  lastPos = tempLastPos-1
	}
	
	
	//fmt.Println("firstPos = ",firstPos, "firstElem = ", firstPos*(firstPos+1), "lastPos = ",lastPos, "lastElem = ", lastPos*(lastPos+1))
	
  //no of terms of type x*(x-1) becomes a simple substraction. 
  return (lastPos - (firstPos - 1))
}

func main() {
	var a, b int = 0, 0
	
	// please enter 2 space, tab or new line seperated valid input integer values in STDIN
	
	fmt.Println("Enter Range Start value A = ")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter Range End value B = ")
	fmt.Scanf("%d", &b)

	
	fmt.Println("Total terms = ", termCount(a,b))
}

