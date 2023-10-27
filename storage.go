package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
*
CRUD Interface
*/
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		account_number serial,
		account_balance serial,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

type PostgresStore struct {
	db *sql.DB
}

// Crud Methods
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `
	insert into
	account (first_name, last_name, account_number, account_balance)
	values ($1, $2, $3, $4)
	`
	resp, err := s.db.Query(query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
	)

	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

// Get account by ID
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {

	return nil, nil
}

// Get all the accounts from acc table
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`select * from account`)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}

// Database Connections
func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=admin dbname=gobank password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
