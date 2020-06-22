/*
@Time : 2020/6/19 上午9:30
@Author : gongyikun
@File : main
@Software: GoLand
*/
package main


//type ClientGrpc struct {
//	Server *grpc.ClientConn
//	Client pb.TransferClient
//}

//func  main() {
//	c := new(ClientGrpc)
//
//	conn, err := grpc.Dial(conf.New().Service, grpc.WithInsecure())
//
//	if err != nil {
//
//		fmt.Printf("连接失败:%v", err)
//
//	}
//	//
//	defer conn.Close()
//
//	c.Server = conn
//
//	c.Client = pb.NewTransferClient(conn)
//
//	ctx := context.Background()
//
//	stream, err := c.Client.WebRequest(ctx, &pb.WebJson{})
//
//	if err != nil {
//		log.Printf("创建数据流失败: [%v] ", err)
//	}
//
//	// 启动一个 goroutine 接收命令行输入的指令
//	for {
//		// 接收从 服务端返回的数据流
//		result, err := stream.Recv()
//
//		if err == io.EOF {
//			log.Println("⚠️ 收到服务端的结束信号")
//			break //如果收到结束信号，则退出“接收循环”，结束客户端程序
//		}
//		if err != nil {
//			// TODO: 处理接收错误
//			log.Println("接收数据出错:", err)
//			break
//		}
//
//		if len(result.Data) > 0 {
//			ctx := context.Background()
//			var requestData pb.WebJson
//
//			datas := []byte("ervice给client发送的数据"+string(result.Data))
//
//			requestData.Data = datas
//			requestData.Id = result.Id
//
//			c.Client.ReturnData(ctx, &requestData)
//		}
//
//	}
//}
