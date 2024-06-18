-- name: UserByID :one
SELECT *
FROM users
WHERE id = @id
LIMIT 1;

-- name: UserByEmail :one
SELECT *
FROM users
WHERE email_address = @email_address
LIMIT 1;

-- name: InsertUser :exec
INSERT INTO users (
    created_at,
    updated_at,
    first_name,
    last_name,
    email_address,
    email_verified,
    password_hash
) VALUES (
    @created_at,
    @updated_at,
    @first_name,
    @last_name,
    @email_address,
    @email_verified,
    @password_hash
);

-- name: UpdateUser :exec
UPDATE users
SET
    updated_at = 'now',
    first_name = @first_name,
    last_name = @last_name,
    email_address = @email_address,
    email_verified = @email_verified,
    password_hash = @password_hash
WHERE id = @id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @id;
