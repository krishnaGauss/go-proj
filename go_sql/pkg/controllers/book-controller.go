package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/krishnaGauss/go-proj/go_sql/pkg/models"
	"github.com/krishnaGauss/go-proj/go_sql/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newbooks := models.GetAllBooks()

	jsondata, _ := json.Marshal(newbooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsondata)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookID"]

	id_num, _ := strconv.ParseInt(id, 0, 0)
	getbooks, _ := models.GetBookById(id_num)

	jsondata, _ := json.Marshal(getbooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsondata)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookID"]
	id_num, _ := strconv.ParseInt(id, 0, 0)

	deleteBook := models.DeleteBook(id_num)

	res, _ := json.Marshal(deleteBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookid := vars["bookID"]
	id, _ := strconv.ParseInt(bookid, 0, 0)

	bookdetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookdetails.Publication = updateBook.Publication
	}

	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}
