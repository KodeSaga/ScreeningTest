//
// - Time Complexity - O(n)
// - Solution is the function maxSameHeight(A,K)
//
package main

import (
	"fmt"
)

func maxSameHeight(A []int, K int) int{
  
  currentHtCount := 0
  nailsAfter := 0
  tempMaxSameHtCount := 0
  maxSameHtCount := 0
  
  
  for i:=0; i<len(A); i++ {
    if i == len(A)-1 || A[i] < A[i+1] {
      currentHtCount++
      nailsAfter = len(A) - (i+1)
      if i == len(A)-1 {
        nailsAfter = 0
      }
      if nailsAfter - K < 0 {
        tempMaxSameHtCount =  currentHtCount + nailsAfter
      } else {
        tempMaxSameHtCount =  currentHtCount + K
      }
      if tempMaxSameHtCount > maxSameHtCount {
        maxSameHtCount = tempMaxSameHtCount
      }
      currentHtCount = 0
      tempMaxSameHtCount = 0
      continue
    }
    currentHtCount++
  }
 
  return (maxSameHtCount)
}

func main() {
  // Edit below variables A and K for custom inputs 
  A := []int{1, 1, 3, 3, 3, 4, 5, 5, 5, 5, 6}
	K := 3
	
	fmt.Println("\nMax nails at same height = ", maxSameHeight(A,K))
}
