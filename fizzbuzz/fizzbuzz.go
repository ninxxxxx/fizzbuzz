package fizzbuzz

import "strconv"

//FizzBuzz return fizz when can devide by 3, buzz when can devide by 5 and fizzbuzz when can by both 3 and 5
//return what it get when not 3 and 5
func FizzBuzz(n int) string {

	if n%3 == 0 && n%5 == 0 {
		return "fizzbuzz"
	}

	if n%3 == 0 {
		return "fizz"
	}

	if n%5 == 0 {
		return "buzz"
	}

	return strconv.Itoa(n)
}
