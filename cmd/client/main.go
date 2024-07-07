package main

import (
	"flag"
	"fmt"
	"hse_mini_course/cmd/client/command"
)

func main() {
	portVal := flag.Int("port", 6969, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "name to change to")
	deltaVal := flag.Int("delta", 0, "change in balance")

	flag.Parse()

	cmd := command.Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Delta:   *deltaVal,
	}

	if err := cmd.Execute(); err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
}
