// kyu6
// https://www.codewars.com/kata/566be96bb3174e155300001b/

package kata

const g float64 = 0.0981

func MaxBall(v0 int) int {
  v := float64(v0) / 36
  t := 0.0
  h := 0.0
  maxH := 0.0
  
  for {
    t++
    h = float64(v) * t - 0.5 * g * t * t
    
    if h < maxH {
      t--
      break
    }
    
    maxH = h
  }
  
  return int(t)
}