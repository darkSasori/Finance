package service

import (
	"context"
	"fmt"

	"github.com/darksasori/finance/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

// NewUser return user service
func NewUser(repository UserRepository) *User {
	return &User{repository}
}

// User service
type User struct {
	repo UserRepository
}

// Save user
func (u *User) Save(ctx context.Context, user *model.User) error {
	userSaved, err := u.repo.FindOne(ctx, user)
	if err != nil {
		return err
	}
	if userSaved == nil {
		if !user.CheckPassword() {
			return fmt.Errorf("Passwords are different")
		}

		password, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = password

		return u.repo.Insert(ctx, user)
	}

	return u.repo.Update(ctx, user)
}

// Login return user using username and password
func (u *User) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.repo.FindOne(ctx, username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("User not found")
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return "", fmt.Errorf("User not found")
	}
	return encode(user.Username)
}

// CheckToken verify token
func (u *User) CheckToken(ctx context.Context, token string) (*model.User, error) {
	return nil, nil
}
