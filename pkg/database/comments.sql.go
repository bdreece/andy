// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: comments.sql

package database

import (
	"context"
)

const commentByID = `-- name: CommentByID :one
SELECT id, created_at, updated_at, content, post_id, user_id, parent_comment_id
FROM comments
WHERE id = ?1
LIMIT 1
`

type CommentByIDParams struct {
	ID int64
}

func (q *Queries) CommentByID(ctx context.Context, arg CommentByIDParams) (*Comment, error) {
	row := q.db.QueryRowContext(ctx, commentByID, arg.ID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.PostID,
		&i.UserID,
		&i.ParentCommentID,
	)
	return &i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = ?1
`

type DeleteCommentParams struct {
	ID int64
}

func (q *Queries) DeleteComment(ctx context.Context, arg DeleteCommentParams) error {
	_, err := q.db.ExecContext(ctx, deleteComment, arg.ID)
	return err
}

const insertComment = `-- name: InsertComment :exec
INSERT INTO comments (
    content,
    post_id,
    user_id,
    parent_comment_id
) VALUES (
    ?1,
    ?2,
    ?3,
    ?4
)
`

type InsertCommentParams struct {
	Content         string
	PostID          int64
	UserID          int64
	ParentCommentID *int64
}

func (q *Queries) InsertComment(ctx context.Context, arg InsertCommentParams) error {
	_, err := q.db.ExecContext(ctx, insertComment,
		arg.Content,
		arg.PostID,
		arg.UserID,
		arg.ParentCommentID,
	)
	return err
}

const updateComment = `-- name: UpdateComment :exec
UPDATE comments
SET
    update_at = 'now',
    content = ?1
WHERE id = ?2
`

type UpdateCommentParams struct {
	Content string
	ID      int64
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) error {
	_, err := q.db.ExecContext(ctx, updateComment, arg.Content, arg.ID)
	return err
}

const upsertCommentVote = `-- name: UpsertCommentVote :exec
INSERT OR REPLACE INTO comment_votes (
    comment_id,
    user_id,
    liked
) VALUES (
    ?1,
    ?2,
    ?3
)
`

type UpsertCommentVoteParams struct {
	CommentID int64
	UserID    int64
	Liked     bool
}

func (q *Queries) UpsertCommentVote(ctx context.Context, arg UpsertCommentVoteParams) error {
	_, err := q.db.ExecContext(ctx, upsertCommentVote, arg.CommentID, arg.UserID, arg.Liked)
	return err
}