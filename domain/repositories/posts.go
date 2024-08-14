package repositories

import (
	"context"
	. "go-api/domain/datasources"
	"go-api/domain/entities"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type postsRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

type IPostsRepository interface {
	InsertPost(data entities.PostDataModel) error
	GetAllPosts() ([]entities.PostDataModel, error)
	GetPostByID(id string) (entities.PostDataModel, error)
	EditPost(postID string, newContent string) error
	LikePost(postID string) error
	CommentPost(postID string, comment entities.Comment) error
	DeletePost(postID string) error
}

func NewPostsRepository(db *MongoDB) IPostsRepository {
	return &postsRepository{
		Context:    db.Context,
		Collection: db.MongoDB.Database(os.Getenv("DATABASE_NAME")).Collection("posts"),
	}
}

func (repo *postsRepository) InsertPost(data entities.PostDataModel) error {
	_, err := repo.Collection.InsertOne(repo.Context, data)
	if err != nil {
		return err
	}
	return nil
}

func (repo *postsRepository) GetAllPosts() ([]entities.PostDataModel, error) {
	var posts []entities.PostDataModel

	cursor, err := repo.Collection.Find(repo.Context, bson.M{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(repo.Context, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo *postsRepository) GetPostByID(id string) (entities.PostDataModel, error) {
	var post entities.PostDataModel
	err := repo.Collection.FindOne(repo.Context, bson.M{"post_id": id}).Decode(&post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (repo *postsRepository) EditPost(postID string, newContent string) error {

	_, err := repo.Collection.UpdateOne(repo.Context, bson.M{"post_id": postID}, bson.M{"$set": bson.M{"content": newContent}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *postsRepository) LikePost(postID string) error {

	_, err := repo.Collection.UpdateOne(repo.Context, bson.M{"post_id": postID}, bson.M{"$inc": bson.M{"likes": 1}})

	if err != nil {
		return err
	}

	return nil
}

func (repo *postsRepository) CommentPost(postID string, comment entities.Comment) error {

	_, err := repo.Collection.UpdateOne(repo.Context, bson.M{"post_id": postID}, bson.M{"$push": bson.M{"comment": comment}})

	if err != nil {
		return err
	}

	return nil
}

func (repo *postsRepository) DeletePost(postID string) error {

	_, err := repo.Collection.DeleteOne(repo.Context, bson.M{"post_id": postID})
	if err != nil {
		return err
	}

	return nil
}
