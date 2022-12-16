package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Prakhar-Agarwal-byte/go-book-management-system/pkg/models"
	"github.com/Prakhar-Agarwal-byte/go-book-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic("error while parsing")
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	b := models.CreateBook(&book)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	params := mux.Vars(r)
	id := params["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic("error while parsing")
	}
	models.DeleteBook(ID)
	models.CreateBook(&book)
	res, _ := json.Marshal(&book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		panic("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(&book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}