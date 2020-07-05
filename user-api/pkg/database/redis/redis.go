package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

const (
	redisExpire = 15
)

var Conn = Connect()

func Connect() redis.Conn {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Print(err.Error())
	}
	return c
}

func Set(key string, value []byte) error {
	_, err := Conn.Do("SET", key, value)
	if err != nil {
		log.Print(err.Error())
	}

	_, err = Conn.Do("EXPIRE", key, redisExpire)
	if err != nil {
		log.Print(err.Error())
	}
		return err
}

func Get(key string) ([]byte, error) {
	var data []byte
	data, err := redis.Bytes(Conn.Do("GET", key))
	if err != nil {
		log.Print(err.Error())
	}

	return data, err
}

func Flush(key string) (bool, error) {
	response ,err := redis.Bool(Conn.Do("DEL", key))
	if err != nil {
		log.Print(err.Error())
	}

	return response, err
}