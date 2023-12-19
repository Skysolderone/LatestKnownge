package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
	Price   float64  `xml:"price"`
}

type BookStore struct {
	XMLName xml.Name `xml:"bookstore"`
	Books   []Book   `xml:"book"`
}

func main() {
	//basic generae xml
	book := Book{
		Title:  "Introduction to do",
		Author: "wws",
		Price:  29.99,
	}
	xmlData, err := xml.MarshalIndent(book, "", "       ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(xmlData))
	//parse
	xmlData1 := []byte(`
	<book>
		<title>Introduction to Go</title>
		<author>John Doe</author>
		<price>29.99</price>
	</book>
`)
	var book4 Book
	err = xml.Unmarshal(xmlData1, &book4)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(book4.Title, book4.Author, book4.Price)
	//read xml
	ls, err := os.ReadFile("data.xml")
	if err != nil {
		log.Fatal(err)
	}
	var bookstore BookStore
	err = xml.Unmarshal(ls, &bookstore)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range bookstore.Books {
		fmt.Println(v.Author, v.Price, v.Title)
	}
	//level decode xml file
	file, err := os.Open("data.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	for {
		token, err := decoder.Token()
		if err != nil {
			log.Fatal(err)
			break
		}
		switch ele := token.(type) {
		case xml.StartElement:
			if ele.Name.Local == "book" {
				var book5 Book
				err := decoder.DecodeElement(&book5, &ele)
				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Println(book5.Title)
			}
		}

	}
}
