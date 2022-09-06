package sessions

import (
	"fmt"
	"testing"
)

func TestMemory(t *testing.T) {
	// 注册一个
	SessionInitByMemory("web")
	SessionInitByMemory("web1")
	manager,_ := NewSessionManager("web","web",600)
	manager1,_ := NewSessionManager("web1","web1",600)
	s, _ := manager.storage.GetSessionObj("web")
	s1,_ := manager1.storage.GetSessionObj("web1")
	_ = s.Set("web", "web")
	_ = s1.Set("web1", "web1")
	fmt.Println(s.Get("web").(string))
	fmt.Println(s1.Get("web1").(string))
	s1.Delete("web1")
	fmt.Println(s1.Get("web1"))
}
