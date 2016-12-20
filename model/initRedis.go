package model

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"github.com/myafeier/meeting/config"
)

var res *redis.Pool


func dialFunc()(redis.Conn,error){
	conn,err:=redis.Dial("tcp",config.RedisConfig.Ip+":"+config.RedisConfig.Port)
	if err!=nil{
		log.Panic(err)
	}
	return conn,nil
}
func InitRedis(){
	res=redis.NewPool(dialFunc,300)
}


