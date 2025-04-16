package main

import (
	"loan-service/app/http"
	"loan-service/app/http/handler"
	"loan-service/internal/config/database"
	"loan-service/internal/entity"
	borrowerRepo "loan-service/internal/repository/borrower"
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
		&entity.Borrower{},
	)

	br := borrowerRepo.New(db)

	bs := borrowerService.New(br)

	h := handler.New(bs)

	server := http.NewServer(h)
	server.Run()
}
