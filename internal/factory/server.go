package factory

import (
	"azuki774/sbiport-server/internal/server"
	"azuki774/sbiport-server/internal/usecase"
)

type ServerRunOption struct{
	Host string
	Port string
}

func NewServer(opts *ServerRunOption, u *usecase.Usecase) (*server.Server, error) {
	l, err := NewLogger()
	if err != nil {
		return nil, err
	}
	if opts.Port == "" {
		opts.Port = "80" // default value
	}

	return &server.Server{Logger: l, Host: "", Port: opts.Port, Usecase: u}, nil
}
