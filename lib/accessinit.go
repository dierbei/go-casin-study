package lib

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
)

var E *casbin.Enforcer

func init() {
	initDB()
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal()
	}
	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	if err != nil {
		log.Fatal()
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal()
	}
	E = e
	initPolicy()
}

//从我们的库里初始化 策略数据
func initPolicy() {
	// add p
	E.AddPolicy("member", "/depts", "GET")
	E.AddPolicy("admin", "/depts", "POST")
	// add g
	E.AddRoleForUser("zhangsan", "member")
}
