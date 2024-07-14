package command

import (
	"context"
	"fmt"
	"hse_mini_course/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func printAccount(acc *proto.GetAccountResponse) {
	fmt.Printf("  account name: %s\n", acc.Name)
	fmt.Printf("  balance:      %d\n", acc.Balance)
}

func handleError(err error, nameNotFound *string, nameTaken *string, nameInvalid *string) error {
	if err == nil {
		panic("only call when err is not nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	switch st.Code() {
	case codes.NotFound:
		if nameNotFound == nil {
			return fmt.Errorf("Not Found response not expected: %v", st.Message())
		}
		fmt.Printf("name \"%s\" not found\n", *nameNotFound)

	case codes.InvalidArgument:
		if nameInvalid == nil {
			return fmt.Errorf("Name Invalid response not expected: %v", st.Message())
		}
		fmt.Printf("name \"%s\" is not valid: %v\n", *nameInvalid, st.Message())

	case codes.AlreadyExists:
		if nameTaken == nil {
			return fmt.Errorf("Name Taken response not expected: %v", st.Message())
		}
		fmt.Printf("name \"%s\" is already taken\n", *nameTaken)

	case codes.Internal:
		fmt.Println("internal server error")

	default:
		return fmt.Errorf("unexpected error from server: code: %d, message: %s",
			st.Code(), st.Message())
	}

	return nil
}

func (cmd *Command) get(ctx context.Context, client proto.Hw3Client) error {
	resp, err := client.GetAccount(ctx, &proto.GetAccountRequest{
		Name: cmd.Name,
	})
	if err != nil {
		return handleError(err, &cmd.Name, nil, nil)
	}

	fmt.Println("successfully retrieved:")
	printAccount(resp)

	return nil
}

func (cmd *Command) create(ctx context.Context, client proto.Hw3Client) error {
	resp, err := client.CreateAccount(ctx, &proto.CreateAccountRequest{
		Name: cmd.Name,
	})
	if err != nil {
		return handleError(err, nil, &cmd.Name, &cmd.Name)
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
		return handleError(err, &cmd.Name, nil, nil)
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
		return handleError(err, &cmd.Name, &cmd.NewName, &cmd.NewName)
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
		return handleError(err, &cmd.Name, nil, nil)
	}

	fmt.Println("successfully deleted")
	return nil
}
