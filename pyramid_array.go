// kyu6
// https://www.codewars.com/kata/515f51d438015969f7000013/

package kata

func Pyramid(n int) [][]int {
  result := [][]int{}
  temp := []int{}
  
  for i := 0; i < n; i++ {
    temp = append(temp, 1)
    
    result = append(result, temp)
  }
  
  return result
}