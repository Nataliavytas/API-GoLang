package library

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

//NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

//Cretion of the endpoints
func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/books",
		function: getAll(s),
	})
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/book/:id",
		function: getBookByID(s),
	})
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/book",
		function: postBook(s),
	})
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/book/:id",
		function: deleteBook(s),
	})
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/book/:id",
		function: putBook(s),
	})

	return list
}

//Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

//Get all the books in the database
func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"books": s.FindAll(),
		})
	}
}

//Get a book by id
func getBookByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Identificador Invalido.",
			})
		}
		book, error := s.FindByID(i)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No se encontro ningun libro con dicho ID.",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"book": *book,
		})
	}
}

//Post a book
func postBook(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book Book
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ocurrio un error, intente nuevamente.",
			})
		}
		postErr := s.PostBook(book)
		if postErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ocurrio un error, intente nuevamente.",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Se inserto el libro con exito",
		})
	}
}

//Delete a book
func deleteBook(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Identificador Invalido.",
			})
		}

		posibleErr := s.DeleteBook(id)
		if posibleErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ocurrio un error. No se pudo eliminar el libro.",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "El libro se elimino correctamente.",
		})
	}
}

//Update a book
func putBook(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book Book
		err := c.BindJSON(&book)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ocurrio un error, intente nuevamente.",
			})
		}

		id := c.Param("id")
		i, err := strconv.Atoi(id)

		updateErr := s.UpdateBook(i, book)

		if updateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ocurrio un error, intente nuevamente.",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Se actualizo el libro con exito",
		})
	}
}
