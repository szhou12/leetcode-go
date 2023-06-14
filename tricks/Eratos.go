package tricks

// Sieve of Eratosthenes 埃氏筛
// Return a list of prime numbers <= n
func eratosthenes(n int) []int {
	var primes []int
	sieve := make([]int, n+1)
	for i := 2; i*i <= n; i++ {
		if sieve[i] == 1 {
			continue
		}
		j := i * 2
		for j <= n {
			sieve[j] = 1
			j += i
		}
	}

	for i := 2; i <= n; i++ {
		if sieve[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}
