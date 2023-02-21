package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

/*

make redis cluster
https://github.com/sam-mix/docker-redis-cluster

redis: 2023/02/21 23:51:36 cluster.go:1760: getting command info: DENIED Redis is running in protected mode because protected mode is enabled and no password is set for
the default user. In this mode connections are only accepted from the loopback interface. If you want to connect from external computers to Redis you may adopt one of the
following solutions: 1) Just disable protected mode sending the command 'CONFIG SET protected-mode no' from the loopback interface by connecting to Redis from the same
host the server is running, however MAKE SURE Redis is not publicly accessible from internet if you do so. Use CONFIG REWRITE to make this change permanent. 2)
Alternatively you can just disable the protected mode by editing the Redis configuration file, and setting the protected mode option to 'no', and then restarting the
server. 3) If you started the server manually just for testing, restart it with the '--protected-mode no' option. 4) Setup a an authentication password for the default
user. NOTE: You only need to do one of the above things in order for the server to start accepting connections from the outside.
get hello: redis: nil

*/

func main() {
	ctx := context.Background()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	rdb.Set(ctx, "hello", "world", 0)
	x := rdb.Get(ctx, "hello")
	fmt.Println(x)
}
