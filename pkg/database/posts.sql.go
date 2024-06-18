// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package database

import (
	"context"
)

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?1
`

type DeletePostParams struct {
	ID int64
}

func (q *Queries) DeletePost(ctx context.Context, arg DeletePostParams) error {
	_, err := q.db.ExecContext(ctx, deletePost, arg.ID)
	return err
}

const insertPost = `-- name: InsertPost :exec
INSERT INTO posts (
    title,
    description,
    user_id
) VALUES (
    ?1,
    ?2,
    ?3
)
`

type InsertPostParams struct {
	Title       string
	Description *string
	UserID      int64
}

func (q *Queries) InsertPost(ctx context.Context, arg InsertPostParams) error {
	_, err := q.db.ExecContext(ctx, insertPost, arg.Title, arg.Description, arg.UserID)
	return err
}

const postByID = `-- name: PostByID :one
SELECT id, created_at, updated_at, title, description, user_id
FROM posts
WHERE id = ?1
LIMIT 1
`

type PostByIDParams struct {
	ID int64
}

func (q *Queries) PostByID(ctx context.Context, arg PostByIDParams) (*Post, error) {
	row := q.db.QueryRowContext(ctx, postByID, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.UserID,
	)
	return &i, err
}

const postsByRating = `-- name: PostsByRating :many
SELECT p.id, p.created_at, p.updated_at, p.title, p.description, p.user_id
FROM posts p
INNER JOIN (
    SELECT
        post_id,
        COUNT(*) AS rating
    FROM post_votes
    WHERE liked = TRUE
    GROUP BY post_id
) r ON r.post_id = p.id
ORDER BY r.rating
`

func (q *Queries) PostsByRating(ctx context.Context) ([]*Post, error) {
	rows, err := q.db.QueryContext(ctx, postsByRating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Description,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :exec
UPDATE posts
SET
    updated_at = 'now',
    title = ?1,
    description = ?2
WHERE id = ?3
`

type UpdatePostParams struct {
	Title       string
	Description *string
	ID          int64
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.ExecContext(ctx, updatePost, arg.Title, arg.Description, arg.ID)
	return err
}

const upsertPostVote = `-- name: UpsertPostVote :exec
INSERT OR REPLACE INTO post_votes (
    post_id,
    user_id,
    liked
) VALUES (
    ?1,
    ?2,
    ?3
)
`

type UpsertPostVoteParams struct {
	PostID int64
	UserID int64
	Liked  bool
}

func (q *Queries) UpsertPostVote(ctx context.Context, arg UpsertPostVoteParams) error {
	_, err := q.db.ExecContext(ctx, upsertPostVote, arg.PostID, arg.UserID, arg.Liked)
	return err
}
