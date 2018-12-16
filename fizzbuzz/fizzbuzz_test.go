package fizzbuzz

import "testing"

func TestFizzBuzzSouldReturnOne(t *testing.T) {
	expected := "1"
	result := FizzBuzz(1)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnTwo(t *testing.T) {
	expected := "2"
	result := FizzBuzz(2)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzSouldReturnFizz(t *testing.T) {
	expected := "fizz"
	result := FizzBuzz(3)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzSouldReturnFour(t *testing.T) {
	expected := "4"
	result := FizzBuzz(4)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnBuzz(t *testing.T) {
	expected := "buzz"
	result := FizzBuzz(5)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnFizzWhenGetSix(t *testing.T) {
	expected := "fizz"
	result := FizzBuzz(6)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzSouldReturnSeven(t *testing.T) {
	expected := "7"
	result := FizzBuzz(7)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnEight(t *testing.T) {
	expected := "8"
	result := FizzBuzz(8)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnFizzWhenGetNine(t *testing.T) {
	expected := "fizz"
	result := FizzBuzz(9)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnFizzBuzz(t *testing.T) {
	expected := "buzz"
	result := FizzBuzz(10)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnFizzWhenGetTwelve(t *testing.T) {
	expected := "fizz"
	result := FizzBuzz(12)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
func TestFizzBuzzSouldReturnFizzBuzzWhenGetFifthteen(t *testing.T) {
	expected := "fizzbuzz"
	result := FizzBuzz(15)
	if expected != result {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
