package dto

import "hse_mini_course/accounts/models"

type GetAccountResponse struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func AccountToResponse(account models.Account) GetAccountResponse {
	return GetAccountResponse{
		Name:    account.Name,
		Balance: account.Balance,
	}
}
