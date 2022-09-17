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
