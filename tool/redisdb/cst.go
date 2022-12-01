package redisdb

import (
	"fmt"
	"hat_logic/tool/consul"
)

type ZSetScoreTyp = float64

//GetRedisKey 当前集群当前服务
func GetRedisKey(key string, sid int32) string {
	serviceInfo := consul.WhoAmI()
	return fmt.Sprintf("%s:%s:%s:%d", serviceInfo.Cluster, serviceInfo.Service, key, sid)
}
