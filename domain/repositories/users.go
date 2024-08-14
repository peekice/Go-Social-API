package repositories

import (
	"context"
	"errors"
	. "go-api/domain/datasources"
	"go-api/domain/entities"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IUsersRepository interface {
	InsertUser(data entities.UserDataModel) error
	GetUserByUsername(username string) (entities.UserDataModel, error)
	GetUserByUserID(userID string) (entities.UserDetailModel, error)
}

func NewUsersRepository(db *MongoDB) IUsersRepository {
	return &usersRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("users"),
	}
}

func (repo *usersRepository) InsertUser(data entities.UserDataModel) error {
	var user entities.UserDataModel
	err := repo.Collection.FindOne(repo.Context, bson.M{"username": data.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			if _, err := repo.Collection.InsertOne(repo.Context, data); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("user already exists")
}

func (repo *usersRepository) GetUserByUsername(username string) (entities.UserDataModel, error) {
	var user entities.UserDataModel
	err := repo.Collection.FindOne(repo.Context, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *usersRepository) GetUserByUserID(userID string) (entities.UserDetailModel, error) {
	var user entities.UserDetailModel

	err := repo.Collection.FindOne(repo.Context, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
