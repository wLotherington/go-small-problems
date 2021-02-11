// 5kyu
// https://www.codewars.com/kata/5550d638a99ddb113e0000a2/

package kata

func Josephus(items []interface{}, k int) []interface{} {
	result := []interface{}{}
	idx := k - 1

	for len(items) > 0 {
		if idx >= len(items) && len(items) > 0 {
			idx = idx % len(items)
		}

		result = append(result, items[idx])
		items = append(items[:idx], items[idx+1:]...)

		idx = idx + k - 1
	}

	return result
}
