// kyu 4
// https://www.codewars.com/kata/5886e082a836a691340000c3/

package kata

import m "math"

func RectangleRotation(a, b int) int {
  p1 := m.Sin(45 * m.Pi / 180) * float64(a) / 2
  p2 := m.Sin(45 * m.Pi / 180) * float64(b) / 2
  
  l1 := pointSlope(p1, p1, -1)
  l2 := pointSlope(-p2, p2, 1)
  l3 := pointSlope(-p1, -p1, -1)
  l4 := pointSlope(p2, -p2, 1)
  
  total := 0
  
  for i := -(a + b); i < (a + b); i++ {
    top := int(m.Floor(m.Min(l1(i), l2(i))))
    btm := int(m.Ceil(m.Max(l3(i), l4(i))))
    
    if top >= btm {
      total += top - btm + 1
    }
  }
  
  return total
}

func pointSlope(x1, y1, m float64) func(int)float64 {
  return func(x int) float64 {
    return m * (float64(x) - x1) + y1
  }
}