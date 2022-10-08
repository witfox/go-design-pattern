/*
单例模式-懒汉模式
比如：数据库连接，redis连接，验证码生成，邮箱初始化，发送短信初始化
*/
package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
