package http

import (
	"loan-service/app/http/handler"

	_ "loan-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{handler}
}

func (s *Server) Run() {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.v1Route(api)

	r.Run()
}

func (s *Server) v1Route(api *gin.RouterGroup) {
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/")

	borrower := v1.Group("/borrower")
	borrower.GET("/", s.handler.GetBorrowers)
	borrower.POST("/", s.handler.CreateBorrower)
	borrower.DELETE("/:id", s.handler.DeleteBorrowerByID)

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
