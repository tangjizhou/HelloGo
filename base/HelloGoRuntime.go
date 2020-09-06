package base

import "fmt"

func add(x, y int, ch chan int) {
	ch <- x + y
}

func Go() {
	ch := make(chan int)
	go add(1, 2, ch)
	go add(2, 3, ch)
	fmt.Println(ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y)

	ch = make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	close(ch)
	//ch <- 4
	fmt.Println(<-ch)
}

func Runtime() {
	ch := make(chan int)
	stopChan := make(chan bool)
	c := 0
	go channel(ch, stopChan)
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

func channel(ch chan int, stopChan chan bool) {
	i := 10
	for j := 0; j < i; j++ {
		ch <- j
	}
	stopChan <- true
}
