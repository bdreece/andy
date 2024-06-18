CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email_address TEXT NOT NULL,
    email_verified BOOLEAN NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS IX_users_name ON users (last_name ASC, first_name ASC);
CREATE UNIQUE INDEX IF NOT EXISTS IX_users_email_address ON users (email_address ASC);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL, 
    title TEXT NOT NULL,
    description TEXT,
    user_id INTEGER NOT NULL REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS IX_posts_created_at ON posts (created_at DESC);
CREATE INDEX IF NOT EXISTS IX_posts_title ON posts (title ASC);
CREATE INDEX IF NOT EXISTS IX_posts_user ON posts (user_id ASC);

CREATE TABLE IF NOT EXISTS post_votes (
    post_id INTEGER NOT NULL REFERENCES posts (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    liked BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS IX_post_votes_post_user ON post_votes (post_id ASC, user_id ASC);
CREATE INDEX IF NOT EXISTS IX_post_votes_liked ON post_votes (post_id ASC, liked ASC);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL, 
    content TEXT NOT NULL,
    post_id INTEGER NOT NULL REFERENCES posts (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    parent_comment_id INTEGER REFERENCES comments (id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS IX_comments_post_user ON comments (post_id ASC, user_id ASC);
CREATE INDEX IF NOT EXISTS IX_comments_parent_comment ON comments (parent_comment_id ASC);

CREATE TABLE IF NOT EXISTS comment_votes (
    comment_id INTEGER NOT NULL REFERENCES comments (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES comments (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    liked BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS IX_comment_votes_comment_user ON comment_votes (comment_id ASC, user_id ASC);
CREATE INDEX IF NOT EXISTS IX_comment_votes_liked ON comment_votes (comment_id ASC, liked ASC);

CREATE TABLE IF NOT EXISTS resources (
    id INTEGER PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL, 
    kind TEXT NOT NULL,
    file_path TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS IX_resources_kind ON resources (kind ASC);
CREATE INDEX IF NOT EXISTS IX_resources_user ON resources (user_id ASC);

CREATE TABLE IF NOT EXISTS post_resources (
    post_id INTEGER NOT NULL REFERENCES posts (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    resource_id INTEGER NOT NULL REFERENCES resources (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS IX_post_resources_post_resource ON post_resources (post_id, resource_id);
