package model

import "time"

type UserCreateReq struct {
	Name    string
	SurName string
	Pass    string
	Login   string
}

type CreateResp struct {
	ID        uint
	CreatedAt time.Time
}
