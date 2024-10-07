package stock

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiny-sky/Tdtm/proto"
)

type Req struct {
	ProductId string `json:"ProductId"`
	Number    int64  `json:"number"`
}

type Srv struct {
}

func (srv *Srv) Try(ctx *gin.Context) {
	var req Req

	time.Sleep(300 * time.Millisecond)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, "[Stock] Try err")
		return
	}
	fmt.Println("[Stock] Try req:", req, time.Now().Unix())
	ctx.JSON(200, nil)
}

func (srv *Srv) Confirm(ctx *gin.Context) {
	var (
		req Req
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, fmt.Sprintf("[Stock] Confirm err:%v", err))
		return
	}
	fmt.Println("[Stock] Confirm req:", req)
	ctx.JSON(200, nil)
}

func (srv *Srv) Cancel(ctx *gin.Context) {
	var (
		req Req
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, "[Stock] Cancel err")
		return
	}
	fmt.Println("[Stock] Cancel req:", req)
	ctx.JSON(200, nil)
}

func NewData() []byte {
	req := &Req{
		ProductId: "520",
		Number:    10,
	}
	b, _ := json.Marshal(req)
	return b
}

func Start(port int) {
	e := gin.Default()
	srv := new(Srv)
	e.POST("/stock/try", srv.Try)
	e.POST("/stock/confirm", srv.Confirm)
	e.POST("/stock/cancel", srv.Cancel)

	go func() {
		err := e.Run(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("failed the stock server:%v", err)
		}
	}()
	fmt.Println("stock server start:", port)
}

func RegisterTCC(port int) (branches []*proto.RegisterReq_Branch) {
	b := NewData()
	uri := fmt.Sprintf("http://localhost:%d", port)

	// try
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/stock/try",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_TRY,
		Level:    1,
	})

	// confirm
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/stock/confirm",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_CONFIRM,
		Level:    1,
	})
	// cancel
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/stock/cancel",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_CANCEL,
		Level:    1,
	})
	return
}
