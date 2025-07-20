-- name: CreateUser :exec
INSERT INTO users (user_id, email, name, surname) VALUES (@user_id, @email,@name, @surname);

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = @user_id AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = @email AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE user_id = @user_id;
