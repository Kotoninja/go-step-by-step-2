package main

import (
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	defer close(prime_nums)
	timeout := time.AfterFunc(100*time.Millisecond, func() {
		stop <- struct{}{}
	})
	for i := 0; i <= N; i++ {
		if numberIsPrime := isPrime(i); !numberIsPrime {
			continue
		}
		select {
		case <-stop:
			timeout.Stop()
			return
		case prime_nums <- i:
			continue
		}
	}
	timeout.Stop()
}
