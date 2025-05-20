package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(c chan int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("close channel c")
		close(c)
	}(c)
	// read from channel
	for {
		select {
		case x, ok := <-c:
			// close(c)前：正确打印x, ok=true
			// close(c)后：打印x=0, ok=false
			fmt.Printf("%v, %v\n", x, ok)
			time.Sleep(time.Second)

			// 提前close, send channel panic
			/*
				if x == 5 {
					fmt.Println("close channel c 2")
					close(c)
				}
			*/
		}
	}
}
