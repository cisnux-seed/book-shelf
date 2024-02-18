package entity

import (
	"encoding/base64"
	"fmt"
	"math/rand"
)

type Book struct {
	id         string
	title      string
	writerName string
}

func (b *Book) Id() string {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) WriterName() string {
	return b.writerName
}

type Builder struct {
	book *Book
}

func BookBuilder() *Builder {
	return &Builder{
		book: &Book{
			id:         "",
			title:      "",
			writerName: "",
		},
	}
}

func (builder *Builder) SetId(id string) {
	builder.book.id = id
}
func (builder *Builder) SetTitle(title string) {
	builder.book.title = title
}

func (builder *Builder) SetWriterName(writerName string) {
	builder.book.writerName = writerName
}

func (builder *Builder) Build() *Book {
	if builder.book.id == "" {
		builder.book.id = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%s%s", rand.Int(),
			builder.book.Title(),
			builder.book.WriterName()),
		))
	}
	return builder.book
}
