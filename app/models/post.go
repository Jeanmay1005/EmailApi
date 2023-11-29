package models

type Email struct {
	ID int64 `db:"id"`
	Title string `db:"title"`
	Content string `db:"content"`
	Author string `db:"author"`
}

type JSONEmail struct {
	ID int64 `db:"id"`
	Title string `db:"title"`
	Content string `db:"content"`
	Author string `db:"author"`
}

type EmailRequest struct{
	Title string `db:"title"`
	Content string `db:"content"`
	Author string `db:"author"`
}