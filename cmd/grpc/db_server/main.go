package main

import (
	"context"
	"fmt"
	"hse_mini_course/proto"
	"hse_mini_course/sqlc"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

const (
	DEFAULT_PORT = "6969"
)

func parseConnectionString() string {
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	if db_host == "" || db_name == "" || db_port == "" || db_user == "" {
		log.Println("DB variables not set up correctly:")
		log.Printf("  db_host: \"%s\"\n", db_host)
		log.Printf("  db_name: \"%s\"\n", db_name)
		log.Printf("  db_port: \"%s\"\n", db_port)
		log.Printf("  db_user: \"%s\"\n", db_user)
		log.Fatalf("Need to specify all four\n")
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		db_user,
		db_password,
		db_host,
		db_port,
		db_name,
	)
}

func main() {
	ctx := context.Background()

	// DB init
	connection_string := parseConnectionString()
	db_connection, err := pgx.Connect(ctx, connection_string)
	if err != nil {
		log.Fatalf("ERR: Unable to connect to the database: %v\n", err)
	}
	defer db_connection.Close(ctx)
	server := newServer(sqlc.New(db_connection))
	log.Printf("Successfully connected to db on %s:%s\n", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	// Server start listening
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Printf("Environment variable empty, using default port %s\n", DEFAULT_PORT)
		port = DEFAULT_PORT
	}
	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("ERR: could not listen: %v\n", err)
	}
	serv := grpc.NewServer()
	proto.RegisterHw3Server(serv, server)
	log.Printf("Serving on address %s\n", addr)
	if err = serv.Serve(listener); err != nil {
		log.Fatalf("ERR: could not serve: %v\n", err)
	}
}
