package session

import (
	"fmt"
	"sync"
)

type Manager struct {
	cookieName string  //private cookiename
	lock sync.Mutex //protects session
	provider Provider
	maxlifetime int64
}
type Provider interface {
	SessionInit(sid string) (Session, error) //session初始化
	SessionRead(sid string) (Session, error) //返回sid代表的session，若不存在就调用SessionInit创建一个新的并返回
	SessionDestroy(sid string) error //销毁session
	SessionGC(maxLifeTime int64) //根据过期时间删除过期数据
}
type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}
var provides  = make(map[string]Provider)
//创建全局session管理器
func NewManager(provideName,cookieName string,maxlifetime int64) (*Manager, error)  {
	provider,ok := provides[provideName] //从provides这个map中查找键为provideName的信息
	if !ok {
		return nil,fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

