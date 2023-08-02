# 使用

## 1、获取
使用Git直接拉取文件就可使用

## 2、配置文件

项目根目录下创建conf文件夹并且存放一个zn.json文件

```
{
    "Name": "Zn_V0.5_DemoServer",
    "Host": "127.0.0.1",
    "TcpPort": 8080,
    "MaxConn" : 10,
    "MaxPackageSize": 1024,
    "WorkerPoolSize": 10
}
```

## 3、创建服务


```

type TestRouter struct {
    znet.Router
}

// Handle 处理conn业务的方法
func (r *TestRouter) Handle(request ziface.IRequest) {

fmt.Printf("Run Handle Data: %s ID : %d \n", request.GetMsgData(), request.GetMsgID())
    err := request.GetConnection().SendMsg(200, []byte("ping...ping..."))
    if err != nil {
        fmt.Println("SendMsg err :", err)
    return
        }
}


func HStart(conn ziface.IConnection) {
    conn.SendMsg(201, []byte("连接成功"))
    fmt.Println("连接成功")
    conn.SetProperty("Name", "张三")
}
func HStop(conn ziface.IConnection) {
    fmt.Println("断开连接")
    if v, err := conn.GetProperty("Name"); err == nil {
        fmt.Println("Name:", v)
    }
}

func main() {
    s := znet.NewZServer()

    s.SetOnConnStartHook(HStart)
    s.SetOnConnStopHook(HStop)

    s.AddRouter(0, &TestRouter{})
    s.AddRouter(1, &TestRouter2{})

    s.Server()
}

```


## 4、连接服务

```
func main() {
	fmt.Println("Client Run ....")

	//等待时长
	time.Sleep(3 * time.Second)

	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Client Dial err", err)
		return
	}

	//进行数据交互
	for {
		//发送信息
		//进行封包
		dp := znet.NewDataPack()

		bytes, err := dp.Pack(znet.NewMessage(0, []byte("ZN-V0.5-你好-0")))
		if err != nil {
			fmt.Println("Client Pack err:", err)
			return
		}

		if _, err := conn.Write(bytes); err != nil {
			fmt.Println("Client Write err:", err)
			return
		}

		//接收消息进行解包

		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("Client Head ReadFull err:", err)
			return
		}

		//进行头部数据解包操作
		pack, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("Client UnPack err:", err)
			return
		}
		if pack.GetDataLen() > 0 {
			data := make([]byte, pack.GetDataLen())
			_, err := io.ReadFull(conn, data)
			if err != nil {
				fmt.Println("Client Data ReadFull err:", err)
				return
			}
			pack.SetData(data)
			fmt.Println("==> Recv Msg: ID=", pack.GetMsgID(), ", len=", pack.GetDataLen(), ", data=", string(pack.GetData()))
		}

		time.Sleep(1 * time.Second)
	}
}
```




