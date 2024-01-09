package uc

import (
	"chatmerger/internal/domain"
	"chatmerger/internal/domain/model"
	"chatmerger/internal/usecase"
	"errors"
	"fmt"
)

var _ usecase.CreateClientSessionUc = (*CreateClientSession)(nil)

type CreateClientSession struct {
	sessionRepo domain.ClientSessionsRepository
	clientRepo  domain.ClientsRepository
}

func NewCreateClientSession(clientRepo domain.ClientsRepository, sessionRepo domain.ClientSessionsRepository) *CreateClientSession {
	return &CreateClientSession{sessionRepo: sessionRepo, clientRepo: clientRepo}
}

var (
	ErrorClientWithGivenApiKeyNotFound = errors.New("client with given ApiKey not found")
	ErrorClientSessionsAlreadyExists   = errors.New("client session already exists")
)

func (c *CreateClientSession) CreateClientSession(input model.CreateClientSession) (*model.ClientSession, error) {
	clients, err := c.clientRepo.GetClients()
	if err != nil {
		return nil, fmt.Errorf("getting client list: %s", err)
	}
	// apikey exists?
	var expectedClient *model.Client
	for _, client := range clients {
		if client.ApiKey == input.ApiKey {
			expectedClient = &client
			break
		}
	}
	// return error if not exist
	if expectedClient == nil {
		return nil, ErrorClientWithGivenApiKeyNotFound
	}

	sessions, err := c.sessionRepo.Connected()
	if err != nil {
		return nil, fmt.Errorf("call sessions connected list getter failed: %s", err)
	}
	// session don't already connected
	for _, session := range sessions {
		if session.Id == expectedClient.Id {
			return nil, ErrorClientSessionsAlreadyExists
		}
	}
	// add client to list
	connect, err := c.sessionRepo.Connect(*expectedClient)
	if err != nil {
		return nil, fmt.Errorf("connect client failed: %s", err)
	}
	return connect, nil
}
