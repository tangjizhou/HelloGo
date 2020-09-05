package base

func RuntimeAdd(x, y int, ch chan int) {
	ch <- x + y
}
