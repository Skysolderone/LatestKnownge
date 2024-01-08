package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	booksAPI := app.Party("/books")
	{
		//booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}

	app.Listen(":8080")
}
