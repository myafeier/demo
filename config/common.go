package config

import (
	"os"
	"log"
)

var SiteCommonConfig *siteCommonConfig
var RedisConfig *redisConfig
type  siteCommonConfig  struct {
	AppPath string
	SiteUrl string
	ListenPort string
	Title string
	TemplatePath string
}

type redisConfig struct {
	Ip string
	Port string
	Passwd string
}


func InitConfig(){
	var err error
	SiteCommonConfig=new(siteCommonConfig)
	SiteCommonConfig.AppPath,err=os.Getwd()
	if err!=nil{
		log.Panic(err)
	}

	SiteCommonConfig.SiteUrl="http://localhost:8080"
	SiteCommonConfig.ListenPort="8080"
	SiteCommonConfig.Title="测试站点"
	SiteCommonConfig.TemplatePath=SiteCommonConfig.AppPath+"/template"

	RedisConfig=new(redisConfig)
	RedisConfig.Ip="127.0.0.1"
	RedisConfig.Port="6379"
}
