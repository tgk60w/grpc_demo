package transfer

import (
	"container/list"
	"context"
	"errors"
	"time"
	conf "grpc_demo/config"
	"grpc_demo/protos"
)

type Transfer struct {
	listWeb     *list.List
	listResult  *list.List
	ch          chan int //等待返回结果
	timingClose chan int //关闭计时器
}

func New() *Transfer {

	instance := new(Transfer)

	instance.listWeb = list.New()
	instance.listResult = list.New()
	instance.ch = make(chan int, conf.New().ChanNum)
	instance.timingClose = make(chan int, conf.New().ChanNum)

	return instance
}
//Push  队列中加入新值    队列不向外开放
func(t *Transfer) Push(data *protos.WebJson)  {

	t.listWeb.PushBack(data)

}

//CustomData  从返回队列中取数据
func (t *Transfer) CustomData() (*protos.WebJson, error, int) {
	var err error
	timingChan := make(chan int)
	//计时器   超时会直接返回
	go func() {

		select {
		case <-time.After(time.Second * 5):
			timingChan <- 1
		case <-t.timingClose:
			return
		}
	}()

	select {
	case <-t.ch:

		dataResult := t.listResult.Front()

		webJson, ok := dataResult.Value.(protos.WebJson)

		if ok == false {

			return nil, errors.New("数据解析错误"), 400
		}

		t.listResult.Remove(dataResult)

		return &webJson, err, 200

	case <-timingChan:

		return nil, errors.New("超时了兄弟"), 500

	}

}

//WebRequest 向client推数据
func (t *Transfer) WebRequest(req *protos.WebJson, stream protos.Transfer_WebRequestServer) error {
	var err error
	var webJson protos.WebJson
	var ok bool

	for {
		//也可用select  触发
		if t.listWeb.Len() > 0 {

			str := t.listWeb.Front()

			webJson, ok = str.Value.(protos.WebJson)

			if ok == false {

				t.listWeb.Remove(str)
				continue
			}

			if len(webJson.Data) > 0 {

				err = stream.Send(&webJson)

				if err != nil {
					return err
				} else {
					t.listWeb.Remove(str)
				}
			}

		}

	}
}
//ReturnData  接收client  返回的数据
func (t *Transfer) ReturnData(ctx context.Context, in *protos.WebJson) (*protos.WebJson, error) {
	var err error

	if in != (&protos.WebJson{}) {

		t.listResult.PushBack(protos.WebJson{Data: in.Data, Id: in.Id})

		t.timingClose <- 1 //关闭计时
		t.ch <- 1          //通道  写入值   释放阻塞

	}

	return &protos.WebJson{}, err
}
