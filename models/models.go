package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

// SaveUser inserts user into MongoDB
func SaveUser(user User) error {
	collection := GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	return err
}

// FindUserByUsername returns user by username from MongoDB
func FindUserByUsername(username string) (*User, error) {
	collection := GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	var user User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
