package main

import "fmt"

func Pyramid(n int) [][]int {
  result := [][]int{}
  temp := []int{}
  
  for i := 0; i < n; i++ {
    temp = append(temp, 1)
    
    result = append(result, temp)
  }
  
  return result
}

func main() {
	fmt.Println(Pyramid(0)) // [ ]
	fmt.Println(Pyramid(1)) // [ [1] ]
	fmt.Println(Pyramid(2)) // [ [1], [1, 1] ]
	fmt.Println(Pyramid(3)) // [ [1], [1, 1], [1, 1, 1] ]
}