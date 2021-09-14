package service

import (
	"errors"

	"estudar.com.br/config/database"
	"estudar.com.br/entity"
)

func ShowAllBooks() ([]entity.Book, error) {
	db := database.GetDatabase()
	var books []entity.Book
	err := db.Find(&books).Error

	if err != nil {
		return nil, errors.New("Error get books")
	}

	return books, nil
}

func GetBooks(id int) (entity.Book, error) {
	db := database.GetDatabase()
	var book entity.Book
	err := db.First(&book, id).Error

	if err != nil {
		return book, errors.New("Book id not found")
	}
	return book, nil
}

func CreateBook(book entity.Book) (entity.Book, error) {
	db := database.GetDatabase()

	err := db.Create(&book).Error
	if err != nil {
		return book, errors.New("Error creating the book")
	}

	return book, nil
}

func EditBook(book entity.Book) (entity.Book, error) {
	_, err := GetBooks(book.ID)
	if err != nil {
		return book, err
	}

	db := database.GetDatabase()

	err = db.Save(&book).Error
	if err != nil {
		return book, errors.New("Error editing the book")
	}

	return book, nil
}

func DeleteBook(id int) error {
	db := database.GetDatabase()

	err := db.Delete(&entity.Book{}, id).Error

	if err != nil {
		return errors.New("Error delete book")
	}
	return nil
}
