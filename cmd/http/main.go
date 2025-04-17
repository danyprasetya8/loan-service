package main

import (
	"loan-service/app/http"
	"loan-service/app/http/handler"
	"loan-service/app/http/middleware"
	"loan-service/internal/config/database"
	"loan-service/internal/entity"
	borrowerRepo "loan-service/internal/repository/borrower"
	fileRepo "loan-service/internal/repository/file"
	loanRepo "loan-service/internal/repository/loan"
	loanApprovalRepo "loan-service/internal/repository/loanapproval"
	loanDisbursementRepo "loan-service/internal/repository/loandisbursement"
	loanInvestmentRepo "loan-service/internal/repository/loaninvestment"
	userRepo "loan-service/internal/repository/user"
	authService "loan-service/internal/service/auth"
	borrowerService "loan-service/internal/service/borrower"
	fileService "loan-service/internal/service/file"
	loanService "loan-service/internal/service/loan"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := database.New()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&entity.File{},
		&entity.User{},
		&entity.Borrower{},
		&entity.Loan{},
		&entity.LoanApproval{},
		&entity.LoanInvestment{},
		&entity.LoanDisbursement{},
	)

	fr := fileRepo.New(db)
	br := borrowerRepo.New(db)
	ur := userRepo.New(db)
	lr := loanRepo.New(db)
	lar := loanApprovalRepo.New(db)
	lir := loanInvestmentRepo.New(db)
	ldr := loanDisbursementRepo.New(db)

	fs := fileService.New(fr)
	as := authService.New(ur)
	bs := borrowerService.New(br)
	ls := loanService.New(&loanService.Depedency{
		FileService:          fs,
		UserRepo:             ur,
		BorrowerRepo:         br,
		LoanRepo:             lr,
		LoanApprovalRepo:     lar,
		LoanInvestmentRepo:   lir,
		LoanDisbursementRepo: ldr,
	})

	h := handler.New(as, bs, ls)
	m := middleware.New(as)

	server := http.NewServer(h, m)
	server.Run()
}
