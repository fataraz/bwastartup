package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// gunakan kata kerja
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error) //user->input.go
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	// mapping sturct input ke struct user
	// simpan struct user melalui repository
	repository Repository // user->repository.go
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// RegisterUserInput => user->input.go
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{} // user->entity.go
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, nil
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "Role"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, err
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// dapatkan user berdasarkan ID
	// update attribute avatar file name
	// simpan perubahan avatar file name

	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
