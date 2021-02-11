// 7kyu
// https://www.codewars.com/kata/59cfc09a86a6fdf6df0000f1/

package kata

import "strings"

func Capitalize(st string, arr []int) (result string) {
	for i, char := range st {
		if len(arr) > 0 && i == arr[0] {
			arr = arr[1:]
			result += strings.ToUpper(string(char))
		} else {
			result += strings.ToLower(string(char))
		}
	}

	return
}
