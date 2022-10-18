package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PieterDup98/go-rest-api-course/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author
		FROM comments
		WHERE id = $1`,
		uuid,
	)

	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid")
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) CreateComment(ctx context.Context, newComment comment.Comment) (comment.Comment, error) {
	newComment.ID = uuid.NewV4().String()

	postRow := CommentRow{
		ID:     newComment.ID,
		Slug:   sql.NullString{String: newComment.Slug, Valid: true},
		Body:   sql.NullString{String: newComment.Body, Valid: true},
		Author: sql.NullString{String: newComment.Author, Valid: true},
	}

	row, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, body, author)
		VALUES (:id, :slug, :body, :author)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert a comment: %w", err)
	}
	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return newComment, nil
}
