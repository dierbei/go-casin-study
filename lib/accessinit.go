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
	m := make([]*RoleRel, 0)
	GetRoles(0, &m, "") //获取角色 对应
	for _, r := range m {
		_, err := E.AddRoleForUser(r.PRole, r.Role)
		if err != nil {
			log.Fatal(err)
		}
	}
	/////// 初始化用户角色
	userRoles := GetUserRoles()
	for _, ur := range userRoles {
		_, err := E.AddRoleForUser(ur.UserName, ur.RoleName)
		if err != nil {
			log.Fatal(err)
		}
	}

	///// 初始化 路由角色
	routerRoles := GetRouterRoles()
	for _, rr := range routerRoles {
		_, err := E.AddPolicy(rr.RoleName, rr.RouterUri, rr.RouterMethod)
		if err != nil {
			log.Fatal(err)
		}
	}
}
