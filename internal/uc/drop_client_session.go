package uc

import (
	"chatmerger/internal/domain"
	"chatmerger/internal/domain/model"
	"chatmerger/internal/usecase"
)

var _ usecase.DropClientSessionUc = (*DropClientSession)(nil)

type DropClientSession struct {
	sessionsRepo domain.ClientSessionsRepository
}

func NewDropClientSession(sessionsRepo domain.ClientSessionsRepository) *DropClientSession {
	return &DropClientSession{sessionsRepo: sessionsRepo}
}

func (d *DropClientSession) DropClientSession(ids []model.ID) error {
	for _, id := range ids {
		d.sessionsRepo.Disconnect(id) // todo
	}
	return nil
}
