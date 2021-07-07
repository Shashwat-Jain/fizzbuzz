package fizzbuzz

// Fizzbuzz takes a number and returns either a string and bool
// true or and empty string and boolean false
//
// callers are supposed to print original number if false is
// returned.
func Fizzbuzz(num int) (string, bool) {
	if num%15 == 0 {
		return "fizz buzz", true
	}

	if num%3 == 0 {
		return "fizz", true
	}

	if num%5 == 0 {
		return "buzz", true
	}

	return "", false
}
