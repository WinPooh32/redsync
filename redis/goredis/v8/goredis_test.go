package goredis

import "github.com/winpooh32/redsync/v4/redis"

var _ redis.Conn = (*conn)(nil)

var _ redis.Pool = (*pool)(nil)
