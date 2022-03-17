package local

import "github.com/nelsongp/proxy/book"

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
