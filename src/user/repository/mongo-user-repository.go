package repository

import (
	"context"
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoUserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(database *mongo.Database) *mongoUserRepository {
	return &mongoUserRepository{database}
}

func (m *mongoUserRepository) FindById(id string) (domain.User, bool, error) {
	var user domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := m.db.Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, false, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, false, domain.ErrNotFound
		}
		return user, false, err
	}
	return user, true, nil
}

func (m *mongoUserRepository) FindByEmail(email string) (domain.User, bool, error) {
	var user domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := m.db.Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, false, domain.ErrNotFound
		}
		return user, false, err
	}
	return user, true, nil
}

func (m *mongoUserRepository) Create(name, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := m.db.Collection("users")
	result, err := collection.InsertOne(ctx, bson.M{"name": name, "email": email, "password": password})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", domain.ErrNotFound
		}
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).String()
	return id, nil
}
