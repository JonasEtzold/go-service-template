package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "gitlab.com/JonasEtzold/go-service-template/internal/pkg/models/example"
	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/persistence"
	http_err "gitlab.com/JonasEtzold/go-service-template/pkg/http-err"
)

// GetExampleById godoc
// @Summary      Retrieves example based on given ID
// @Description  get Example by ID
// @Produce      json
// @Param        id   path      integer  true  "Example ID"
// @Success      200  {object}  example.Example
// @Failure      403  {object}  http_err.HTTPError
// @Failure      404  {object}  http_err.HTTPError
// @Failure      500  {object}  http_err.HTTPError
// @Router       /api/v1/example/{id} [get]
// @Security     Authorization Token
func GetExampleById(c *gin.Context) {
	s := persistence.GetExampleRepository()
	id := c.Param("id")
	if example, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("example not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, example)
	}
}

// GetExample godoc
// @Summary      Retrieves example based on query
// @Description  Get Example based on example name or text provided as body JSON
// @Produce      json
// @Param        name  body      string  false  "Name"
// @Param        text  body      string  false  "Text"
// @Success      200   {array}   []example.Example
// @Failure      403   {object}  http_err.HTTPError
// @Failure      404   {object}  http_err.HTTPError
// @Failure      500   {object}  http_err.HTTPError
// @Router       /api/v1/example [get]
// @Security     Authorization Token
func GetExample(c *gin.Context) {
	s := persistence.GetExampleRepository()
	var q models.Example
	_ = c.Bind(&q)
	if examples, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("example not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, examples)
	}
}

// CreateExample godoc
// @Summary      Creates an example
// @Description  Create an example with a name and a descriptive text.
// @Produce      json
// @Param        name  body      string  true   "Name"
// @Param        text  body      string  false  "Text"
// @Success      201   {object}  example.Example
// @Failure      400   {object}  http_err.HTTPError
// @Failure      403   {object}  http_err.HTTPError
// @Failure      500   {object}  http_err.HTTPError
// @Router       /api/v1/example [post]
// @Security     Authorization Token
func CreateExample(c *gin.Context) {
	s := persistence.GetExampleRepository()
	var exampleInput models.Example
	if err := c.BindJSON(&exampleInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := s.Add(&exampleInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, exampleInput)
	}
}

// UpdateExample godoc
// @Summary      Updates an example
// @Description  Updates an existing example with a name and a descriptive text based on the given ID.
// @Produce      json
// @Param        id    path      integer  true   "Example ID"
// @Param        name  body      string   true   "Name"
// @Param        text  body      string   false  "Text"
// @Success      200   {object}  example.Example
// @Failure      403   {object}  http_err.HTTPError
// @Failure      404   {object}  http_err.HTTPError
// @Failure      500   {object}  http_err.HTTPError
// @Router       /api/v1/example [put]
// @Security     Authorization Token
func UpdateExample(c *gin.Context) {
	s := persistence.GetExampleRepository()
	id := c.Params.ByName("id")
	var exampleInput models.Example
	if err := c.BindJSON(&exampleInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("example not found"))
		log.Println(err)
	} else {
		if err := s.Update(&exampleInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, exampleInput)
		}
	}
}

// DeleteExample godoc
// @Summary      Deletes an example
// @Description  Deletes an existing example based on the given ID.
// @Produce      json
// @Param        id   path      integer  true  "Example ID"
// @Success      204  {object}  string
// @Failure      403  {object}  http_err.HTTPError
// @Failure      404  {object}  http_err.HTTPError
// @Failure      500  {object}  http_err.HTTPError
// @Router       /api/v1/example [delete]
// @Security     Authorization Token
func DeleteExample(c *gin.Context) {
	s := persistence.GetExampleRepository()
	id := c.Params.ByName("id")
	if example, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("example not found"))
		log.Println(err)
	} else {
		if err := s.Delete(example); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
