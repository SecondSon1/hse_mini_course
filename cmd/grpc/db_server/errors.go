package main

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	DUPLICATE_KEY_CODE = "23505"
)

type ErrorType int

const (
	NoRows ErrorType = iota
	Duplicates
	DBError
	NonDBError
)

func IdentifyErrorAfterTransaction(err error) ErrorType {
	if err == pgx.ErrNoRows {
		return NoRows
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == DUPLICATE_KEY_CODE {
			return Duplicates
		} else {
			return DBError
		}
	} else {
		return NonDBError
	}
}

func NameIsTaken(name *string) error {
	return status.Errorf(
		codes.AlreadyExists,
		"name \"%s\" is already taken",
		*name,
	)
}

func NameNotFound(name *string) error {
	return status.Errorf(
		codes.NotFound,
		"name \"%s\" not found",
		*name,
	)
}
