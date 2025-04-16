package main

import (
	"loan-service/app/http"
	"loan-service/app/http/handler"
	"loan-service/app/http/middleware"
	"loan-service/internal/config/database"
	"loan-service/internal/entity"
	borrowerRepo "loan-service/internal/repository/borrower"
	userRepo "loan-service/internal/repository/user"
	authService "loan-service/internal/service/auth"
	borrowerService "loan-service/internal/service/borrower"

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
		&entity.User{},
		&entity.Borrower{},
	)

	br := borrowerRepo.New(db)
	ur := userRepo.New(db)

	as := authService.New(ur)
	bs := borrowerService.New(br)

	h := handler.New(as, bs)
	m := middleware.New(as)

	server := http.NewServer(h, m)
	server.Run()
}
