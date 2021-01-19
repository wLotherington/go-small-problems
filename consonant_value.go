// kyu6
// https://www.codewars.com/kata/59c633e7dcc4053512000073/

package kata

func solve(str string) int {
  max := 0
  subtotal := 0
  vowels := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true,}
  
  for _, c := range str {
    s := string(c)
    
    if vowels[s] {
      if subtotal > max {
        max = subtotal
      }
      subtotal = 0
    } else {
      subtotal += (int(c) - 96)
    }
  }
  
  if subtotal > max {
    max = subtotal
  }
  
  return max
}