package handlers

import (
	"errors"
	"fmt"
	e "golang/src/web-api-gin/entities"
	m "golang/src/web-api-gin/models"
	s "golang/src/web-api-gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService s.Service
}

func NewBookHandler(bookService s.Service) *bookHandler {
	return &bookHandler{bookService}
}

func convertToBookResponse(book e.Book) m.BookResponse {
	return m.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		Description: book.Description,
		Rating:      book.Rating,
		Discount:    book.Discount,
	}
}

func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []m.BookResponse

	for _, book := range books {
		bookResponse := m.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Price:       book.Price,
			Description: book.Description,
			Discount:    book.Discount,
		}
		booksResponse = append(booksResponse, bookResponse)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := m.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		Discount:    book.Discount,
		Description: book.Description,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Pandega Haqqi Sadieda",
		"bio":  "Karyawan Swasta",
	})
}

func (h *bookHandler) CreateBook(ctx *gin.Context) {
	var bookRequest m.BookCreateRequest

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {

			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf(
					"Error on field %s, condition %s",
					e.Field(),
					e.ActualTag(),
				)
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *bookHandler) UpdateBook(ctx *gin.Context) {
	var bookRequest m.BookUpdateRequest

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {

			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf(
					"Error on field %s, condition %s",
					e.Field(),
					e.ActualTag(),
				)
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *bookHandler) DeleteBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Delete(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(book)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}
