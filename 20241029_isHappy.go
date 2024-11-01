package main

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n = getSum(n)
		m[n] = true
	}
	return n == 1
}
func getSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
