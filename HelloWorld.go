package main

import (
	"./base"
	"fmt"
	"os"
)

func main() {

	var slice = []int{0, 1}
	slice = append(slice, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(cap(slice), len(slice))

	for key, arg := range os.Args {
		fmt.Println(key, arg)
	}

	for i, v := range slice {
		fmt.Println(&i, &v, &slice[i])
	}

	kv := map[string]string{"1": "2"}
	for key, value := range kv {
		fmt.Printf("key %s,value %s", key, value)
		fmt.Println()
	}

	base.Add(1, 3)
	printBook(Book{"Kafka权威指南", "BomThor", 67.6, "最权威的kafka学习资料"})
	var book = Book{"Kafka权威指南", "BomThor", 67.6, "最权威的kafka学习资料"}
	book.sale()
	empty := new(Empty)
	boo := new(Book)
	fmt.Println(boo)
	fmt.Println(empty)
	fmt.Println(book)

	ch := make(chan int)
	go base.RuntimeAdd(1, 2, ch)
	go base.RuntimeAdd(2, 3, ch)
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






type Book struct {
	title       string
	author      string
	price       float32
	description string
}
type Empty struct {
}

type Salable interface {
	sale() bool
}

func (book *Book) sale() {
	fmt.Println(book.title)
	book.title = "这是修改之后的title"
}

func printBook(book Book) {
	println(book.title)
	println(book.author)
	println(book.price)
	println(book.description)

}
