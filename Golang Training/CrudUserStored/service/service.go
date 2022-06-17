package service

import (
	"UserServicePractice/DataUser"
	"errors"
	"net/mail"
)

type UserServiceInterface interface{
	RegisterService(user *DataUser.User) (*DataUser.User, error)
}

type UserService struct{
}

func NewUserService() UserServiceInterface{
	return &UserService{}
}

func (s UserService) RegisterService(user *DataUser.User) (*DataUser.User, error){
	email := user.Email
	_,err := mail.ParseAddress(email)

	if user.Email == ""{
		return nil, errors.New("email harus diisi")
	}
	if err != nil{
		return nil, errors.New("email tidak valid")
	}
	if user.Username == ""{
		return nil, errors.New("username harus diisi")
	}
	if user.Password == "" || len(user.Password) < 6{
		return nil, errors.New("password harus diisi dan harus lebih dari 6 karakter")
	}
	if user.Age <= 8{
		return nil, errors.New("umur tidak boleh di bawah 8 tahun")
	}
	return user, nil
}