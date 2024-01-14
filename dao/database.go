package dao

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Config 数据库配置
type Config struct {
	// DSN "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
	DSN string `toml:"dsn"`

	// 设置空闲连接池中连接的最大数量
	MaxIdleCount int `toml:"max_idle_count"`

	// 设置打开数据库连接的最大数量
	MaxOpen int `toml:"max_open"`

	// 设置了连接可复用的最大时间
	MaxLifetime time.Duration `toml:"max_lifetime"`
}

// InitDatabase 数据库初始化
func InitDatabase(cfg *Config) (err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.DSN, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.Ping(); err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleCount)
	sqlDB.SetMaxOpenConns(cfg.MaxOpen)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	// 仅支持建表，不支持修改字段和删除字段，避免意外导致丢失数据。
	if err = db.AutoMigrate(new(User)); err != nil {
		return err
	}

	if err = InitUserDB(); err != nil {
		return err
	}

	return nil
}
