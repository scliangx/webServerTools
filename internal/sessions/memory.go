package sessions

import (
	"container/list"
	"sync"
	"time"
)

var storage = &FromMemory{list: list.New()}

// FromMemory session来自内存 实现
type FromMemory struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}

// SessionStoreByMemory session实现
type SessionStoreByMemory struct {
	sid      string                      //session id 唯一标示
	LastTime time.Time                   //最后访问时间
	value    map[interface{}]interface{} //session 里面存储的值
}

func SessionInitByMemory(sessionName string) {
	storage.sessions = make(map[string]*list.Element, 0)
	//注册  memory 调用的时候一定有一致
	Register(sessionName, storage)
}

/**
	实现session对象
*/

// Set 设置
func (st *SessionStoreByMemory) Set(key, value interface{}) error {
	st.value[key] = value
	return storage.update(st.sid)
}

// Get 获取session
func (st *SessionStoreByMemory) Get(key interface{}) interface{} {
	_ = storage.update(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

// Delete 删除
func (st *SessionStoreByMemory) Delete(key interface{}) error {
	delete(st.value, key)
	return storage.update(st.sid)
}

// SessionID 获取sessionId
func (st *SessionStoreByMemory) SessionID() string {
	return st.sid
}


func (memory *FromMemory) InitSessionObj(sid string) (Session, error) {
	memory.lock.Lock()
	defer memory.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newS := &SessionStoreByMemory{sid: sid, LastTime: time.Now(), value: v}
	element := memory.list.PushBack(newS)
	memory.sessions[sid] = element
	return newS, nil
}

func (memory *FromMemory) GetSessionObj(sid string) (Session, error) {
	if element, ok := memory.sessions[sid]; ok {
		return element.Value.(*SessionStoreByMemory), nil
	} else {
		sess, err := memory.InitSessionObj(sid)
		return sess, err
	}
}

func (memory *FromMemory) DestroySession(sid string) error {
	if element, ok := memory.sessions[sid]; ok {
		delete(memory.sessions, sid)
		memory.list.Remove(element)
		return nil
	}
	return nil
}

func (memory *FromMemory) GC(maxLifeTime int64) {
	memory.lock.Lock()
	defer memory.lock.Unlock()
	for {
		element := memory.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStoreByMemory).LastTime.Unix() + maxLifeTime) <
			time.Now().Unix() {
			memory.list.Remove(element)
			delete(memory.sessions, element.Value.(*SessionStoreByMemory).sid)
		} else {
			break
		}
	}
}

func (memory *FromMemory) update(sid string) error {
	memory.lock.Lock()
	defer memory.lock.Unlock()
	if element, ok := memory.sessions[sid]; ok {
		element.Value.(*SessionStoreByMemory).LastTime = time.Now()
		memory.list.MoveToFront(element)
		return nil
	}
	return nil
}
