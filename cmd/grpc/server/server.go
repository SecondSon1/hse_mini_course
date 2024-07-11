package main

import (
	"context"
	"fmt"
	"hse_mini_course/accounts/models"
	"hse_mini_course/proto"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	PORT = 6969

	IDS_FROM    = 100000
	IDS_TO      = 999999
	GUARD_LIMIT = 20
)

type server struct {
	proto.UnimplementedHw3Server
	accounts      map[string]*models.Account
	ids           map[uint32]bool
	accountsGuard *sync.RWMutex
	idsGuard      *sync.RWMutex
}

func newServer() *server {
	return &server{
		accounts:      make(map[string]*models.Account),
		ids:           make(map[uint32]bool),
		accountsGuard: &sync.RWMutex{},
		idsGuard:      &sync.RWMutex{},
	}
}

func (s *server) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.GetAccountResponse, error) {
	name := request.Name
	err := validateName(name)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	s.accountsGuard.Lock()

	_, taken := s.accounts[name]
	if taken {
		s.accountsGuard.Unlock()
		return nil, status.Errorf(
			codes.AlreadyExists,
			"name \"%s\" is already taken",
			name,
		)
	}

	new_acc := models.Account{
		Id:      generateId(),
		Name:    name,
		Balance: 0,
	}

	var iter_limit int
	s.idsGuard.Lock()

	for iter_limit = 0; iter_limit < GUARD_LIMIT && s.ids[new_acc.Id]; iter_limit++ {
		new_acc.Id = generateId()
	}
	if iter_limit == GUARD_LIMIT {
		s.idsGuard.Unlock()
		s.accountsGuard.Unlock()
		log.Printf("[WARNING]: Hit iteration limit while generating ids\n")
		return nil, status.Error(codes.Internal, "internal error")
	}
	s.ids[new_acc.Id] = true
	s.accounts[name] = &new_acc

	s.idsGuard.Unlock()
	s.accountsGuard.Unlock()

	return modelToResponse(&new_acc), nil
}

func (s *server) GetAccount(ctx context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	name := request.Name
	s.accountsGuard.RLock()

	account, found := s.accounts[name]
	if !found {
		s.accountsGuard.RUnlock()
		return nil, status.Errorf(
			codes.NotFound,
			"name \"%s\" not found",
			name,
		)
	}
	s.accountsGuard.RUnlock()
	return modelToResponse(account), nil
}

func (s *server) NewTransaction(ctx context.Context, request *proto.NewTransactionRequest) (*proto.GetAccountResponse, error) {
	name := request.Name
	delta := request.Delta

	s.accountsGuard.Lock()

	acc, found := s.accounts[name]
	if !found {
		s.accountsGuard.Unlock()
		return nil, status.Errorf(
			codes.NotFound,
			"name \"%s\" not found",
			name,
		)
	}
	acc.Balance += delta
	s.accountsGuard.Unlock()
	return modelToResponse(acc), nil
}

func (s *server) ChangeName(ctx context.Context, request *proto.ChangeNameRequest) (*proto.GetAccountResponse, error) {
	name := request.Name
	newName := request.NewName
	err := validateName(newName)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	s.accountsGuard.Lock()
	acc, found := s.accounts[name]
	if !found {
		s.accountsGuard.Unlock()
		return nil, status.Errorf(
			codes.NotFound,
			"name \"%s\" not found",
			name,
		)
	}
	_, found = s.accounts[newName]
	if found {
		s.accountsGuard.Unlock()
		return nil, status.Errorf(
			codes.AlreadyExists,
			"name \"%s\" already taken",
			newName,
		)
	}

	delete(s.accounts, name)
	acc.Name = newName
	s.accounts[newName] = acc
	s.accountsGuard.Unlock()
	return modelToResponse(acc), nil
}

func (s *server) DeleteAccount(ctx context.Context, request *proto.DeleteAccountRequest) (*proto.Empty, error) {
	name := request.Name

	s.accountsGuard.Lock()
	account, found := s.accounts[name]
	if !found {
		s.accountsGuard.Unlock()
		return nil, status.Errorf(
			codes.NotFound,
			"name \"%s\" not found",
			name,
		)
	}
	id := account.Id

	s.idsGuard.Lock()

	delete(s.ids, id)
	delete(s.accounts, name)

	s.idsGuard.Unlock()
	s.accountsGuard.Unlock()

	return &proto.Empty{}, nil
}

func main() {
	addr := fmt.Sprintf(":%d", PORT)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("ERR: could not listen: %v\n", err)
	}
	serv := grpc.NewServer()
	proto.RegisterHw3Server(serv, newServer())
	log.Printf("Serving on address %s\n", addr)
	if err = serv.Serve(listener); err != nil {
		log.Fatalf("ERR: could not serve: %v\n", err)
	}
}
