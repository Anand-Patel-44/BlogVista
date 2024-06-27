// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package api

import (
	"context"
	"database/sql"
)

const createBlog = `-- name: CreateBlog :execresult
INSERT INTO blogs(title,blogText,username) VALUES(?,?,?)
`

type CreateBlogParams struct {
	Title    string
	Blogtext string
	Username string
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createBlog, arg.Title, arg.Blogtext, arg.Username)
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users(username,password,bio,color,profession) VALUES(?,?,?,?,?)
`

type CreateUserParams struct {
	Username   string
	Password   string
	Bio        string
	Color      string
	Profession string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.Bio,
		arg.Color,
		arg.Profession,
	)
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blogs WHERE blogID=?
`

func (q *Queries) DeleteBlog(ctx context.Context, blogid int32) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, blogid)
	return err
}

const selectBlogByUserName = `-- name: SelectBlogByUserName :one
SELECT blogid, title, blogtext, username FROM blogs WHERE blogID=?
`

func (q *Queries) SelectBlogByUserName(ctx context.Context, blogid int32) (Blog, error) {
	row := q.db.QueryRowContext(ctx, selectBlogByUserName, blogid)
	var i Blog
	err := row.Scan(
		&i.Blogid,
		&i.Title,
		&i.Blogtext,
		&i.Username,
	)
	return i, err
}

const selectBlogs = `-- name: SelectBlogs :many
SELECT blogid, title, blogtext, username FROM blogs
`

func (q *Queries) SelectBlogs(ctx context.Context) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, selectBlogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.Blogid,
			&i.Title,
			&i.Blogtext,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectBlogsByUserName = `-- name: SelectBlogsByUserName :many
SELECT blogid, title, blogtext, username FROM blogs WHERE username=?
`

func (q *Queries) SelectBlogsByUserName(ctx context.Context, username string) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, selectBlogsByUserName, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.Blogid,
			&i.Title,
			&i.Blogtext,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserById = `-- name: SelectUserById :one
SELECT username, password, bio, color, profession FROM users WHERE username=?
`

func (q *Queries) SelectUserById(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, selectUserById, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Password,
		&i.Bio,
		&i.Color,
		&i.Profession,
	)
	return i, err
}

const updateBlog = `-- name: UpdateBlog :execresult
UPDATE blogs SET title=?,blogText=? WHERE blogId=?
`

type UpdateBlogParams struct {
	Title    string
	Blogtext string
	Blogid   int32
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateBlog, arg.Title, arg.Blogtext, arg.Blogid)
}

const updateUser = `-- name: UpdateUser :execresult
UPDATE users SET bio=?,color=?,profession=?,username=? WHERE username=?
`

type UpdateUserParams struct {
	Bio        string
	Color      string
	Profession string
	Username   string
	Username_2 string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUser,
		arg.Bio,
		arg.Color,
		arg.Profession,
		arg.Username,
		arg.Username_2,
	)
}
