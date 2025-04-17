package auth

import (
	"errors"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/internal/repository/user"
	"loan-service/pkg/model/request"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

type IAuthService interface {
	GetUserRole(email string) (constant.UserRole, error)
	MockLogin(req *request.MockLogin) (string, error)
	ParseToken(tokenStr string) (string, error)
	generateAccessToken(email string) (string, error)
}

type Auth struct {
	userRepo user.IUserRepository
}

func New(userRepo user.IUserRepository) IAuthService {
	return &Auth{userRepo}
}

func (a *Auth) GetUserRole(email string) (constant.UserRole, error) {
	existUser := a.userRepo.Get(email)

	if existUser == nil {
		return "", errors.New("user not exist")
	}

	return existUser.Role, nil
}

func (a *Auth) MockLogin(req *request.MockLogin) (string, error) {
	existUser := a.userRepo.Get(req.Email)

	if existUser == nil {
		err := a.userRepo.Create(&entity.User{
			Email: req.Email,
			Role:  req.Role,
			Audit: entity.Audit{
				CreatedBy: req.Email,
				UpdatedBy: req.Email,
			},
		})

		if err != nil {
			log.Errorf("Error creating user: %s", err.Error())
			return "", err
		}
	} else {
		existUser.Role = req.Role
		a.userRepo.Save(existUser)
	}

	return a.generateAccessToken(req.Email)
}

func (a *Auth) generateAccessToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Errorf("Error signing token: %s", err.Error())
		return "", err
	}

	return t, nil
}

func (a *Auth) ParseToken(tokenStr string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		log.Errorf("Error parsing token: %s", err.Error())
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims["email"].(string), nil
}
