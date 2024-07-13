package main

import (
	"context"
	"hse_mini_course/proto"
	"hse_mini_course/sqlc"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedHw3Server
	queries *sqlc.Queries
}

func newServer(queries *sqlc.Queries) *server {
	return &server{
		queries: queries,
	}
}

func (s *server) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.GetAccountResponse, error) {
	name := request.Name
	err := validateName(name)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	account, err := s.queries.CreateAccount(ctx, sqlc.CreateAccountParams{
		Name:    name,
		Balance: pgtype.Int4{Int32: 0, Valid: true},
	})
  if err != nil {
    problem := IdentifyErrorAfterTransaction(err)
    switch problem {
    case NoRows:
      panic("NoRows is impossible in CreateAccount")
    case Duplicates:
      return nil, NameIsTaken(&name)
    case DBError, NonDBError:
      log.Printf("ERR: while creating account: %v\n", err)
      return nil, status.Error(codes.Internal, "internal error")
    }
    panic("Switch/case was supposed to be exhaustive")
  }

	return accountModelToDto(&account), nil
}

func (s *server) GetAccount(ctx context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	name := request.Name

  account, err := s.queries.GetAccount(ctx, name)
  if err != nil {
    if err == pgx.ErrNoRows { // No error, name not found
      return nil, NameNotFound(&name)
    } else { // Actual db error
      log.Printf("ERR: while getting account: %v\n", err)
      return nil, status.Error(codes.Internal, "internal error")
    }
  }

	return accountModelToDto(&account), nil
}

func (s *server) NewTransaction(ctx context.Context, request *proto.NewTransactionRequest) (*proto.GetAccountResponse, error) {
	name := request.Name
	delta := request.Delta

  account, err := s.queries.UpdateBalance(ctx, sqlc.UpdateBalanceParams{
    Name: name,
    Balance: pgtype.Int4{ Int32: delta, Valid: true },
  })
  if err != nil {
    if err == pgx.ErrNoRows { // No error, name not found
      return nil, NameNotFound(&name)
    } else { // Actual db error
      log.Printf("ERR: while getting account: %v\n", err)
      return nil, status.Error(codes.Internal, "internal error")
    }
  }

	return accountModelToDto(&account), nil
}

func (s *server) ChangeName(ctx context.Context, request *proto.ChangeNameRequest) (*proto.GetAccountResponse, error) {
	name := request.Name
	newName := request.NewName
	err := validateName(newName)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

  account, err := s.queries.UpdateName(ctx, sqlc.UpdateNameParams{
    Name: name,
    Name_2: newName,
  })
  if err != nil {
    problem := IdentifyErrorAfterTransaction(err)
    switch problem {
    case NoRows:
      return nil, NameNotFound(&name)
    case Duplicates:
      return nil, NameIsTaken(&newName)
    case DBError, NonDBError:
      log.Printf("ERR: while creating account: %v\n", err)
      return nil, status.Error(codes.Internal, "internal error")
    }
    panic("Switch/case was supposed to be exhaustive")
  }

	return accountModelToDto(&account), nil
}

func (s *server) DeleteAccount(ctx context.Context, request *proto.DeleteAccountRequest) (*proto.Empty, error) {
	name := request.Name

  err := s.queries.DeleteAccount(ctx, name)
  if err != nil {
    if err == pgx.ErrNoRows { // No error, name not found
      return nil, NameNotFound(&name)
    } else { // Actual db error
      log.Printf("ERR: while getting account: %v\n", err)
      return nil, status.Error(codes.Internal, "internal error")
    }
  }

	return &proto.Empty{}, nil
}
