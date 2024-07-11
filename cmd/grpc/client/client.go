package main

import (
	"context"
	"flag"
	"fmt"
	"hse_mini_course/cmd/grpc/client/command"
	"hse_mini_course/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	portVal := flag.Int("port", 6969, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "name to change to")
	deltaVal := flag.Int("delta", 0, "change in balance")

	flag.Parse()

	if *nameVal == "" {
		fmt.Println("Name cannot be empty")
		return
	}

	cmd := command.Command{
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Delta:   int32(*deltaVal),
	}

	addr := fmt.Sprintf("%s:%d", *hostVal, *portVal)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("ERR: could not create client: %v", err)
	}
	client := proto.NewHw3Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := cmd.Execute(ctx, client); err != nil {
		fmt.Printf("ERR: could not execute: %s\n", err)
	}
}
