package account

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiny-sky/Tdtm/proto"
)

type Req struct {
	UserId string `json:"userId"`
	Amount int64  `json:"amount"`
}

type Srv struct {
}

func (srv *Srv) Try(ctx *gin.Context) {
	var req Req

	time.Sleep(300 * time.Millisecond)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, "[Account] Try err")
		return
	}
	fmt.Println("[Account] Try req:", req, time.Now().Unix())
	ctx.JSON(200, nil)
}

func (srv *Srv) Confirm(ctx *gin.Context) {
	var (
		req Req
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, fmt.Sprintf("[Account] Confirm err:%v", err))
		return
	}
	fmt.Println("[Account] Confirm req:", req)
	ctx.JSON(200, nil)
}

func (srv *Srv) Cancel(ctx *gin.Context) {
	var (
		req Req
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, "[Account] Cancel err")
		return
	}
	fmt.Println("[Account] Cancel req:", req)
	ctx.JSON(200, nil)
}

func Start(port int) {
	e := gin.Default()
	srv := new(Srv)
	e.POST("/account/try", srv.Try)
	e.POST("/account/confirm", srv.Confirm)
	e.POST("/account/cancel", srv.Cancel)

	go func() {
		err := e.Run(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("failed the account server:%v", err)
		}
	}()
	fmt.Println("account server start:", port)
}

func NewData() []byte {
	reqData := Req{
		UserId: "tiny_sky",
		Amount: 100,
	}
	b, _ := json.Marshal(reqData)
	return b
}

func RegisterTCC(port int) (branches []*proto.RegisterReq_Branch) {
	b := NewData()
	uri := fmt.Sprintf("http://localhost:%d", port)

	// try
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/account/try",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_TRY,
		Level:    1,
	})

	// confirm
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/account/confirm",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_CONFIRM,
		Level:    1,
	})
	// cancel
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/account/cancel",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_CANCEL,
		Level:    1,
	})
	return
}
