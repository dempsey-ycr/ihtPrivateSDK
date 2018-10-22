package models

import (
	"ihtPrivateSDK/iht/ipfs/config"
	"ihtPrivateSDK/share/logging"
	"ihtPrivateSDK/share/models"
)

var (
	FStore *config.FileStore
	TTL    *config.CacheTTL
	FCat   *config.FileCatalog
)

func init() {

	cfg := config.Default(APP_PID)

	//初始化 MySQL 配置
	err := models.Init(cfg.Db.DriverName, cfg.Db.DataSource)
	if err != nil {
		logging.Fatal(err)
		return
	}

	//初始化 米领 MySQL 配置
	err = models.InitMicrolink(cfg.DbMicroLink.DriverName, cfg.DbMicroLink.DataSource)
	if err != nil {
		logging.Fatal(err)
		return
	}

	//初始化 首证 MySQL 配置
	err = models.InitSZ(cfg.DbSZ.DriverName, cfg.DbSZ.DataSource)
	if err != nil {
		logging.Fatal(err)
		return
	}

	FStore = &cfg.File
	TTL = &cfg.TTL
	FCat = &cfg.Catalog

	// 初始化 Redis 配置
	InitRedisFrame(&cfg.RedisCache, &cfg.Redis, &cfg.RedisMicroLink)
}
