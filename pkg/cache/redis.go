package cache

import (
	"github.com/go-redis/redis/v8"
	"strings"
)

type cache struct {
	client redis.Cmdable
}

func SetUpRedis(Db int, Cluster int) {

}

func SetUpRedisCluster(nodes []string) error {
	//redisMaster = redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: nodes,
	//})
	//redisSlave = redisMaster
	//if redisMaster == nil || redisSlave == nil {
	//	return fmt.Errorf("redisCluster: nil")
	//}
	return nil
}

func formatRedisClusterAddress(list string) []string {
	var returnList []string

	hosts := strings.Split(list, ",")

	for _, v := range hosts {
		returnList = append(returnList, v)
	}
	return returnList
}
