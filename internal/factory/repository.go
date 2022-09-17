package factory

import (
	"net"
	"time"

	"azuki774/sbi-port-admin/internal/repository"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DBConnectRetry = 5
const DBConnectRetryInterval = 10

func NewDBRepo(user string, password string, host string, port string, dbName string) (*repository.DBRepository, error) {
	l, err := NewLogger()
	if err != nil {
		return nil, err
	}

	addr := net.JoinHostPort(host, port)
	dsn := user + ":" + password + "@(" + addr + ")/" + dbName + "?parseTime=true&loc=Local"
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
