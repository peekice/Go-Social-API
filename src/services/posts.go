package services

import (
	"errors"
	"go-api/domain/entities"
	"go-api/domain/repositories"
	"time"

	"github.com/google/uuid"
)

type postsService struct {
	postsRepository repositories.IPostsRepository
	usersRepository repositories.IUsersRepository
}

type IPostsService interface {
	CreatePost(userID string, payloadData entities.UserPostModel) error
	GetAllPosts() ([]entities.PostDataModel, error)
	EditPost(userID string, postID string, newContent string) error
	GetPostByID(postID string) (entities.PostDataModel, error)
	LikePost(postID string) error
	CommentPost(userID string, postID string, newComment string) error
	DeletePost(userID string, postID string) error
	EditComment(userID string, postID string, commentID string, newContent string) error
	DeleteComment(userID string, postID string, commentID string) error
}

func NewPostsService(repo0 repositories.IPostsRepository, repo1 repositories.IUsersRepository) IPostsService {
	return &postsService{
		postsRepository: repo0,
		usersRepository: repo1,
	}
}

func (sv *postsService) GetAllPosts() ([]entities.PostDataModel, error) {
	posts, err := sv.postsRepository.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (sv *postsService) GetPostByID(postID string) (entities.PostDataModel, error) {

	postData, err := sv.postsRepository.GetPostByID(postID)

	if err != nil {
		return entities.PostDataModel{}, err
	}
	return postData, nil
}

func (sv *postsService) CreatePost(userID string, payloadData entities.UserPostModel) error {
	if payloadData.Content == "" {
		return errors.New("content cannot be empty")
	}

	// var user entities.UserDataModel
	user, err := sv.usersRepository.GetUserByUserID(userID)
	if err != nil {
		return err
	}

	postData := entities.PostDataModel{
		PostID:  uuid.New().String(),
		Content: payloadData.Content,
		PostBy:  user,
		Likes:   0,
		Comment: []entities.Comment{},
		PostdAt: time.Now(),
	}

	err = sv.postsRepository.InsertPost(postData)

	if err != nil {
		return err
	}
	return nil
}

func (sv *postsService) EditPost(userID string, postID string, newContent string) error {

	if newContent == "" {
		return errors.New("content cannot be empty")
	}

	existPost, err := sv.postsRepository.GetPostByID(postID)

	if err != nil {
		return err
	}

	if existPost.PostBy.UserID != userID {
		return errors.New("you don't have permission to edit this post")
	}

	err = sv.postsRepository.EditPost(postID, newContent)
	if err != nil {
		return err
	}
	return nil
}

func (sv *postsService) LikePost(postID string) error {

	_, err := sv.postsRepository.GetPostByID(postID)

	if err != nil {
		return err
	}

	err = sv.postsRepository.LikePost(postID)
	if err != nil {
		return err
	}
	return nil
}

func (sv *postsService) CommentPost(userID string, postID string, newComment string) error {

	if newComment == "" {
		return errors.New("comment cannot be empty")
	}

	_, err := sv.postsRepository.GetPostByID(postID)
	if err != nil {
		return err
	}

	user, err := sv.usersRepository.GetUserByUserID(userID)
	if err != nil {
		return err
	}

	CommentData := entities.Comment{
		CommentID: uuid.New().String(),
		Content:   newComment,
		PostBy:    user,
		CommentAt: time.Now(),
		Edited:    false,
	}

	err = sv.postsRepository.CommentPost(postID, CommentData)
	if err != nil {
		return err
	}

	return nil
}

func (sv *postsService) DeletePost(userID string, postID string) error {

	existPost, err := sv.postsRepository.GetPostByID(postID)
	if err != nil {
		return err
	}

	if existPost.PostBy.UserID != userID {
		return errors.New("you don't have permission to delete this post")
	}

	err = sv.postsRepository.DeletePost(postID)

	if err != nil {
		return err
	}

	return nil
}

func (sv *postsService) EditComment(userID string, postID string, commentID string, newContent string) error {

	existPost, err := sv.postsRepository.GetPostByID(postID)
	if err != nil {
		return err
	}

	for _, comment := range existPost.Comment {
		if comment.CommentID == commentID {
			if comment.PostBy.UserID != userID {
				return errors.New("you don't have permission to edit this comment")
			}
			if newContent == "" {
				return errors.New("content cannot be empty")
			}
			err = sv.postsRepository.EditComment(postID, commentID, newContent)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("comment not found")
}

func (sv *postsService) DeleteComment(userID string, postID string, commentID string) error {

	existPost, err := sv.postsRepository.GetPostByID(postID)

	if err != nil {
		return err
	}

	for _, comment := range existPost.Comment {
		if comment.CommentID == commentID {
			if comment.PostBy.UserID != userID {
				return errors.New("you don't have permission to delete this comment")
			}

			err = sv.postsRepository.DeleteComment(postID, commentID)

			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("comment not found")
}
