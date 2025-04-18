package http

import (
	"loan-service/app/http/handler"
	"loan-service/app/http/middleware"
	"loan-service/internal/constant"

	_ "loan-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handler    *handler.Handler
	middleware *middleware.Middleware
}

func NewServer(
	handler *handler.Handler,
	middleware *middleware.Middleware,
) *Server {
	return &Server{handler, middleware}
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
	auth.GET("/user", s.handler.GetAllUsers)
	auth.POST("/mock-login", s.handler.MockLogin)

	file := v1.Group("/file", s.middleware.Authenticate)
	file.POST("/:id/_download", s.handler.DownloadFile)

	s.borrowerRoute(v1)
	s.loanRoute(v1)
}

func (s *Server) borrowerRoute(v1 *gin.RouterGroup) {
	borrower := v1.Group("/borrower", s.middleware.Authenticate)
	borrower.GET(
		"/",
		s.middleware.Authorize(constant.FieldOfficer),
		s.handler.GetBorrowers)
	borrower.POST(
		"/",
		s.middleware.Authorize(constant.FieldOfficer),
		s.handler.CreateBorrower)
	borrower.DELETE(
		"/:id",
		s.middleware.Authorize(constant.FieldOfficer),
		s.handler.DeleteBorrowerByID)
}

func (s *Server) loanRoute(v1 *gin.RouterGroup) {
	loan := v1.Group("/loan", s.middleware.Authenticate)

	loan.GET("/", s.handler.GetLoans)
	loan.GET("/:id", s.handler.GetLoanDetail)

	loan.POST(
		"/",
		s.middleware.Authorize(constant.FieldOfficer),
		s.handler.ProposeLoan)

	loan.POST(
		"/:id/_approve",
		s.middleware.Authorize(constant.Internal),
		s.handler.ApproveLoan)
	loan.POST(
		"/:id/proof",
		s.middleware.Authorize(constant.Internal),
		s.handler.UploadLoanProofOfPicture)
	loan.POST(
		"/:id/_disburse",
		s.middleware.Authorize(constant.Internal),
		s.handler.DisburseLoan)
	loan.POST(
		"/:id/borrower-agreement-letter",
		s.middleware.Authorize(constant.Internal),
		s.handler.UploadBorrowerAgreementLetter)

	loan.POST(
		"/:id/_invest",
		s.middleware.Authorize(constant.Investor),
		s.handler.InvestLoan)
}
