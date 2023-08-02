package utils

import (
	"ZN/ZN/ziface"
	"encoding/json"
	"io/ioutil"
)

type GlobalObj struct {
	//Server服务配置
	TcpServer ziface.IServer //ZN服务全局的Server对象
	Host      string         //当前服务器监听的IP地址
	TcpPort   int            //当前服务器监听的端口
	Name      string         //当前服务器的名称
	//ZN服务配置
	Version           string //当前版本
	MaxConn           int    //最大的连接数
	MaxPackageSize    uint32 //最大文件大小
	WorkerPoolSize    uint32 //当前业务线程池中的线程的数量
	MaxWorkerTaskSize uint32 //每个线程对应队列中的数据的数量的长度
}

// ReloadJSON 初始化读取配置文件的函数
func (o *GlobalObj) ReloadJSON() {
	//读取配置文件
	file, err := ioutil.ReadFile("conf/zn.json")
	if err != nil {
		panic(err)
	}

	//解析json数据
	err = json.Unmarshal(file, &GlobalOBJ)
	if err != nil {
		panic(err)
	}
}

// GlobalOBJ 定义一个唯一配置对象
var GlobalOBJ *GlobalObj

// 使用init函数初始化加载配置文件的数据
func init() {
	//设置默认参数
	GlobalOBJ = &GlobalObj{
		Name:              "ZNServer", //默认服务名称
		Version:           "V0.4",     //默认版本
		Host:              "0.0.0.0",  //默认IP地址
		TcpPort:           8080,       //默认端口
		MaxConn:           1000,       //默认最大连接数量
		MaxPackageSize:    4096,       //默认最大文件大小
		WorkerPoolSize:    10,         //当前业务线程池中的线程的数量
		MaxWorkerTaskSize: 1024,       //每个线程对应队列中的数据的数量的长度
	}

	//进行解析json配置文件
	GlobalOBJ.ReloadJSON()
}
