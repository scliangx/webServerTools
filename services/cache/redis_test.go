package cache

import "testing"

func RedisTest(t *testing.T){
	client := NewRedisClient()

	val := client.Get("test")
	t.Logf("get redis value: %s",val)
}
