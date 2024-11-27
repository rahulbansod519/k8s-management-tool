package repositories

import (
	"context"
	"time"

	"k8s-management-tool/config"
	"k8s-management-tool/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateUser adds a new user to the database
func CreateUser(user models.User) error {
	collection := config.DB.Collection("users") // Initialize collection dynamically
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	return err
}

// FindUserByUsername retrieves a user by username
func FindUserByUsername(username string) (*models.User, error) {
	collection := config.DB.Collection("users") // Initialize collection dynamically
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
