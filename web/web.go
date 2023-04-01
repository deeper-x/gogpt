package web

import (
	"net/http"

	"github.com/deeper-x/gogpt/ai"
	"github.com/gin-gonic/gin"
)

// Server core object
type Server struct {
	Eng *gin.Engine
}

// NewServer return main instance
func NewServer() *Server {
	return &Server{
		Eng: gin.Default(),
	}
}

// GET server REST API wrapper
func (s *Server) GET(p string, h gin.HandlerFunc) {
	s.Eng.GET(p, h)
}

// Run wrapper for engin run call
func (s *Server) Run() {
	s.Eng.Run(":8989")
}

// Query handler return main query
var Query = func(c *gin.Context) {
	key := c.Param("key")
	query := c.Param("query")

	e := ai.NewEngine(key)
	res, err := e.Query(query)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"AI reply": res,
	})
}
