package redis

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Persistent *redis.Pool //持久保存(redis会固化)
	Temporary  *redis.Pool //临时缓存(redis仅做缓存)
)

var (
	ErrNil = redis.ErrNil
)

func InitPersistanRedisPool(addr string, pwd string, db int) (err error) {
	pool := newPool(addr, pwd, db)
	err = pool.Get().Err()
	if err != nil {
		return
	}
	Persistent = pool
	return
}

func InitTemporarRedisPool(addr string, pwd string, db int) (err error) {
	pool := newPool(addr, pwd, db)
	err = pool.Get().Err()
	if err != nil {
		return
	}
	Temporary = pool
	return
}

func newPool(addr string, pwd string, db int) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     100,
		MaxActive:   100,
		Wait:        true,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Println("[ERROR] failed to dial redis server,err:", err)
				return nil, err
			}

			if len(pwd) > 0 {
				_, err = conn.Do(Auth, pwd)
				if err != nil {
					log.Println("[ERROR] wrong password ,err:", err)
					conn.Close()
					return nil, err
				}
			}

			_, err = conn.Do(Select, db)
			if err != nil {
				log.Println("[ERROR] failed to select db ,err:", err)
				conn.Close()
				return nil, err
			}
			return conn, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do(Ping)
			return err
		},
	}
}

func String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

func StringMap(result interface{}, err error) (map[string]string, error) {
	return redis.StringMap(result, err)
}

func Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

func Int64s(reply interface{}, err error) ([]int64, error) {
	return redis.Int64s(reply, err)
}

func Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

func Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}
