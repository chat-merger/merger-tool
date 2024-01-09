package uc

import (
	"chatmerger/internal/domain"
	"chatmerger/internal/domain/model"
	"chatmerger/internal/usecase"
	"fmt"
	"slices"
)

var _ usecase.DeleteClientUc = (*DeleteClient)(nil)

type DeleteClient struct {
	clientRepos domain.ClientsRepository
}

func NewDeleteClient(clientRepos domain.ClientsRepository) *DeleteClient {
	return &DeleteClient{clientRepos: clientRepos}
}

func (r *DeleteClient) DeleteClients(ids []model.ID) error {
	clients, err := r.clientRepos.GetClients()
	if err != nil {
		return fmt.Errorf("getting client list: %s", err)
	}
	// remove elements..
	var newClientsList = slices.DeleteFunc(clients, func(client model.Client) bool {
		// witch in the list
		return slices.ContainsFunc(ids, func(id model.ID) bool {
			return id == client.Id
		})
	})

	err = r.clientRepos.SetClients(newClientsList)
	if err != nil {
		return fmt.Errorf("setting clients list: %s", err)
	}

	return nil
}
