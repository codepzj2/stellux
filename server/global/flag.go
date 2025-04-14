package global

import (
	"flag"
	"fmt"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
)

var (
	Mode = flag.String("mode", "development", "运行模式,eg: development/production")
)

func ParseFlag() {
	flag.Parse()
}

func NewEnforcer(url string) *casbin.Enforcer {
	adapter, err := mongodbadapter.NewAdapter(url)
	if err != nil {
		panic(err)
	}
	enforcer, err := casbin.NewEnforcer("config/policy.conf", adapter)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		panic(fmt.Sprintf("权限加载失败：%s", err.Error()))
	}
	return enforcer
}
