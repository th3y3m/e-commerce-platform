package Provider

import (
	"th3y3m/e-commerce-platform/Util"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IDb interface {
	GetDB() (*gorm.DB, error)
	GetRedis() (*redis.Client, error)
	// GetMockDb() (*mock.Mock, error)
}

type DbProvider struct {
	log *logrus.Logger
}

func NewDbProvider(log *logrus.Logger) IDb {
	return &DbProvider{log: log}
}

var _db *gorm.DB

func (d *DbProvider) GetDB() (*gorm.DB, error) {
	if _db != nil {
		return _db, nil
	}

	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}
	_db = db
	return db, nil
}

var _redis *redis.Client

func (d *DbProvider) GetRedis() (*redis.Client, error) {
	if _redis != nil {
		return _redis, nil
	}

	redis, err := Util.ConnectToRedis()
	if err != nil {
		return nil, err
	}
	_redis = redis
	return redis, nil
}

// func (d *DbProvider) GetMockDb() (*mock.Mock, error) {
// 	if _mock == nil {
// 		_mock = new(mock.Mock)
// 	}
// 	return _mock, nil
// }
