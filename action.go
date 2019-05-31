package main

import (
	"fmt"

	"github.com/dzhcool/genid/memcachep"
	"github.com/dzhcool/genid/snow"
)

var data map[string]string = make(map[string]string)

//初始化绑定处理程序
func init() {
	memcachep.BindAction(memcachep.GET, GetAction)
	memcachep.BindAction(memcachep.SET, SetAction)
}

func GetAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {
	id := snow.New().GenId()
	res.Fatal = false
	res.Value = []byte(fmt.Sprintf("%d", id))
}

func SetAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {
	res.Fatal = false
	res.Status = memcachep.STORED
}
