package mongodb

import (
	"context"

	"github.com/darksasori/finance/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewUser create a user repository
func NewUser() *User {
	return &User{conn.Collection("user")}
}

// User is a repository to handle User model
type User struct {
	coll *mongo.Collection
}

// Insert user
func (u *User) Insert(ctx context.Context, user *model.User) error {
	if _, err := u.coll.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

// Update user
func (u *User) Update(ctx context.Context, user *model.User) error {
	q := bson.M{"_id": user.Username}
	update := bson.D{
		{"$set", bson.D{
			{"displayname", user.Displayname},
			{"password", user.Password},
		}},
	}
	_, err := u.coll.UpdateOne(ctx, q, update)
	return err
}

// Delete user
func (u *User) Delete(ctx context.Context, user *model.User) error {
	q := bson.M{"_id": user.Username}
	_, err := u.coll.DeleteOne(ctx, q)
	return err
}

// FindOne return model.User
func (u *User) FindOne(ctx context.Context, id interface{}) (*model.User, error) {
	var result model.User
	q := bson.M{"_id": id}
	if err := u.coll.FindOne(ctx, q).Decode(&result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
