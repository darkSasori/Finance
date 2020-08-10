package service

import (
	"context"
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/darksasori/finance/pkg/model"
)

type userRepo struct {
	user *model.User
}

func (u *userRepo) Insert(ctx context.Context, user *model.User) error {
	u.user = user
	password, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.user.Password = password
	return nil
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	u.user = user
	return nil
}

func (u *userRepo) Delete(ctx context.Context, user *model.User) error {
	u.user = nil
	return nil
}

func (u *userRepo) FindOne(ctx context.Context, id interface{}) (*model.User, error) {
	return u.user, nil
}

func TestUserSave(t *testing.T) {
	ctx := context.TODO()
	repo := &userRepo{}
	user := model.NewUser("test", "test", "test2", "test")
	service := NewUser(repo)

	if err := service.Save(ctx, user); err.Error() != "Passwords are different" {
		t.Error(err)
	}

	user.Password = []byte("test")
	if err := service.Save(ctx, user); err != nil {
		t.Error(err)
	}

	user.Displayname = "testing"
	if err := service.Save(ctx, user); err != nil {
		t.Error(err)
	}
}

func TestUserLogin(t *testing.T) {
	ctx := context.TODO()
	repo := &userRepo{}
	user := model.NewUser("test", "test", "test", "test")
	service := NewUser(repo)

	_, err := service.Login(ctx, "test", "test")
	if err != nil && err.Error() != "User not found" {
		t.Error("Expected user not found")
	}

	repo.Insert(ctx, user)
	_, err = service.Login(ctx, "test", "test2")
	if err != nil && err.Error() != "User not found" {
		t.Error("Expected user not found")
	}

	_, err = service.Login(ctx, "test", "test")
	if err != nil {
		t.Error(err)
	}
}

func TestCheckToken(t *testing.T) {
	ctx := context.TODO()
	repo := &userRepo{}
	user := model.NewUser("test", "test", "test", "test")
	service := NewUser(repo)
	repo.Insert(ctx, user)

	token, err := service.Login(ctx, "test", "test")
	if err != nil {
		t.Error(err)
	}

	user2, err := service.CheckToken(ctx, fmt.Sprintf("%stoken", token))
	if err == nil {
		t.Error("Expected an error")
	}

	user2, err = service.CheckToken(ctx, token)
	if err != nil {
		t.Error(err)
	}

	if user.Username != user2.Username {
		t.Errorf("Expected %s, received %s", user.Username, user2.Username)
	}
}
