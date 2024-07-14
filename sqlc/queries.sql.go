// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (name, balance)
VALUES ($1, $2) RETURNING id, name, balance
`

type CreateAccountParams struct {
	Name    string
	Balance pgtype.Int4
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Name, arg.Balance)
	var i Account
	err := row.Scan(&i.ID, &i.Name, &i.Balance)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :one
DELETE FROM accounts
WHERE name=$1 RETURNING id
`

func (q *Queries) DeleteAccount(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, deleteAccount, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, balance FROM accounts
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, name string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, name)
	var i Account
	err := row.Scan(&i.ID, &i.Name, &i.Balance)
	return i, err
}

const updateBalance = `-- name: UpdateBalance :one
UPDATE accounts SET
  balance = balance + $2
WHERE name=$1 RETURNING id, name, balance
`

type UpdateBalanceParams struct {
	Name    string
	Balance pgtype.Int4
}

func (q *Queries) UpdateBalance(ctx context.Context, arg UpdateBalanceParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateBalance, arg.Name, arg.Balance)
	var i Account
	err := row.Scan(&i.ID, &i.Name, &i.Balance)
	return i, err
}

const updateName = `-- name: UpdateName :one
UPDATE accounts SET
  name = $2
WHERE name=$1 RETURNING id, name, balance
`

type UpdateNameParams struct {
	Name   string
	Name_2 string
}

func (q *Queries) UpdateName(ctx context.Context, arg UpdateNameParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateName, arg.Name, arg.Name_2)
	var i Account
	err := row.Scan(&i.ID, &i.Name, &i.Balance)
	return i, err
}
