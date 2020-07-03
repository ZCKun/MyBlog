package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{
		Title: title,
		Body:  body,
	}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	fmt.Printf("get %s path with request from %s \n",
		url.Path,
		url.Host)

	page, err := loadPage("TestPage")
	if err != nil {
		page = &Page{
			Title: "Error Page",
			Body:  []byte("Load TestPage File has error: " + err.Error()),
		}
	}
	_, _ = fmt.Fprint(w, string(page.Body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
