package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// name := os.Args[1]

	// str := fmt.Sprint(`
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	// <meta charset="UTF-8">
	// <title>Hello World!</title>
	// </head>
	// <body>
	// <h1>` + name + `</h1>
	// </body>
	// </html>
	// `)

	// nf, err := os.Create("index.html")
	// if err != nil {
	// 	log.Fatal("Error creating file", err)
	// }
	// defer nf.Close()

	// io.Copy(nf, strings.NewReader(str))

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
