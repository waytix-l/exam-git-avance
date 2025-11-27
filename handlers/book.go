package handlers

import (
	"net/http"
	"main/data"
)

var books = []data.Book{
	{
		ID:     1,
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Year:   2015,
	},
	{
		ID:     2,
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		Year:   2008,
	},
	{
		ID:     3,
		Title:  "The Pragmatic Programmer",
		Author: "Andrew Hunt",
		Year:   1999,
	},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	return
}