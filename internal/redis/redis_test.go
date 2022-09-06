package cache

import "testing"

func RedisTest(t *testing.T){
	client := GetClient()

	val := client.Get("test")
	t.Logf("get redis value: %s",val)
}
