package main

import (
"fmt"
"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func gort_app() {
	go say("world")
	say("hello")

}


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}


func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}



func gort_app01() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)


	//############
	ch := make(chan int, 10)	//Sends block only when the buffer is full. Receives block when the buffer is empty.
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)



	//	######
	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}


}
