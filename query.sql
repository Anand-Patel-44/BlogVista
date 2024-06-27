-- name: CreateUser :execresult
INSERT INTO users(username,password,bio,color,profession) VALUES(?,?,?,?,?);

-- name: CreateBlog :execresult
INSERT INTO blogs(title,blogText,username) VALUES(?,?,?);

-- name: UpdateUser :execresult 
UPDATE users SET bio=?,color=?,profession=?,username=? WHERE username=?;

-- name: DeleteBlog :exec
DELETE FROM blogs WHERE blogID=?;

-- name: UpdateBlog :execresult
UPDATE blogs SET title=?,blogText=? WHERE blogId=?;

-- name: SelectBlogs :many
SELECT * FROM blogs;

-- name: SelectBlogByUserName :one
SELECT * FROM blogs WHERE blogID=?;

-- name: SelectUserById :one
SELECT * FROM users WHERE username=?;

-- name: SelectBlogsByUserName :many
SELECT * FROM blogs WHERE username=?;