package main

func Send(ch1, ch2 chan int) {
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	go func() {
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
	}()
}
