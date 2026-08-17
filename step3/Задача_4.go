package main

func Process(nums []int) chan int {
	ch := make(chan int)

	go func() {
		for _, number := range nums {
			ch <- number
		}
	}()

	return ch
}
