package factory

import (
	"net"
	"time"

	"azuki774/sbiport-server/internal/repository"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DBConnectRetry = 5
const DBConnectRetryInterval = 10

type DBInfo struct {
	Host     string
	Port     string
	DBName   string
	UserName string
	UserPass string
}

type HTTPClientInfo struct {
	Scheme string // http or https
	Host   string
	Port   string
}

type RegistInfo struct {
	TargetDir string
}

func setDefaultValue(opts *DBInfo) {
	if opts.Host == "" {
		opts.Host = "localhost"
	}

	if opts.Port == "" {
		opts.Port = "3306"
	}

	if opts.DBName == "" {
		opts.DBName = "sbiport"
	}

	if opts.UserName == "" {
		opts.UserName = "root"
	}

	if opts.UserPass == "" {
		opts.UserPass = "password"
	}
}

func setDefaultValueHTTPClient(opts *HTTPClientInfo) {
	if opts.Scheme == "" {
		opts.Scheme = "http"
	}
	if opts.Host == "" {
		opts.Host = "localhost"
	}
	if opts.Port == "" {
		opts.Port = "80"
	}
}

func NewDBRepo(opts *DBInfo) (*repository.DBRepository, error) {
	l, err := NewLogger()
	if err != nil {
		return nil, err
	}

	setDefaultValue(opts)
	addr := net.JoinHostPort(opts.Host, opts.Port)
	dsn := opts.UserName + ":" + opts.UserPass + "@(" + addr + ")/" + opts.DBName + "?parseTime=true&loc=Local"
	var gormdb *gorm.DB
	for i := 0; i < DBConnectRetry; i++ {
		gormdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			// Success DB connect
			l.Info("DB connect")
			break
		}
		l.Warn("DB connection retry")

		if i == DBConnectRetry {
			l.Error("failed to connect (DB)", zap.Error(err))
			return nil, err
		}

		time.Sleep(DBConnectRetryInterval * time.Second)
	}

	return &repository.DBRepository{Conn: gormdb}, nil
}

func NewHTTPClient(opts *HTTPClientInfo) *repository.Client {
	setDefaultValueHTTPClient(opts)
	return &repository.Client{Scheme: opts.Scheme, Host: opts.Host, Port: opts.Port}
}
