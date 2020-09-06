package base

import "fmt"

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
	fmt.Println("*Book title的内存地址为", book, &book.title)
	book.title = "change title"
}

func printTitleAddress(book *Book) {
	fmt.Println("title的内存地址为", book, &book.title)
}

func printBook(book Book) {
	println(book.title)
	println(book.author)
	println(book.price)
	println(book.description)

}

func Type() {
	printBook(Book{"Kafka权威指南", "BomThor", 67.6, "最权威的kafka学习资料"})
	var book = Book{"Kafka权威指南", "BomThor", 67.6, "最权威的kafka学习资料"}
	book.sale()
	printTitleAddress(&book)
	empty := new(Empty)
	fmt.Println(empty)
	fmt.Println(book)
}
