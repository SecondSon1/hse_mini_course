package command

import (
	"encoding/json"
	"fmt"
	"hse_mini_course/accounts/dto"
	"io"
	"net/http"
)

func parseGetAccountResponse(body io.ReadCloser) (dto.GetAccountResponse, error) {
	var response dto.GetAccountResponse
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return response, fmt.Errorf("json decode failed: %w", err)
	}
	return response, nil
}

func printAccount(acc *dto.GetAccountResponse) {
	fmt.Printf("  account name: %s\n", acc.Name)
	fmt.Printf("  balance:      %d\n", acc.Balance)
}

func makeRequest(method, url, contentType string,
	payload io.Reader) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)

	return client.Do(request)
}
