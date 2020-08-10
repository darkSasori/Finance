package mongodb

import (
	"context"
	"testing"

	"github.com/darksasori/finance/pkg/model"
)

func TestUserFindOne(t *testing.T) {
	ctx := context.TODO()
	if err := Connect(ctx); err != nil {
		t.Error(err)
		return
	}

	user := model.NewUser("test", "test", "test", "test")
	repo := NewUser()

	u, err := repo.FindOne(ctx, user.Username)
	if err != nil {
		t.Error(err)
	}
	if u != nil {
		t.Error("Expected return nil")
	}

	if err := repo.Insert(ctx, user); err != nil {
		t.Error(err)
	}
	u, err = repo.FindOne(ctx, user.Username)
	if err != nil {
		t.Error(err)
	}
	if u.Username != user.Username {
		t.Error("Expected return same user")
	}

	if err := repo.Delete(ctx, user); err != nil {
		t.Error(err)
	}
}

func TestUserInsert(t *testing.T) {
	ctx := context.TODO()
	if err := Connect(ctx); err != nil {
		t.Error(err)
		return
	}

	user := model.NewUser("test", "test", "test", "test")
	repo := NewUser()
	if err := repo.Insert(ctx, user); err != nil {
		t.Error(err)
	}

	if err := repo.Insert(ctx, user); err == nil {
		t.Error("Excepted duplicate key error")
	}
}

func TestUserUpdate(t *testing.T) {
	ctx := context.TODO()
	if err := Connect(ctx); err != nil {
		t.Error(err)
		return
	}

	user := model.NewUser("test", "test", "test", "test")
	repo := NewUser()

	user.Displayname = "testing"
	if err := repo.Update(ctx, user); err != nil {
		t.Error(err)
	}

	user2, err := repo.FindOne(ctx, "test")
	if err != nil {
		t.Error(err)
	}
	if user.Displayname != user2.Displayname {
		t.Error("Excepted same displayname")
	}
}

func TestUserDelete(t *testing.T) {
	ctx := context.TODO()
	if err := Connect(ctx); err != nil {
		t.Error(err)
		return
	}

	user := model.NewUser("test", "test", "test", "test")
	repo := NewUser()

	if err := repo.Delete(ctx, user); err != nil {
		t.Error(err)
	}
}
