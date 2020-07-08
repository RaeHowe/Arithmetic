package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

var COUNT = 100

func main(){

	f, err := os.Create("trace.out")
	if err != nil{
		panic(err)
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil{
		panic(err)
	}

	defer trace.Stop()

	var tmpChan = make(chan int)

	go fun1(tmpChan)

	go fun2(tmpChan)

	time.Sleep(1 * time.Second)
}

func fun1(tmpChan chan int)  {
	for i := 0; i < COUNT; i++{
		tmpChan <- i

		if i % 2 != 0{
			fmt.Println(i)
		}
	}
}

func fun2(tmpChan chan int)  {
	for i := 0; i < COUNT; i++{
		<- tmpChan

		if i % 2 == 0{
			fmt.Println(i)
		}
	}
}