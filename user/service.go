package user

import "golang.org/x/crypto/bcrypt"

// gunakan kata kerja
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error) //user->input.go
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
