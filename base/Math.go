package base

import "fmt"

func Chan(ch chan int, stopChan chan bool) {
	i := 10
	for j := 0; j < i; j++ {
		ch <- i
	}
	stopChan <- true
}

func Add(x, y int) {
	ch := make(chan int)
	stopChan := make(chan bool)
	c := 0
	go Chan(ch, stopChan)
	for true {
		select {
		case c = <-ch:
			fmt.Println("Receive", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-stopChan:
			goto end
		}
	}
end:
	fmt.Println("this is end")
}

type Driver interface {
	Add(x, y int) int
	Sub(x, y int) int
}
