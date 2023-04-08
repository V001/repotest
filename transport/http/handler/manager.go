package handler

import (
	"repotest/config"
	"repotest/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(conf *config.Config, srv *service.Manager) *Manager {
	return &Manager{
		srv: srv,
	}
}
