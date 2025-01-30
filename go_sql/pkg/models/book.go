package models

import (
	"github.com/krishnaGauss/go-proj/go_sql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
    gorm.Model
    Name        string `gorm:"type:varchar(255);not null" json:"name"`
    Author      string `gorm:"type:varchar(255);not null" json:"author"`
    Publication string `gorm:"type:varchar(255);not null" json:"publication"`
}


func init(){
	config.Connect()
	db=config.GetDB()

	db.AutoMigrate(&Book{})
}


func (b *Book) CreateBook() *Book{
	db.Create(b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book

	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db	
}

func DeleteBook(Id int64) Book{
	var deleteBook Book

	db.Where("ID=?", Id).Delete(&deleteBook)
	return deleteBook
}