package domain

type Book struct {
	ID     string
	Title  string
	Author string
	ISBN   string
}

func NewBook(id, title, author, isbn string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
		ISBN:   isbn,
	}
}

func (b *Book) ChangeTitle(newtitle string) {
	b.Title = newtitle
}
