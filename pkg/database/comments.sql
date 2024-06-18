-- name: CommentByID :one
SELECT *
FROM comments
WHERE id = @id
LIMIT 1;

-- name: InsertComment :exec
INSERT INTO comments (
    content,
    post_id,
    user_id,
    parent_comment_id
) VALUES (
    @content,
    @post_id,
    @user_id,
    @parent_comment_id
);

-- name: UpdateComment :exec
UPDATE comments
SET
    update_at = 'now',
    content = @content
WHERE id = @id;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = @id;

-- name: UpsertCommentVote :exec
INSERT OR REPLACE INTO comment_votes (
    comment_id,
    user_id,
    liked
) VALUES (
    @comment_id,
    @user_id,
    @liked
);
