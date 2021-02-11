// 7kyu
// https://www.codewars.com/kata/569218bc919ccba77000000b/

package kata

import "time"

func DateNbDays(a0 int, a int, p int) string {
	i := float64(p) / 36000
	start := float64(a0)
	end := float64(a)

	days := 0

	for start < end {
		start *= (1 + i)
		days++
	}

	startDate := time.Date(2016, 01, 01, 0, 0, 0, 0, time.UTC)
	startDate = startDate.AddDate(0, 0, days)

	return startDate.Format("2006-01-02")
}
