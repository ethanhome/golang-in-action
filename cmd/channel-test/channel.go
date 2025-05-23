package main

import (
	"fmt"
	"time"
)

/*
channel:
unbuffer channel, buffer channel
ch := make(chan int) // ch has type 'chan int'
ch <- x  // a send statement
x = <-ch // a receive expression in an assignment statement
<-ch     // a receive statement; result is discarded

Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。
对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；
如果channel中已经没有数据的话将产生一个零值的数据。

使用内置的close函数就可以关闭一个channel：
close(ch)

ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3

没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式：
它多接收一个结果，多接收的第二个结果是一个布尔值ok，ture表示成功从channels接收到值，
false表示channels已经被关闭并且里面没有值可接收

// Squarer
go func() {
    for {
        x, ok := <-naturals
        if !ok {
            break // channel was closed and drained
        }
        squares <- x * x
    }
    close(squares)
}()


类型chan<- int表示一个只发送int的channel，只能发送不能接收。
相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送。
（箭头<-和关键字chan的相对位置表明了channel的方向。）

*/

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
