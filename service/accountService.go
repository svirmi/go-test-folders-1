package service

import (
	"folders-one/dto"
	"folders-one/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, errs.AppError)
}
