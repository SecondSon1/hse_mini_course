package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hse_mini_course/accounts/dto"
	"io"
	"net/http"
	"net/url"
)

func (cmd *Command) get() error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account/%s", cmd.Host, cmd.Port, url.QueryEscape(cmd.Name)),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error: %s", string(body))
	}

	response, err := parseGetAccountResponse(resp.Body)
	if err != nil {
    return fmt.Errorf("parse account response failed: %w", err)
	}
	fmt.Println("response:")
	printAccount(&response)

	return nil
}

func (cmd *Command) create() error {
	request := dto.CreateAccountRequest{
		Name: cmd.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/new", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

    return fmt.Errorf("resp error: %s", string(body))
	}

	response, err := parseGetAccountResponse(resp.Body)
	if err != nil {
    return fmt.Errorf("parse account response failed: %w", err)
	}
	fmt.Println("created acount:")
	printAccount(&response)
	return nil
}

func (cmd *Command) newTransaction() error {
	request := dto.NewTransactionRequest{
    Delta: cmd.Delta,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/%s", cmd.Host, cmd.Port, url.QueryEscape(cmd.Name)),
    "application/json",
    bytes.NewReader(data),
	)
  if err != nil {
    return fmt.Errorf("http post failed: %w", err)
  }

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

    return fmt.Errorf("resp error: %s", string(body))
	}

  response, err := parseGetAccountResponse(resp.Body)
	if err != nil {
    return fmt.Errorf("parse account response failed: %w", err)
	}
	fmt.Println("success")
  fmt.Println("new account status:")
  printAccount(&response)
	return nil
}

func (cmd *Command) changeName() error {
	request := dto.ChangeNameRequest{
    NewName: cmd.NewName,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := makeRequest(
    http.MethodPatch,
		fmt.Sprintf("http://%s:%d/account/%s", cmd.Host, cmd.Port, url.QueryEscape(cmd.Name)),
    "application/json",
    bytes.NewReader(data),
	)
  if err != nil {
    return fmt.Errorf("http patch failed: %w", err)
  }

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

    return fmt.Errorf("resp error: %s", string(body))
	}

  response, err := parseGetAccountResponse(resp.Body)
	if err != nil {
    return fmt.Errorf("parse account response failed: %w", err)
	}
	fmt.Println("success")
  fmt.Println("new account status:")
  printAccount(&response)
	return nil
}

func (cmd *Command) delete() error {
	resp, err := makeRequest(
    http.MethodDelete,
		fmt.Sprintf("http://%s:%d/account/%s", cmd.Host, cmd.Port, url.QueryEscape(cmd.Name)),
    "application/json",
    nil,
	)
  if err != nil {
    return fmt.Errorf("http delete failed: %w", err)
  }

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

    return fmt.Errorf("resp error: %s", string(body))
	}

	fmt.Println("successfully deleted")
	return nil
}
