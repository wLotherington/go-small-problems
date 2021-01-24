// kyu6
// https://www.codewars.com/kata/586538146b56991861000293

package kata
 
import s "strings"
 
func ToNato(words string) string {
  words = s.ToLower(words)
  nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot","Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
  result := []string{}
  
  for _, v := range words {
    idx := int(v) - 97
    
    if v == '.' || v == '!' || v == '?' {
      result = append(result, string(v))
    } else if idx <= 25 && idx >= 0 {
      result = append(result, nato[idx])
    }
  }
  
  return s.Join(result, " ")
}