package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// interface ini berupa business login mewakili kata kerja
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error) //mengembalikan objek User
	Login(input LoginInput) (User, error)
}

// dependensi atau kebergantungan kepada repository
type service struct {
	// tujuannnya untuk mengakses method di interface Repository
	repo Repository // repository dengan tipe Repository interface
}

// akan dibuat new service 
func NewService(repository Repository) *service{
	return &service{repo: repository}
}

//mengembalikan objek User
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	//panggil repository method save
	newUser, err := s.repo.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, err
}

func (s *service) Login(input LoginInput) (User, error) {
	// olah data yang dikirim dari postman
	email := input.Email
	password := input.Password

	//panggil repository apakah ada user dengan email yang diinput
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	// bandingkan passwordhash di database dengan password input
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return user, err
	}

	return user, nil
}