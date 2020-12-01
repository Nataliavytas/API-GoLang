package library

import (
	"github.com/Nataliavytas/API-GoLang/internal/config"
	"github.com/jmoiron/sqlx"
)

// Book struct
type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

//Service interface
type Service interface {
	FindByID(int) (*Book, error)
	FindAll() []*Book
	PostBook(Book) error
	UpdateBook(id int, book Book) error
	DeleteBook(int) error
}

//service struct
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//New method for service creation
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

//Finds every book in the library table
func (s service) FindAll() []*Book {
	var list []*Book
	if err := s.db.Select(&list, "SELECT * FROM library"); err != nil {
		panic(err)
	}
	return list
}

//Find the book with the given id
func (s service) FindByID(ID int) (*Book, error) {
	var book Book

	sentence := `SELECT * FROM library WHERE id=?;`
	s.db.MustExec(sentence, ID)

	err := s.db.QueryRowx(sentence, ID).StructScan(&book)

	if err != nil {
		return nil, err
	}
	return &book, nil
}

//Saves the book given in the database
func (s service) PostBook(book Book) error {
	sentence := "INSERT INTO library (title, author, price) VALUES (?, ?, ?)"
	_, err := s.db.Exec(sentence, book.Title, book.Author, book.Price)
	if err != nil {
		return err
	}

	return nil
}

//Delete de book with the id given
func (s service) DeleteBook(id int) error {
	sentence := `DELETE FROM library WHERE id=?;`
	_, err := s.db.Exec(sentence, id)

	if err != nil {
		return err
	}

	return nil
}

//Updates the book with the id given and the new information
func (s service) UpdateBook(id int, book Book) error {
	sentence := `UPDATE library SET title = ?, author = ?, price = ? WHERE id=?;`
	_, err := s.db.Exec(sentence, book.Title, book.Author, book.Price, id)

	if err != nil {
		return err
	}

	return nil
}
