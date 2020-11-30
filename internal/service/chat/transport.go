package chat

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

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/messages",
		function: getAll(s),
	})
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/message/:id",
		function: getMessageByID(s),
	})
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/message",
		function: postMessage(s),
	})
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/message/:id",
		function: deleteMessage(s),
	})
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/message/:id",
		function: putMessage(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": s.FindAll(),
		})
	}
}

func deleteMessage(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Identificador Invalido.",
			})
		}

		posibleErr := s.DeleteMessage(id)
		if posibleErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Ocurrio un error. No se pudo eliminar el elemento.",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Elemento eliminadao": "El elemento se elimino correctamente.",
		})
	}
}

func getMessageByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Identificador Invalido.",
			})
		}
		message, error := s.FindByID(i)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "No se encontro ningun elemento con dicho ID.",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": *message,
		})
	}
}

//Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

func postMessage(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message Message
		err := c.BindJSON(&message)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Ocurrio un error, intente nuevamente.",
			})
		}
		postErr := s.PostMessage(message)
		if postErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Ocurrio un error, intente nuevamente.",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": "Se inserto el elemento con exito",
		})
	}
}

func putMessage(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message Message
		err := c.BindJSON(&message)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Ocurrio un error, intente nuevamente.",
			})
		}

		id := c.Param("id")
		i, err := strconv.Atoi(id)

		updateErr := s.UpdateMessage(i, message)

		if updateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Ocurrio un error, intente nuevamente.",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": "Se actualizo el elemento con exito",
		})
	}
}
