package command

import "fmt"

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Delta   int32
}

func (cmd *Command) Execute() error {
	switch cmd.Cmd {
	case "create":
		if err := cmd.create(); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := cmd.get(); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "transact":
		if err := cmd.newTransaction(); err != nil {
			return fmt.Errorf("new transaction failed: %w", err)
		}

		return nil

	case "change":
		if err := cmd.changeName(); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}

		return nil
	case "delete":
		if err := cmd.delete(); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s\n  available commands: create, get, transact, change, delete", cmd.Cmd)
	}
}
