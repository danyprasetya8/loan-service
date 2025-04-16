package http

import "github.com/gin-gonic/gin"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	r := gin.Default()
	api := r.Group("/api")

	s.v1Route(api)

	r.Run()
}

func (s *Server) v1Route(api *gin.RouterGroup) {
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/")

	borrower := v1.Group("/borrower")
	borrower.GET("/")
	borrower.POST("/")
	borrower.DELETE("/")

	loan := v1.Group("/loan")
	loan.GET("/")
	loan.GET("/:id")

	// Route for FieldOfficer
	loan.POST("/:id")
	loan.POST("/:id/agreement-letter")

	// Route for Internal
	loan.POST("/:id/_approve")
	loan.POST("/:id/proof")
	loan.POST("/:id/_disburse")

	// Route for Investor
	loan.POST("/:id/_invest")
}
