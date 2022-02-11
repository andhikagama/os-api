package resource

import (
	"fmt"
	"time"

	"github.com/andhikagama/os-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB .
func NewDB(config *config.EnvConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%v:%v@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True`,
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLHost,
		config.MySQLPort,
		config.MySQLDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if config.MySQLDebug {
		return db.Debug(), nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.MySQLMaxIdleConn)
	sqlDB.SetMaxOpenConns(config.MySQLMaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MySQLMaxLifetimeConn) * time.Second)

	return db, nil
}
