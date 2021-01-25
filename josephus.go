// 5 kyu
// https://www.codewars.com/kata/555624b601231dc7a400017

package kata

func JosephusSurvivor(n, k int) int {
  nums := []int{}
  i := 0
  
  for i := 1; i <= n; i++ {
    nums = append(nums, i)
  }
  
  for {
    if len(nums) <= 1 {
      break
    }
    
    i = (i + k - 1) % len(nums)
    
    nums = append(nums[:i], nums[i+1:]...)
  }
  
  return nums[0]
}