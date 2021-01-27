// kyu6
// https://www.codewars.com/kata/534d2f5b5371ecf8d2000a08/

package multiplicationtable

func MultiplicationTable(size int) [][]int {
  var table [][]int
  var row []int
  
  for i := 1; i <= size; i++ {
    row = []int{}
    
    for j := 1; j <= size; j++ {
      row = append(row, i * j)
    }
    
    table = append(table, row)
  }
  
  return table
}