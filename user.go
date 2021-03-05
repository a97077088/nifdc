package nifdc

import (
	jsoniter "github.com/json-iterator/go"
	"sync"
	"test.com/a/app"
)

type NifdcUser struct {
	*app.AppUser
	mp map[string]interface{}
	mp_lk sync.Mutex
	evmp map[string]interface{}
	evmp_lk sync.Mutex
}


func (u *NifdcUser)V(k string)jsoniter.Any{
	u.mp_lk.Lock()
	defer u.mp_lk.Unlock()
	return jsoniter.Wrap(u.mp)
}
func (u *NifdcUser)SV(k string,v interface{}){
	u.mp_lk.Lock()
	defer u.mp_lk.Unlock()
	u.mp[k]=v
}
func (u *NifdcUser)EV(k string)jsoniter.Any{
	u.evmp_lk.Lock()
	defer u.evmp_lk.Unlock()
	return jsoniter.Wrap(u.evmp)
}
func (u *NifdcUser)EInterface(k string)interface{}{
	u.evmp_lk.Lock()
	defer u.evmp_lk.Unlock()
	return u.evmp
}
func (u *NifdcUser)SEV(k string,v interface{}){
	u.evmp_lk.Lock()
	defer u.evmp_lk.Unlock()
	u.evmp[k]=v
}
func NifdcUserFromInterface(v interface{})*NifdcUser{
	return v.(*NifdcUser)
}