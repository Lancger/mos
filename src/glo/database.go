package glo

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db    *gorm.DB
	RdsDb *redis.Pool
)

// 连接数据库
func DbConnect() {
	db, err := gorm.Open(Config.MosAPI.Database.Dialect, Config.MosAPI.Database.Addr)
	if err != nil {
		log.Panicln("打开数据库失败:", err)
	}
	db.LogMode(true)
	Db = db.Set("gorm:table_options", "charset=utf8")
}

// 断开数据库
func DbDisconnect() {
	Db.Close()
}

// RdsInit func redis pool
func RdsInit() {
	maxIdle := Config.MosAPI.Redis.MaxIdle
	RdsDb = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Config.MosAPI.Redis.Addr)
			if err != nil {
				log.Panicln("打开redis数据库失败:", err)
				return c, fmt.Errorf("redis connection error: %s", err)
			}
			if Config.MosAPI.Redis.Password != `` {
				if _, authErr := c.Do("AUTH", Config.MosAPI.Redis.Password); authErr != nil {
					log.Panicln("验证redis密码数据库失败:", authErr)
					return c, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

// RdsDisConnect DisConnect Redis Pool
func RdsDisConnect() {
	RdsDb.Close()
}
