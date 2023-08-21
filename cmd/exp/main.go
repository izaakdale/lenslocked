package main

import (
	"html/template"
	"os"
)

type user struct {
	Name string
	Bio  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	u := user{"izaak", `<script>alert("hi")</script>`}

	err = t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
