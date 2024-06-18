-- name: PostByID :one
SELECT *
FROM posts
WHERE id = @id
LIMIT 1;

-- name: PostsByRating :many
SELECT p.*
FROM posts p
INNER JOIN (
    SELECT
        post_id,
        COUNT(*) AS rating
    FROM post_votes
    WHERE liked = TRUE
    GROUP BY post_id
) r ON r.post_id = p.id
ORDER BY r.rating;

-- name: InsertPost :exec
INSERT INTO posts (
    title,
    description,
    user_id
) VALUES (
    @title,
    @description,
    @user_id
);

-- name: UpdatePost :exec
UPDATE posts
SET
    updated_at = 'now',
    title = @title,
    description = @description
WHERE id = @id;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = @id;

-- name: UpsertPostVote :exec
INSERT OR REPLACE INTO post_votes (
    post_id,
    user_id,
    liked
) VALUES (
    @post_id,
    @user_id,
    @liked
);
