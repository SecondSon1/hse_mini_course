package command

import (
	"context"
	"fmt"
	"hse_mini_course/proto"
)

type Command struct {
	Cmd     string
	Name    string
	NewName string
	Delta   int32
}

func (cmd *Command) Execute(ctx context.Context, client proto.Hw3Client) error {
	switch cmd.Cmd {
	case "create":
		if err := cmd.create(ctx, client); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := cmd.get(ctx, client); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "transact":
		if err := cmd.newTransaction(ctx, client); err != nil {
			return fmt.Errorf("new transaction failed: %w", err)
		}

		return nil

	case "change":
		if err := cmd.changeName(ctx, client); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}

		return nil
	case "delete":
		if err := cmd.delete(ctx, client); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s\n  available commands: create, get, transact, change, delete", cmd.Cmd)
	}
}
