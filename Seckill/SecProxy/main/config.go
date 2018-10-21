package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golearn/Seckill/SecProxy/service"
)

var(
	secKillConf = &service.SecKillConf{
		SecProductInfoConfMap:make(map[int]*service.SecProductInfoConf,1024),
	}
)

func initConfig()(err error){
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")

	logs.Debug("read config succ, redis addr:%v", redisAddr)
	logs.Debug("read config succ, etcd addr:%v", etcdAddr)

	secKillConf.RedisConf.RedisAddr = redisAddr
	secKillConf.EtcdConf.EtcdAddr = etcdAddr

	if len(redisAddr) == 0 || len(etcdAddr) == 0{
		err = fmt.Errorf("init config failed.redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return
	}

	redisMaxIdle,err := beego.AppConfig.Int("redis_max_idle")
	if err !=nil{
		err = fmt.Errorf("init config failed,read redis_max_idle err:%v",err)
		return
	}

	redisMaxActive,err := beego.AppConfig.Int("redis_max_active")
	if err !=nil{
		err = fmt.Errorf("init config failed,read redis_max_active err:%v",err)
		return
	}

	redisIdleTimeout,err := beego.AppConfig.Int("redis_max_timeout")
	if err !=nil{
		err = fmt.Errorf("init config failed,read redis_max_timeout err:%v",err)
		return
	}
	secKillConf.RedisConf.RedisMaxIdle=redisMaxIdle
	secKillConf.RedisConf.RedisMaxActive=redisMaxActive
	secKillConf.RedisConf.RedisIdleTimeout=redisIdleTimeout

	etcdTimeout,err := beego.AppConfig.Int("etcd_timeout")
	if err !=nil{
		err = fmt.Errorf("connect etcd failed,err:",err)
		return
	}
	secKillConf.EtcdConf.Timeout = etcdTimeout
	secKillConf.EtcdConf.EtcdSecKeyPrefix= beego.AppConfig.String("etcd_sec_key_prefix")
	if len(secKillConf.EtcdConf.EtcdSecKeyPrefix) == 0{
		err = fmt.Errorf("init config failed ,read etcd_sec_key error:%v",err)
		return
	}

	productKey := beego.AppConfig.String("etcd_product_key")
	if len(productKey) == 0{
		err = fmt.Errorf("init config failed ,read etcd_product_key error:%v",err)
		return
	}
	secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s/%s",secKillConf.EtcdConf.EtcdSecKeyPrefix,productKey)

	secKillConf.LogPath = beego.AppConfig.String("log_path")
	secKillConf.LogLevel = beego.AppConfig.String("log_level")





	return
}