package dto

type CreateAccountRequest struct {
	Name string `json:"name"`
}

type ChangeNameRequest struct {
	NewName string `json:"new_name"`
}

type NewTransactionRequest struct {
	Delta int32 `json:"delta"`
}
