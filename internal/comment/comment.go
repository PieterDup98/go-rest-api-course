package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImpemented   = errors.New("not implemented")
)

// Comment - a represenation of the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that our service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

// Service - is the struct on which all our
// logic will be build on top of
type Service struct {
	Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment, id:", id)

	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("updating a comment, id:", cmt.ID)

	cmt, err := s.Store.UpdateComment(ctx, cmt)

	if err != nil {
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	fmt.Println("deleting a comment, id:", id)

	err := s.Store.DeleteComment(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("creating a new comment")

	insertedComment, err := s.Store.CreateComment(ctx, cmt)

	if err != nil {
		return Comment{}, err
	}
	return insertedComment, nil
}
