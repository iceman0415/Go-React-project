package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=./mocks/mock_db.go -package=db github.com/bogdanguranda/go-react-example/db DB

// DB describes CRUD methods for a database managing persons.
type DB interface {
	CreatePerson(person *Person) error
	DeletePerson(email string) error
	ListPersons(orderBy string) ([]*Person, error)

	RetrievePerson(email string) (*Person, error)
	UpdatePerson(email string, updatedPerson *Person) error

	Close()
}

// MySqlDB implements DB interface using MySQL as database.
type MySqlDB struct {
	db *sql.DB
}

// NewMySqlDB creates a new MySqlDB.
func NewMySqlDB(pwd string) (*MySqlDB, error) {
	db, err := TryConnect("root:"+pwd+"@tcp(db:3306)/api_db", 3, 5)
	if err != nil {
		return nil, err
	}

	return &MySqlDB{db: db}, nil
}

// TryConnect retry system for connecting to MySQL database.
func TryConnect(dsn string, delay, retries int) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	for ; err != nil && retries > 0; retries-- {
		time.Sleep(time.Second * time.Duration(delay))
		db, err = sql.Open("mysql", dsn)
	}
	return db, err
}

// Close closes the database connection.
func (my *MySqlDB) Close() {
	if err := my.db.Close(); err != nil {
		logrus.Panic("Failed to close MySQL db.")
	}
}

// CreatePerson creates a new person in the DB.
func (my *MySqlDB) CreatePerson(person *Person) error {
	insertQuery, err := my.db.Prepare("INSERT INTO Persons VALUES(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	if _, err := insertQuery.Exec(person.Name, person.Age, person.Balance, person.Email, person.Address); err != nil {
		return err
	}

	return nil
}

// DeletePerson deletes a person from the DB.
func (my *MySqlDB) DeletePerson(email string) error {
	deleteQuery, err := my.db.Prepare("DELETE FROM Persons WHERE Email = ?;")
	if err != nil {
		return err
	}

	if _, err = deleteQuery.Exec(email); err != nil {
		return err
	}

	return nil
}

// ListPersons retrieves all persons sorting by the given column.
func (my *MySqlDB) ListPersons(orderBy string) ([]*Person, error) {
	rows, err := my.db.Query("SELECT * FROM Persons ORDER BY " + orderBy + " DESC;")
	if err != nil {
		return nil, err
	}

	persons := []*Person{}
	for rows.Next() {
		person := &Person{}
		if err := rows.Scan(&person.Name, &person.Age, &person.Balance, &person.Email, &person.Address); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	return persons, nil
}

// RetrievePerson retrieves a specific person by email.
func (my *MySqlDB) RetrievePerson(email string) (*Person, error) {
	selectQuery, err := my.db.Prepare("SELECT * FROM Persons WHERE Email = ?;")
	if err != nil {
		return nil, err
	}

	row := selectQuery.QueryRow(email)
	person := &Person{}
	return person, row.Scan(&person.Name, &person.Age, &person.Balance, &person.Email, &person.Address)
}

// UpdatePerson updates information for a specific person.
func (my *MySqlDB) UpdatePerson(email string, updatedPerson *Person) error {
	updateQuery, err := my.db.Prepare("UPDATE Persons SET Name = ?, Age = ?, Balance = ?, Email = ?, Address = ? WHERE Email = ?;")
	if err != nil {
		return err
	}

	if _, err = updateQuery.Exec(updatedPerson.Name, updatedPerson.Age, updatedPerson.Balance, updatedPerson.Email, updatedPerson.Address, email); err != nil {
		return err
	}

	return nil
}
