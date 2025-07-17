// internal/adapter/storage/mongo/repository/user.go
package repository

import (
	"context"
	"hexagonal-test-v2/internal/core/domain"
	"hexagonal-test-v2/internal/core/port"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []*domain.User
	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"email": email}

	err := r.coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindByGoogleID(googleID string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"google_id": googleID}

	err := r.coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) UpdateUser(user *domain.User) (*domain.User, error) {
	user.UpdatedAt = time.Now()
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := r.coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return user, nil
}
