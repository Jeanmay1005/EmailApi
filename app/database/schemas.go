package database

const createSchema = `
CREATE TABLE IF NOT EXISTS email
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	author TEXT
)`

var insertEmailSchema = `
	INSERT INTO email (title, content, author) VALUES ($1, $2, $3) RETURNING ID`