package service

import (
	"echo-server/internal/model"
	"echo-server/internal/repository"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("service: invalid credentials")
var ErrUserExist = errors.New("service: can't register this user")

type UserService struct{
	userRepo *repository.UserRepo
}

func NewUserService(r *repository.UserRepo) UserService{
	return UserService{userRepo: r}
}

// Register a new user
func (s *UserService) Register(username, email, password string) (*model.User, error) {
	existing, _ := s.userRepo.GetByUserName(username)
	if existing != nil {
		return nil, ErrUserExist
	} 

	hashedPwd, hashErr := hashPassowrd(password)
	if hashErr != nil {
		return nil, fmt.Errorf("service: hash password: %w", hashErr)
	}
	u := &model.User{
		Username: username,
		Email: email,
		PasswordHash: hashedPwd,
	}

	createUserErr := s.userRepo.CreateUser(u)
	if createUserErr != nil{
		return nil, createUserErr
	} else {
		return u, nil
	}
}

// Login
func (s *UserService) Login(username, password string) (*model.User, error) {
	u, fetchingErr := s.userRepo.GetByUserName(username)

	if fetchingErr != nil {
		if errors.Is(fetchingErr, pgx.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("service: user lookup: %w", fetchingErr)
	}

	pwdErr := checkPassword(u.PasswordHash, password)
	if pwdErr != nil {
		return nil, ErrInvalidCredentials
	} else {
		return u, nil
	}
}


// hashPassword
func hashPassowrd(password string) (string, error) {
	hashed, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return "", hashErr
	} else {
		return string(hashed), nil
	}
}

// checkPassword
func checkPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}