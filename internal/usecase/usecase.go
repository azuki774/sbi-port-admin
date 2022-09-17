package usecase

import (
	"go.uber.org/zap"
)

type DBRepository interface {
}

type Usecase struct {
	Logger *zap.Logger
	DBRepo DBRepository
}
