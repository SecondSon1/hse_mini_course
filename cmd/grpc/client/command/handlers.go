package command

import (
	"context"
	"fmt"
	"hse_mini_course/proto"
)

func printAccount(acc *proto.GetAccountResponse) {
	fmt.Printf("  account name: %s\n", acc.Name)
	fmt.Printf("  balance:      %d\n", acc.Balance)
}

func (cmd *Command) get(ctx context.Context, client proto.Hw3Client) error {
	resp, err := client.GetAccount(ctx, &proto.GetAccountRequest{
		Name: cmd.Name,
	})
	if err != nil {
		return err
	}

	fmt.Println("response:")
	printAccount(resp)

	return nil
}

func (cmd *Command) create(ctx context.Context, client proto.Hw3Client) error {
	resp, err := client.CreateAccount(ctx, &proto.CreateAccountRequest{
		Name: cmd.Name,
	})
	if err != nil {
		return err
	}

	fmt.Println("created acount:")
	printAccount(resp)

	return nil
}

func (cmd *Command) newTransaction(ctx context.Context, client proto.Hw3Client) error {
	resp, err := client.NewTransaction(ctx, &proto.NewTransactionRequest{
		Name:  cmd.Name,
		Delta: cmd.Delta,
	})
	if err != nil {
		return err
	}

	fmt.Println("success")
	fmt.Println("new account status:")
	printAccount(resp)

	return nil
}

func (cmd *Command) changeName(ctx context.Context, client proto.Hw3Client) error {
	resp, err := client.ChangeName(ctx, &proto.ChangeNameRequest{
		Name:    cmd.Name,
		NewName: cmd.NewName,
	})
	if err != nil {
		return err
	}

	fmt.Println("success")
	fmt.Println("new account status:")
	printAccount(resp)
	return nil
}

func (cmd *Command) delete(ctx context.Context, client proto.Hw3Client) error {
	_, err := client.DeleteAccount(ctx, &proto.DeleteAccountRequest{
		Name: cmd.Name,
	})
	if err != nil {
		return err
	}

	fmt.Println("successfully deleted")
	return nil
}
