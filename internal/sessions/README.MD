#### session 实现流程


- 该结构体主要是管理全局的session,因为session可能有多个，基于内存的，基于redis的等等，通过cookieName可以在改管理器中获取到对应的session对象
```go

type Manager struct {
	cookieName  string
	lock        sync.Mutex //互斥锁
	storage SessionStorage   //存储session方式
	maxAge int64      //有效期
}
```

- 该结构体主要是暴露出来session存储方式，可以自己实现
```go
// SessionStorage session存储方式接口
type SessionStorage interface {
    // InitSessionObj 初始化一个session，sid根据需要生成后传入
    InitSessionObj(sid string) (Session, error)
    // GetSessionObj 根据sid,获取session
    GetSessionObj(sid string) (Session, error)
    // DestroySession 销毁session
    DestroySession(sid string) error
    //GC 回收
    GC(maxLifeTime int64)
}
```

- 该结构体主要是session的实际增删改查操作
```go
//Session 操作接口
type Session interface {
    Set(key, value interface{}) error
    Get(key interface{}) interface{}
    Delete(ket interface{}) error
    SessionID() string
}
```

- case
```go
func TestMemory(t *testing.T) {
	// 基于多重存储方式实现session，并且有一个唯一名称
	SessionInitByMemory("storeageName")
	SessionInitByMemory("storeageName1")
	// 实例化session管理器
	// 会根据storeageName拿去session，初始化为cookieName的一个管理器对象
	manager,_ := NewSessionManager("storeageName","cookieName",600)
	manager1,_ := NewSessionManager("storeageName1","cookieName1",600)
	// 根据storeageName 去管理器中获取对应的session对象，如果没有的话会进行初始化
	s, _ := manager.storage.GetSessionObj("storeageName")
	s1,_ := manager1.storage.GetSessionObj("storeageName1")
	// 操作
	_ = s.Set("web", "web")
	_ = s1.Set("web1", "web1")
	fmt.Println(s.Get("web").(string))
	fmt.Println(s1.Get("web1").(string))
	s1.Delete("web1")
	fmt.Println(s1.Get("web1"))
}
```

