package prototype

import "fmt"

type Book struct {
	Title  string
	Page   int
	Author string
}

type BookAction interface {
	Read()
	CurrentPage(page int)
}

func (b *Book) Read() {
	fmt.Println("Ciee lagi baca buku")
}

func (b *Book) CurrentPage(page int) {
	fmt.Println("Halaman sekarang : ", page)
}

type Novel struct {
	Book  *Book
	Genre string
}

type Comic struct {
	Book
	genre string
}

func (n *Novel) Read() {
	fmt.Println("Cie lagi baca novel")
}

func (n *Novel) CurrentPage(page int) {
	fmt.Println("Novel sejauh ", page)
}
