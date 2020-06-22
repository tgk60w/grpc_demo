package main

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gookit/rux"
	"google.golang.org/grpc"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
	clientFunc "grpc_demo/app/client/test"
	"grpc_demo/app/service/transfer"
	conf "grpc_demo/config"
	pb "grpc_demo/protos"
)

var instance *transfer.Transfer           //service
var connectsServer *clientFunc.ClientGrpc //client

type message struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

func main() {

	ch := make(chan os.Signal, 1)
	//监控 ctrl+c和 kill
	signal.Notify(ch, os.Interrupt, os.Kill)

	instance = transfer.New()

	connectsServer = new(clientFunc.ClientGrpc)

	go grpcService()                //启动grpc  service
	go webService()                 //启动web
	time.Sleep(1 * time.Second)
	go connectsServer.StartClient() //client连接service  为了好测试  所以一起启动

	<-ch

	close()
}
//close  关闭grpc连接    有时间再优化
func close() {
	connectsServer.Server.Close()
}

func webService() {
	// nat  队列包
	r := rux.New()

	r.POST("[/{name}]", func(c *rux.Context) {

		var data pb.WebJson
		var dataResult *pb.WebJson
		var err error
		var code int

		if c.Req.FormValue("name") == "" {
			c.JSON(http.StatusOK, message{Code: 400, Data: "请求参数为空"})
			return
		}

		uniquely := UniqueId()

		data.Data = []byte(c.Req.FormValue("name"))
		data.Id = uniquely

		instance.Push(&data)

		dataResult, err, code = instance.CustomData()

		if err != nil {

			c.JSON(http.StatusOK, message{Code: code, Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, message{Code: code, Data: fmt.Sprintf("%s", dataResult.Data)})

		return

	})

	r.GET("/web/{name}", func(c *rux.Context) {

		var data pb.WebJson
		var dataResult *pb.WebJson
		var err error
		var code int

		if c.Param("name") == "" {
			c.JSON(http.StatusOK, message{Code: 400, Data: "请求参数为空"})
			return
		}

		uniquely := UniqueId()

		data.Data = []byte(c.Param("name"))
		data.Id = uniquely

		instance.Push(&data)

		dataResult, err, code = instance.CustomData()

		if err != nil {

			c.JSON(http.StatusOK, message{Code: code, Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, message{Code: code, Data: fmt.Sprintf("%s", dataResult.Data)})

		return

	})
	// quick start
	r.Listen(conf.New().WebAddr)
}

func grpcService() {
	var err error

	if err != nil {
		panic(err)
	}

	conn, err := net.Listen("tcp", conf.New().Service)

	if err != nil {
		panic(err)
	}

	fmt.Printf("启动服务...")

	server := grpc.NewServer(grpc.UnaryInterceptor(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			//return nil, status.Error(codes.Canceled, "canceled")
			return handler(ctx, req)
		}), grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		//return status.Error(codes.Canceled, "canceled")
		return handler(srv, ss)
	}))

	pb.RegisterTransferServer(server, instance)

	if err := server.Serve(conn); err != nil {
		panic(err)
	}
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
