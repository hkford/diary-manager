// must not import a package called "init", so package name is a little bit long.
// https://go-review.googlesource.com/c/go/+/287494/
package initialize

func IsLeapYear(year int64) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
