package main

import (
	"fmt"
	"time"
)

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciExe(n int, c chan<- int) {
	res := Fibonacci(n)
	c <- res
}

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be non-negative")
	} else if timeout == 0 {
		return 0, fmt.Errorf("timeout")
	}
	c := make(chan int, 1)

	go FibonacciExe(n, c)
	select {
	case res := <-c:
		return res, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("timeout")
	}
}

func main() {
	res, err := TimeoutFibonacci(10, time.Second)
	fmt.Println(res, err)
}
