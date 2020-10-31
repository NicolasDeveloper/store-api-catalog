package domain

import "time"

//DateNow get current date with location configured
func DateNow() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}
