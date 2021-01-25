// kyu6
// https://www.codewars.com/kata/5269452810342858ec000951/

package kata

import (
  sc "strconv"
  s "strings"
)

func validCoord (coord string) bool {
  counts := map[string]int{"decimal": 0, "nums": 0, "minus": 0}
  
  for _, v := range coord {
    switch v {
      case '-':
        if counts["minus"] > 0 || counts["decimal"] > 0 || counts["nums"] > 0 {
          return false
        }
        counts["minus"]++
      case '.':
        if counts["decimal"] > 0 {
          return false
        }
        counts["decimal"]++
      case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        counts["nums"]++
      default:
        return false
    }
  }
  
  return true
}

func IsValidCoordinates(coordinates string) bool {
  coords := s.Split(coordinates, ", ")
  
  if !validCoord(coords[0]) || !validCoord(coords[1]) {
    return false
  }
  
  lat, _ := sc.ParseFloat(coords[0], 64)
  long, _ := sc.ParseFloat(coords[1], 64)
  
  if lat > 90 || lat < -90 {
    return false
  }
  
  if long > 180 || long < -180 {
    return false
  }
  
  return true
}