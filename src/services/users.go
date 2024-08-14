package services

import (
	"errors"
	"go-api/domain/entities"
	"go-api/domain/repositories"
	"go-api/src/middlewares"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type usersService struct {
	UsersRepository repositories.IUsersRepository
}

type IUsersService interface {
	InsertUser(user entities.UserRegisterModel) error
	Login(user entities.UserLoginModel) (string, error)
	GetUserByUserID(userID string) (entities.UserDetailModel, error)
}

func NewUsersService(repo0 repositories.IUsersRepository) IUsersService {
	return &usersService{
		UsersRepository: repo0,
	}
}

func (sv *usersService) InsertUser(user entities.UserRegisterModel) error {
	hashpass, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	userData := entities.UserDataModel{
		UserID:   uuid.New().String(),
		Username: user.Username,
		Email:    user.Email,
		Password: hashpass,
	}

	err = sv.UsersRepository.InsertUser(userData)
	if err != nil {
		return err
	}
	return nil
}

func (sv *usersService) Login(user entities.UserLoginModel) (string, error) {
	userData, err := sv.UsersRepository.GetUserByUsername(user.Username)
	if err != nil {
		return "", err
	}
	if !CheckPasswordHash(user.Password, userData.Password) {
		return "", errors.New("wrong password")
	}
	token, err := middlewares.GenerateJWTToken(userData.UserID)
	if err != nil {
		return "", err
	}
	return *token.Token, nil
}

func (sv *usersService) GetUserByUserID(userID string) (entities.UserDetailModel, error) {
	userData, err := sv.UsersRepository.GetUserByUserID(userID)
	if err != nil {
		return entities.UserDetailModel{}, err
	}
	return userData, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
