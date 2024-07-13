package main

import (
	"errors"
	"fmt"
	"hse_mini_course/proto"
	"hse_mini_course/sqlc"
)

func validateName(name string) error {
	if len(name) == 0 {
		return errors.New("name cannot be empty")
	}
	if len(name) > 32 {
		return errors.New("name cannot be longer than 32 symbols")
	}
	for _, ch := range name {
		if !(('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')) {
			return errors.New(fmt.Sprintf("name \"%s\" must consist of only latin letters", name))
		}
	}
	return nil
}

func accountModelToDto(account *sqlc.Account) *proto.GetAccountResponse {
	return &proto.GetAccountResponse{
		Name:    account.Name,
		Balance: account.Balance.Int32,
	}
}
