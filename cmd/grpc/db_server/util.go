package main

import (
	"errors"
	"fmt"
	"hse_mini_course/accounts/models"
	"hse_mini_course/proto"
	"math/rand/v2"
)

func generateId() uint32 {
	return rand.Uint32N(IDS_TO-IDS_FROM+1) + IDS_FROM
}

func validateName(name string) error {
	if len(name) == 0 {
		return errors.New("name cannot be empty")
	}
	for _, ch := range name {
		if !(('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')) {
			return errors.New(fmt.Sprintf("name \"%s\" must consist of only latin letters", name))
		}
	}
	return nil
}

func modelToResponse(account *models.Account) *proto.GetAccountResponse {
	return &proto.GetAccountResponse{
		Name:    account.Name,
		Balance: account.Balance,
	}
}
