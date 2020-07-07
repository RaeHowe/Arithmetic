package main

import (
	"fmt"
	"time"
)

var COUNT = 100

func main()  {
	var tmpChan = make(chan int)

	go fun1(tmpChan)

	go fun2(tmpChan)

	time.Sleep(1 * time.Second)
}

func fun1(tmpChan chan int)  {
	for i := 0; i <= COUNT; i++{
		tmpChan <- i

		if i % 2 != 0{
			fmt.Println(i)
		}
	}
}

func fun2(tmpChan chan int)  {
	for i := 0; i <= COUNT; i++{
		<-tmpChan

		if i % 2 == 0{
			fmt.Println(i)
		}
	}
}
