package database

import (
	"emailapi/app/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type EmailDB interface {
	Open() error
	Close() error
	// recieves email object and create in db
	CreateEmail(e *models.Email) error
	// get email objects from db
	GetEmail() ([]*models.Email, error)
}

type DB struct{
	db *sqlx.DB
}

// Open method opens a new connection to posrgres database
func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to database!")
	pg.MustExec(createSchema)
	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return nil
}

func (d *DB) CreateEmail(e *models.Email) error{
	res, err := d.db.Exec(insertEmailSchema, e.Title, e.Content, e.Author)
	if err != nil{
		return err
	}
	res.LastInsertId()
	return err
}

func (d *DB) GetEmail()([]*models.Email, error){
	
	var emails []*models.Email
	err := d.db.Select(&emails, "SELECT * FROM email")	
	if err != nil{
		return emails, err
	}
	return emails, nil
}