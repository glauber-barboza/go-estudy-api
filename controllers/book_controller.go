package controllers

import (
	"net/http"
	"strconv"

	"estudar.com.br/entity"
	"estudar.com.br/httputil"
	"estudar.com.br/service"
	"github.com/gin-gonic/gin"
)

// @Summary Get all books
// @Description Get all books
// @Tags Book
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Book
// @Failure 417 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /book [get]
func ShowAllBooks(contex *gin.Context) {
	books, err := service.ShowAllBooks()

	if err != nil {
		httputil.ExpectedExceptionErr(contex, err)
		return
	}

	contex.JSON(http.StatusOK, books)
}

// @Summary Get books
// @Description get book by ID
// @Tags Book
// @Accept  json
// @Produce  json
// @Param id path int true "Book id"
// @Success 200 {object} entity.Book
// @Failure 417 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /book/{id} [get]
func ShowBook(contex *gin.Context) {
	id := contex.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		httputil.ExpectedException(contex, "ID has to be integer", err)
		return
	}

	book, err := service.GetBooks(newid)

	if err != nil {
		httputil.ExpectedExceptionErr(contex, err)
		return
	}

	contex.JSON(http.StatusOK, book)
}

// @Summary Create book
// @Description Create books
// @Tags Book
// @Accept  json
// @Produce  json
// @Param book body request.BookReq true "Add book"
// @Success 200 {object} entity.Book
// @Failure 417 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /book [post]
func CreateBook(contex *gin.Context) {
	var book entity.Book

	err := contex.ShouldBindJSON(&book)
	if err != nil {
		httputil.ExpectedException(contex, "json invalid", err)
		return
	}

	book, err = service.CreateBook(book)
	if err != nil {
		httputil.ExpectedExceptionErr(contex, err)
		return
	}

	contex.JSON(http.StatusCreated, book)
}

// @Summary Edit book
// @Description Edit book
// @Tags Book
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body request.BookReq true "Edit book"
// @Success 200 {object} entity.Book
// @Failure 417 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /book/{id} [put]
func EditBook(contex *gin.Context) {
	var book entity.Book

	id := contex.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		httputil.ExpectedException(contex, "id deve ser numérico", err)
		return
	}

	err = contex.ShouldBindJSON(&book)
	if err != nil {
		httputil.ExpectedException(contex, "json invalid", err)
		return
	}

	book.SetId(newid)

	book, err = service.EditBook(book)
	if err != nil {
		httputil.ExpectedExceptionErr(contex, err)
		return
	}

	contex.JSON(http.StatusOK, book)
}

// @Summary Delete book
// @Description Delete books
// @Tags Book
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 204 {object} entity.Book
// @Failure 417 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /book/{id} [delete]
func DeleteBook(contex *gin.Context) {
	id := contex.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		httputil.ExpectedException(contex, "ID has to be integer", err)
		return
	}
	err = service.DeleteBook(newid)
	if err != nil {
		httputil.ExpectedExceptionErr(contex, err)
		return
	}

	contex.Status(http.StatusNoContent)
}
