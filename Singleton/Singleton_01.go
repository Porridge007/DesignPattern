package Singleton

// 单例模式采用 饿汉式和懒汉式两种实现
// Singleton_01.go 采用

type Singleton struct{}

var singleton *Singleton

func init(){
	singleton = &Singleton{}
}

func GetInstance() *Singleton{
	return singleton
}