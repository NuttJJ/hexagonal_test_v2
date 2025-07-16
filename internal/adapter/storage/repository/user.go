package repository

import (
	"context"
	"fmt"
	"hexagonal-test-v2/internal/core/domain"
	"hexagonal-test-v2/internal/core/port"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	coll *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) port.UserRepository {
	return &userRepo{coll: db.Collection("users")}
}

func (r *userRepo) CreateUser(user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.coll.InsertOne(ctx, user)
	return user, err
}

func (r *userRepo) ListUsers() ([]*domain.User, error) {
	var users []*domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.coll.Find(ctx, map[string]interface{}{})
	fmt.Println("Cursor:", cursor)
	fmt.Println("Error:", err)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
