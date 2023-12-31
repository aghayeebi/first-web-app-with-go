package main

import (
	"fmt"
	"os"
)

// define page structure
type Page struct {
	Title string
	Body  []byte
}

// save page
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 6000)
}

// load page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// main function

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
