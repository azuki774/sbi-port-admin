package factory

import (
	"azuki774/sbiport-server/internal/repository"
	"azuki774/sbiport-server/internal/usecase"
)

func NewUsecase(dbRepo *repository.DBRepository) (*usecase.Usecase, error) {
	l, err := NewLogger()
	if err != nil {
		return nil, err
	}

	return &usecase.Usecase{Logger: l, DBRepo: dbRepo}, nil
}

func NewUsecaseClient(HTTPClient *repository.Client) (*usecase.UsecaseClient, error) {
	l, err := NewLogger()
	if err != nil {
		return nil, err
	}

	return &usecase.UsecaseClient{Logger: l, HTTPClient: HTTPClient}, nil
}
