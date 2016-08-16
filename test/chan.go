package main

import "fmt"
import "time"

func fib(c, qiut chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-qiut:
			fmt.Println("qiut")
			return

		case <-time.After(time.Second * 2):
			fmt.Println("timeout")
			break
		}
	}
}

func main1() {
	c := make(chan int)
	qiut := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i, "\t", <-c)
			time.Sleep(time.Second)
			if i == 3 {
				time.Sleep(time.Second * 3)
			}
		}
		qiut <- 0
	}()
	fib(c, qiut)
}
