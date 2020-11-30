package chat

import (
	"github.com/Nataliavytas/API-GoLang/internal/config"
	"github.com/jmoiron/sqlx"
)

// Message ...
type Message struct {
	ID   int64  `json:"id"`
	Text string `json:"message"`
}

//Service ...
type Service interface {
	AddMessage(Message) error
	DeleteMessage(int) error
	FindByID(int) (*Message, error)
	FindAll() []*Message
	PostMessage(Message) error
	UpdateMessage(id int, message Message) error
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//New ..
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddMessage(m Message) error {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	if err := s.db.Select(&list, "SELECT * FROM messages"); err != nil {
		panic(err)
	}
	return list
}

func (s service) FindByID(ID int) (*Message, error) {
	var message Message

	sentence := `SELECT * FROM messages WHERE id=?;`
	s.db.MustExec(sentence, ID)

	err := s.db.QueryRowx(sentence, ID).StructScan(&message)

	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (s service) PostMessage(message Message) error {
	sentence := "INSERT INTO messages (text) VALUES (?)"
	_, err := s.db.Exec(sentence, message.Text)
	if err != nil {
		return err
	}

	return nil
}

func (s service) DeleteMessage(id int) error {
	sentence := `DELETE FROM messages WHERE id=?;`
	_, err := s.db.Exec(sentence, id)

	if err != nil {
		return err
	}

	return nil
}

func (s service) UpdateMessage(id int, message Message) error {
	sentence := `UPDATE messages SET text = ? WHERE id=?;`
	_, err := s.db.Exec(sentence, message.Text, id)

	if err != nil {
		return err
	}

	return nil
}
