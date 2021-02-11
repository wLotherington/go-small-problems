// 6kyu
// https://www.codewars.com/kata/585d7d5adb20cf33cb000235/train/go

package kata

func FindUniq(arr []float32) float32 {
  nums := map[float32]int{}

  for _, n := range arr {
    nums[n]++
  }

  for k, v := range nums {
    if v == 1 {
      return k
    }
  }

  return 0
}
