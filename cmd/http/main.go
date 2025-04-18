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
	mailerService "loan-service/internal/service/mailer"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	logSetup()

	if err := godotenv.Load(); err != nil {
		log.Error(err.Error())
		panic(err)
	}

	db, err := database.New()

	if err != nil {
		log.Error(err.Error())
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

	ms := mailerService.New()
	fs := fileService.New(fr)
	as := authService.New(ur)
	bs := borrowerService.New(br)
	ls := loanService.New(&loanService.Dependency{
		FileService:          fs,
		MailerService:        ms,
		UserRepo:             ur,
		BorrowerRepo:         br,
		LoanRepo:             lr,
		LoanApprovalRepo:     lar,
		LoanInvestmentRepo:   lir,
		LoanDisbursementRepo: ldr,
	})

	h := handler.New(fs, as, bs, ls)
	m := middleware.New(as)

	server := http.NewServer(h, m)
	server.Run()
}

func logSetup() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
